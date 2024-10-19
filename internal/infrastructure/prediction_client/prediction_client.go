// Create a new file named torque_and_drag_client.go under `internal/infrastructure/client`

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
)

const (
	timeoutSeconds = 10
)

type TorqueAndDragClient interface {
	CalculateEffectiveTension(data requests.EffectiveTensionFromMLModelRequest) (*responses.EffectiveTensionFromMLModelResponse, error)
}

type torqueAndDragClient struct {
	baseURL string
}

func NewTorqueAndDragClient(baseURL string) TorqueAndDragClient {
	return &torqueAndDragClient{
		baseURL: baseURL,
	}
}

// CalculateEffectiveTension sends a request to the FastAPI service to calculate effective tension
func (c *torqueAndDragClient) CalculateEffectiveTension(data requests.EffectiveTensionFromMLModelRequest) (*responses.EffectiveTensionFromMLModelResponse, error) {
	url := fmt.Sprintf("%s/effect_na/", c.baseURL)

	// Serialize the data to JSON
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %v", err)
	}

	// Set up the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Execute the request with a timeout
	client := &http.Client{Timeout: time.Duration(timeoutSeconds) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	// Decode the response
	var effectiveTensionResponse responses.EffectiveTensionFromMLModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&effectiveTensionResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &effectiveTensionResponse, nil
}