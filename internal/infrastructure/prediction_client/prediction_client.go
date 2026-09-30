package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
)

// Family selects one of the model service's prediction routes.
type Family string

const (
	EffectiveTension Family = "effect_na"
	HookLoad         Family = "ves_na_kru"
	Torque           Family = "moment"
	MinWeight        Family = "min_ves"
)

var (
	// ErrUnavailable means the model service could not be reached or is not ready (HTTP 503).
	ErrUnavailable = errors.New("the prediction service is unavailable; try again later")
	// ErrInvalidResponse means the model service answered with something unusable (HTTP 502).
	ErrInvalidResponse = errors.New("the prediction service returned an invalid response")
)

// depthKey is filled by the API from the request; the model does not return it.
const depthKey = "Глубина"

const (
	requestTimeout   = 60 * time.Second
	maxResponseBytes = 32 << 20
)

type TorqueAndDragClient interface {
	// Predict posts one feature request and decodes a validated response into out, a pointer
	// to a struct whose json tags name the expected series.
	Predict(ctx context.Context, family Family, data requests.TorqueAndDragFromMLModelRequest, out any) error
	// Ready reports whether the model service has loaded and smoke-tested its artifacts.
	Ready(ctx context.Context) error
}

type torqueAndDragClient struct {
	client  *http.Client
	baseURL string
}

func NewTorqueAndDragClient(baseURL string) TorqueAndDragClient {
	return &torqueAndDragClient{client: &http.Client{Timeout: requestTimeout}, baseURL: strings.TrimRight(baseURL, "/")}
}

func (c *torqueAndDragClient) Predict(ctx context.Context, family Family, data requests.TorqueAndDragFromMLModelRequest, out any) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encode prediction request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s/", c.baseURL, family), bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("build prediction request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes))
	if err != nil {
		return fmt.Errorf("%w: read body: %v", ErrUnavailable, err)
	}
	switch {
	case resp.StatusCode == http.StatusServiceUnavailable:
		return fmt.Errorf("%w: model artifacts are not ready", ErrUnavailable)
	case resp.StatusCode != http.StatusOK:
		return fmt.Errorf("%w: %s: %.300s", ErrInvalidResponse, resp.Status, body)
	}
	return decodeValidated(body, len(data.MD), out)
}

func (c *torqueAndDragClient) Ready(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	// The prediction routes live under /predict; readiness is at the service root.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.TrimSuffix(c.baseURL, "/predict")+"/ready", nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnavailable, err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: readiness returned %s", ErrUnavailable, resp.Status)
	}
	return nil
}

// decodeValidated requires every expected series, one finite value per station, before decoding.
func decodeValidated(body []byte, stations int, out any) error {
	var series map[string][]float64
	if err := json.Unmarshal(body, &series); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}
	for _, key := range expectedSeries(out) {
		values, ok := series[key]
		if !ok {
			return fmt.Errorf("%w: series %q is missing", ErrInvalidResponse, key)
		}
		if len(values) != stations {
			return fmt.Errorf("%w: series %q has %d values for %d stations", ErrInvalidResponse, key, len(values), stations)
		}
		for _, v := range values {
			if math.IsNaN(v) || math.IsInf(v, 0) {
				return fmt.Errorf("%w: series %q contains a non-finite value", ErrInvalidResponse, key)
			}
		}
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}
	return nil
}

func expectedSeries(out any) []string {
	t := reflect.TypeOf(out)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	var keys []string
	for i := 0; i < t.NumField(); i++ {
		tag := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		if tag != "" && tag != "-" && tag != depthKey {
			keys = append(keys, tag)
		}
	}
	return keys
}
