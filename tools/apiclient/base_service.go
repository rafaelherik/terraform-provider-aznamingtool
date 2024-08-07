package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type BaseService struct {
	client *APIClient
}

// NewBaseService creates a new instance of BaseService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created BaseService instance.
func NewBaseService(client *APIClient) *BaseService {
	return &BaseService{client: client}
}

// DoGet performs a GET request to the specified endpoint with URL interpolation
// and decodes the response into the provided response object.
//
// Parameters:
//   - endpointKey: The key to the API endpoint in the client's endpoint map.
//   - uriData: A map containing data to be interpolated into the endpoint URL.
//   - response: A pointer to a variable where the response should be decoded.
//
// Returns:
//   - error: An error if any of the following occurs:
//   - The endpoint is not found in the client's endpoint map.
//   - The request creation fails.
//   - The request execution fails.
//   - The response body decoding fails.
//   - The response status code is 400 or greater.
func (s *BaseService) DoGet(endpointKey string, uriData map[string]string, response interface{}) error {

	if s == nil {
		return fmt.Errorf("BaseService is nil")
	}
	if endpointKey == "" {
		return fmt.Errorf("url is empty")
	}
	if s.client == nil {
		return fmt.Errorf("client is nil")
	}
	if s.client.ApiEndpoints == nil {
		return fmt.Errorf("ApiEndpoints is nil")
	}

	endpoint, exists := s.client.ApiEndpoints[endpointKey]
	if !exists {
		return fmt.Errorf("endpoint not found for %s", endpointKey)
	}

	// Perform string interpolation with uriData
	for key, value := range uriData {
		placeholder := fmt.Sprintf("{%s}", key)
		endpoint = strings.Replace(endpoint, placeholder, value, -1)
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.DoRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	return nil
}

// DoPost performs a POST request to the specified endpoint with a JSON-encoded body
// and decodes the response into the provided response object.
//
// Parameters:
//   - endpointKey: A string representing the key to the API endpoint in the client's endpoint map.
//   - requestData: An object that will be serialized into a JSON object to be included in the POST request body.
//   - response: A pointer to a variable where the decoded response should be stored.
//
// Returns:
//   - error: An error if any of the following occurs:
//   - The endpoint is not found in the client's endpoint map.
//   - The request creation fails.
//   - The request execution fails.
//   - The response body decoding fails.
//   - The response status code is 400 or greater.
func (s *BaseService) DoPost(endpointKey string, requestData interface{}, response interface{}) error {
	endpoint, exists := s.client.ApiEndpoints[endpointKey]
	if !exists {
		return fmt.Errorf("endpoint not found for %s", endpointKey)
	}

	var req *http.Request
	var err error

	if requestData != nil {
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			return err
		}
		req, err = http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest("POST", endpoint, nil)
		if err != nil {
			return err
		}
	}

	resp, err := s.client.DoRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}

// DoDelete performs a DELETE request to the specified endpoint with URL interpolation
// and decodes the response into the provided response object.
//
// Parameters:
//   - endpointKey: The key to the API endpoint in the client's endpoint map.
//   - uriData: A map containing data to be interpolated into the endpoint URL.
//   - response: A pointer to a variable where the response should be decoded.
//
// Returns://
//   - error: An error if any of the following occurs:
//   - The endpoint is not found in the client's endpoint map.
//   - The request creation fails.
//   - The request execution fails.
//   - The response body decoding fails.
func (s *BaseService) DoDelete(endpointKey string, uriData map[string]string) error {
	endpoint, exists := s.client.ApiEndpoints[endpointKey]
	if !exists {
		return fmt.Errorf("endpoint not found for %s", endpointKey)
	}

	// Perform string interpolation with uriData
	for key, value := range uriData {
		placeholder := fmt.Sprintf("{%s}", key)
		endpoint = strings.Replace(endpoint, placeholder, value, -1)
	}

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.DoRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
