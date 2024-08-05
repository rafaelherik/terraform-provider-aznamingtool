package apiclient

import "github.com/rafaelherik/terraform-provider-aznamingtool/apiclient/models"

type ResourceTypeService struct {
	baseService *BaseService
}

// NewResourceTypeService creates a new instance of ResourceTypeService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceTypeService instance.
func NewResourceTypeService(client *APIClient) *ResourceTypeService {
	return &ResourceTypeService{baseService: NewBaseService(client)}
}

// GetAllResourceTypes retrieves all resource types.
//
// Returns:
//   - A pointer to a slice of models.ResourceType containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceTypeService) GetAllResourceTypes() (*[]models.ResourceType, error) {
	var response []models.ResourceType
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceType retrieves a resource type based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource type.
//
// Returns:
//   - A pointer to models.ResourceType containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceTypeService) GetResourceType(id string) (*models.ResourceType, error) {
	var response models.ResourceType
	err := s.baseService.DoGet("GetResourceType", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
