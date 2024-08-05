package apiclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaelherik/terraform-provider-aznamingtool/apiclient/models"
	"github.com/stretchr/testify/assert"
)

func TestDoGet(t *testing.T) {
	expectedResponse := []models.ResourceType{{Id: 1, Resource: "Resource1"}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody, _ := json.Marshal(expectedResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}))
	defer server.Close()
	httpClient := server.Client()
	client := NewAPIClient(server.URL, "123456", "123456", httpClient)
	service := NewBaseService(client)

	var actualResponse []models.ResourceType
	err := service.DoGet("RequestName", nil, &actualResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestDoPost(t *testing.T) {

	requestData := models.ResourceUnit{ResourceBaseEntity: models.ResourceBaseEntity{Id: 1, Name: "Resource1"}}
	var expectedResponse = requestData

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqData models.ResourceUnit
		json.NewDecoder(r.Body).Decode(&reqData)
		responseBody, _ := json.Marshal(expectedResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}))
	defer server.Close()

	httpClient := server.Client()
	client := NewAPIClient(server.URL, "123456", "123456", httpClient)
	service := NewBaseService(client)

	var actualResponse models.ResourceUnit
	err := service.DoPost("CreateOrUpdateResourceUnit", requestData, &actualResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestDoDelete(t *testing.T) {
	requestData := models.ResourceUnit{ResourceBaseEntity: models.ResourceBaseEntity{Id: 1, Name: "Resource1"}}
	var expectedResponse = requestData

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody, _ := json.Marshal(expectedResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}))
	defer server.Close()

	httpClient := server.Client()
	client := NewAPIClient(server.URL, "123456", "123456", httpClient)
	service := NewBaseService(client)

	var actualResponse models.ResourceUnit
	service.DoDelete("DeleteResourceUnit", map[string]string{"id": "1"}, &actualResponse)

	assert.Equal(t, expectedResponse, actualResponse)

}
