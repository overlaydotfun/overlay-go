package overlay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

// sendRequest is a helper function to handle API requests.
// It returns a flattened map of the response.
func (c *Client) sendRequest(method, path string, params map[string]interface{}, withAuth bool) (map[string]interface{}, error) {
	var jsonBody []byte
	var err error
	if params != nil {
		jsonBody, err = json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}
	url := c.baseURL + path
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("overlay-go/1.0.0 (%s)", runtime.Version()))

	// Conditionally add auth headers
	if withAuth {
		if c.config.APIKey != "" {
			req.Header.Set("api-key", c.config.APIKey)
		}
		if c.config.AuthKey != "" {
			req.Header.Set("auth-key", c.config.AuthKey)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the raw response into the Response struct
	rawResponse := &Response{}
	if err := json.Unmarshal(body, rawResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(body))
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("http error: %d, message: %s", resp.StatusCode, rawResponse.Message)
	}

	// Flatten the response: merge data into the top level
	flatResponse := make(map[string]interface{})
	flatResponse["success"] = rawResponse.Success
	flatResponse["message"] = rawResponse.Message

	if rawResponse.Data != nil {
		for key, value := range rawResponse.Data {
			flatResponse[key] = value
		}
	}

	return flatResponse, nil
}
