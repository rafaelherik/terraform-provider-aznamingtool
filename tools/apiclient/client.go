package apiclient

import (
	"fmt"
	"net/http"
)

// APIClient provides a client for making API requests to the resource naming service.
type APIClient struct {
	BaseURL       string            // The base URL of the API.
	APIKey        string            // The API key for authenticating requests.
	AdminPassword string            // The admin password for authenticating requests.
	ApiEndpoints  map[string]string // A map of endpoint keys to endpoint URLs.
	HttpClient    *http.Client      // The HTTP client used to make requests.
	requestQueue  chan requestEntry // A channel to queue requests
}

type requestEntry struct {
	req  *http.Request
	resp chan responseEntry
}

type responseEntry struct {
	resp *http.Response
	err  error
}

// NewAPIClient creates a new instance of APIClient with the provided base URL and API key.
func NewAPIClient(baseURL string, apiKey string, adminPassword string, httpClient *http.Client) *APIClient {
	if baseURL == "" {
		panic("baseURL cannot be empty")
	}
	if apiKey == "" {
		panic("apiKey cannot be empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	client := &APIClient{
		BaseURL:       baseURL,
		APIKey:        apiKey,
		AdminPassword: adminPassword,
		ApiEndpoints: map[string]string{
			//Resource Naming
			"RequestName":               baseURL + "/api/ResourceNamingRequests/RequestName",
			"RequestNameWithComponents": baseURL + "/api/ResourceNamingRequests/RequestNameWithComponents",
			"ValidateName":              baseURL + "/api/ResourceNamingRequests/ValidateName",
			"GetGeneratedName":          baseURL + "/api/Admin/GetGeneratedName/{id}",
			"DeleteGeneratedName":       baseURL + "/api/Admin/DeleteGeneratedName/{id}",

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

			// Resource Types
			"GetAllResourceTypes": baseURL + "/api/ResourceTypes",

			//Resource Units
			"CreateOrUpdateResourceUnit": baseURL + "/api/ResourceUnitDepts",
			"DeleteResourceUnit":         baseURL + "/api/ResourceUnitDepts/{id}",
		},
		HttpClient:   httpClient,
		requestQueue: make(chan requestEntry, 100), // Buffered channel to queue requests
	}

	go client.processQueue()

	return client
}

// processQueue processes the queued requests one by one.
func (c *APIClient) processQueue() {
	for entry := range c.requestQueue {
		resp, err := c.doRequest(entry.req)
		entry.resp <- responseEntry{resp: resp, err: err}
		close(entry.resp)
	}
}

// doRequest sends an HTTP request using the client's HTTP client, adding the API key to the request headers.
func (c *APIClient) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("APIKey", c.APIKey)

	if c.AdminPassword != "" {
		req.Header.Set("AdminPassword", c.AdminPassword)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

// DoRequest adds the request to the queue to be processed sequentially.
func (c *APIClient) DoRequest(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	// Create a copy of the request to avoid concurrent modifications
	reqCopy := req.Clone(req.Context())

	respChan := make(chan responseEntry)
	c.requestQueue <- requestEntry{req: reqCopy, resp: respChan}

	result := <-respChan
	return result.resp, result.err
}
