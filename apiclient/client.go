package apiclient

import (
	"fmt"
	"net/http"
)

// APIClient provides a client for making API requests to the resource naming service.
type APIClient struct {
	BaseURL      string            // The base URL of the API.
	APIKey       string            // The API key for authenticating requests.
	AdminPassord string            // The admin password for authenticating requests.
	ApiEndpoints map[string]string // A map of endpoint keys to endpoint URLs.
	HttpClient   *http.Client      // The HTTP client used to make requests.
}

// NewAPIClient creates a new instance of APIClient with the provided base URL and API key.
//
// Parameters:
//   - baseURL: A string representing the base URL of the API.
//   - apiKey: A string representing the API key for authentication.
//
// Returns:
//   - A pointer to the newly created APIClient instance.
func NewAPIClient(baseURL string, apiKey string, admin_password string, HttpClient *http.Client) *APIClient {
	if baseURL == "" {
		panic("baseURL cannot be empty")
	}
	if apiKey == "" {
		panic("apiKey cannot be empty")
	}

	httpClientInstance := HttpClient

	if HttpClient == nil {
		httpClientInstance = &http.Client{}
	}

	return &APIClient{
		BaseURL:      baseURL,
		APIKey:       apiKey,
		AdminPassord: admin_password,
		ApiEndpoints: map[string]string{
			//Resource Naming
			"RequestName":               baseURL + "/api/ResourceNamingRequests/RequestName",
			"RequestNameWithComponents": baseURL + "/api/ResourceNamingRequests/RequestNameWithComponents",
			"ValidateName":              baseURL + "/api/ResourceNamingRequests/ValidateName",
			"GetGeneratedName":          baseURL + "/api/Admin/GetGeneratedName/{id}",

			// Custom Components
			"GetAllCustomComponents":          baseURL + "/api/CustomComponents",
			"GetCustomComponent":              baseURL + "/api/CustomComponents/{id}",
			"GetCustomComponentByParentId":    baseURL + "/api/CustomComponents/GetByParentId/{parentComponentId}",
			"GetCustomComponentByParentType":  baseURL + "/api/CustomComponents/GetByParentType/{parentComponentType}",
			"CreateOrUpdateCustomComponent":   baseURL + "/api/CustomComponents",
			"DeleteCustomComponent":           baseURL + "/api/CustomComponents/{id}",
			"DeleteCustomComponentByParentId": baseURL + "/api/CustomComponents/DeleteByParentId/{parentComponentId}",

			// Resource Components
			"GetAllResourceComponents":        baseURL + "/api/ResourceComponents",
			"GetResourceComponent":            baseURL + "/api/ResourceComponents/{id}",
			"CreateOrUpdateResourceComponent": baseURL + "/api/ResourceComponents",

			// Resource Delimiters
			"GetAllResourceDelimiters":        baseURL + "/api/ResourceDelimiters",
			"GetResourceDelimiter":            baseURL + "/api/ResourceDelimiters/{id}",
			"CreateOrUpdateResourceDelimiter": baseURL + "/api/ResourceDelimiters",

			// Resource Environments
			"GetAllResourceEnvironments":        baseURL + "/api/ResourceEnvironments",
			"CreateOrUpdateResourceEnvironment": baseURL + "/api/ResourceEnvironments",
			"DeleteResourceEnvironment":         baseURL + "/api/ResourceEnvironments/{id}",

			// Resource Functions
			"GetAllResourceFunctions":        baseURL + "/api/ResourceFunctions",
			"GetResourceFunction":            baseURL + "/api/ResourceFunctions/{id}",
			"CreateOrUpdateResourceFunction": baseURL + "/api/ResourceFunctions",
			"DeleteResourceFunction":         baseURL + "/api/ResourceFunctions/{id}",

			// Resource Locations
			"GetAllResourceLocations":        baseURL + "/api/ResourceLocations",
			"GetResourceLocation":            baseURL + "/api/ResourceLocations/{id}",
			"CreateOrUpdateResourceLocation": baseURL + "/api/ResourceLocations",
			"DeleteResourceLocation":         baseURL + "/api/ResourceLocations/{id}",

			// Resource Organizations

			//Resource Units
			"CreateOrUpdateResourceUnit": baseURL + "/api/ResourceUnitDepts",
			"DeleteResourceUnit":         baseURL + "/api/ResourceUnitDepts/{id}",
		},
		HttpClient: httpClientInstance,
	}
}

// DoRequest sends an HTTP request using the client's HTTP client, adding the API key to the request headers.
//
// Parameters:
//   - req: A pointer to the http.Request to be sent.
//
// Returns:
//   - A pointer to the http.Response received.
//   - An error if the request fails.
func (c *APIClient) DoRequest(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	req.Header.Set("APIKey", c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	if resp.StatusCode >= 400 {
		return resp, fmt.Errorf("received error status code %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return resp, nil
}
