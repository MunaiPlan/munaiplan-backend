// Create a new file named torque_and_drag_client.go under `internal/infrastructure/client`

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
)

const (
	timeoutSeconds = 10
)

type TorqueAndDragClient interface {
	CalculateEffectiveTension(data requests.TorqueAndDragFromMLModelRequest) (*responses.EffectiveTensionFromMLModelResponse, error)
	CalculateWeightOnBit(data requests.TorqueAndDragFromMLModelRequest) (*responses.WeightOnBitFromMLModelResponse, error)
	CalculateMoment(data requests.TorqueAndDragFromMLModelRequest) (*responses.MomentFromMLModelResponse, error)
	CalculateMinWeight(data requests.TorqueAndDragFromMLModelRequest) (*responses.MinWeightFromMLModelResponse, error)
}

type torqueAndDragClient struct {
	client *http.Client
	baseURL string
}

func NewTorqueAndDragClient(baseURL string) TorqueAndDragClient {
	return &torqueAndDragClient{
		client: &http.Client{
			Timeout: time.Duration(timeoutSeconds) * time.Second,
		},
		baseURL: baseURL,
	}
}

// CalculateEffectiveTension sends a request to the FastAPI service to calculate effective tension
func (c *torqueAndDragClient) CalculateEffectiveTension(data requests.TorqueAndDragFromMLModelRequest) (*responses.EffectiveTensionFromMLModelResponse, error) {
	url := fmt.Sprintf("%s/effect_na/", c.baseURL)

	req, err := preparePostRequest(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %v", err)
	}

	// Execute the request with a timeout
	resp, err := c.client.Do(req)
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

// CalculateWeightOnBit sends a request to the FastAPI service to calculate weight on bit
func (c *torqueAndDragClient) CalculateWeightOnBit(data requests.TorqueAndDragFromMLModelRequest) (*responses.WeightOnBitFromMLModelResponse, error) {
	url := fmt.Sprintf("%s/ves_na_kru/", c.baseURL)

	req, err := preparePostRequest(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %v", err)
	}

	// Execute the request with a timeout
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	// Decode the response
	var response responses.WeightOnBitFromMLModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}

// // CalculateMoment sends a request to the FastAPI service to calculate moment
func (c *torqueAndDragClient) CalculateMoment(data requests.TorqueAndDragFromMLModelRequest) (*responses.MomentFromMLModelResponse, error) {
	url := fmt.Sprintf("%s/moment/", c.baseURL)

	req, err := preparePostRequest(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %v", err)
	}

	// Execute the request with a timeout
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	// Decode the response
	var response responses.MomentFromMLModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}

// CalculateMinWeight sends a request to the FastAPI service to calculate minimum weight
func (c *torqueAndDragClient) CalculateMinWeight(data requests.TorqueAndDragFromMLModelRequest) (*responses.MinWeightFromMLModelResponse, error) {
	url := fmt.Sprintf("%s/min_ves/", c.baseURL)

	req, err := preparePostRequest(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %v", err)
	}

	// Execute the request with a timeout
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	// Decode the response
	var response responses.MinWeightFromMLModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}

// preparePostRequest prepares an HTTP POST request with the given URL and data, setting the appropriate headers.
func preparePostRequest(url string, data requests.TorqueAndDragFromMLModelRequest) (*http.Request, error) {
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

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
