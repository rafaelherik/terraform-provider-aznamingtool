package apiclient

import (
	"fmt"

	"github.com/rafaelherik/azure-naming-provider/tf/apiclient/models"
)

// ResourceNamingService provides methods for requesting and validating resource names.
type ResourceNamingService struct {
	baseService *BaseService
}

// NewResourceNamingService creates a new instance of ResourceNamingService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceNamingService instance.
func NewResourceNamingService(client *APIClient) *ResourceNamingService {
	return &ResourceNamingService{baseService: NewBaseService(client)}
}

// RequestName requests a new resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceNamingService) RequestName(request models.ResourceNameRequest) (*models.ResourceNameResponse, error) {
	var response models.ResourceNameResponse
	err := s.baseService.DoPost("RequestName", request, &response)
	if err != nil {
		return nil, err
	}

	if !response.Success {
		return &response, fmt.Errorf("request failed: %v", response)
	}

	return &response, nil
}

// RequestNameWithComponents requests a new resource name with components based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceNameRequestWithComponents containing the request data.
//
// Returns:
//   - A pointer to models.ResourceNameResponse containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceNamingService) RequestNameWithComponents(request models.ResourceNameRequestWithComponents) (*models.ResourceNameResponse, error) {
	var response models.ResourceNameResponse
	err := s.baseService.DoPost("RequestName", request, &response)
	if err != nil {
		return nil, err
	}

	if !response.Success {
		return &response, fmt.Errorf("request failed: %v", response)
	}

	return &response, nil
}

// ValidatetName validates a resource name based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ValidateNameRequest containing the request data.
//
// Returns:
//   - A pointer to models.ValidateNameResponse containing the validation response.
//   - An error if the request fails.
func (s *ResourceNamingService) ValidatetName(request models.ValidateNameRequest) (*models.ValidateNameResponse, error) {
	var response models.ValidateNameResponse
	err := s.baseService.DoPost("ValidateName", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetGeneratedName retrieves a generated resource name by its ID.
//
// Parameters:
//   - id: A string representing the ID of the generated resource name.
//
// Returns:
//   - A pointer to models.ResourceGeneratedName containing the generated name data.
//   - An error if the request fails.
func (s *ResourceNamingService) GetGeneratedName(id string) (*models.ResourceGeneratedName, error) {
	var response models.ResourceGeneratedName
	err := s.baseService.DoGet("GetGeneratedName", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
