package apiclient

import "github.com/rafaelherik/terraform-provider-aznamingtool/apiclient/models"

type ResourceComponentService struct {
	baseService *BaseService
}

// NewResourceComponentService creates a new instance of ResourceComponentService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceComponentService instance.
func NewResourceComponentService(client *APIClient) *ResourceComponentService {
	return &ResourceComponentService{baseService: NewBaseService(client)}
}

// GetAllResourceComponents retrieves all resource components.
//
// Returns:
//   - A pointer to a slice of models.ResourceComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceComponentService) GetAllResourceComponents() (*[]models.ResourceComponent, error) {
	var response []models.ResourceComponent
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceComponent retrieves a resource component based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource component.
//
// Returns:
//   - A pointer to models.ResourceComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceComponentService) GetResourceComponent(id string) (*models.ResourceComponent, error) {
	var response models.ResourceComponent
	err := s.baseService.DoGet("GetResourceComponent", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceComponent creates or updates a resource component based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceComponent containing the request data.
//
// Returns:
//   - A pointer to models.ResourceComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceComponentService) CreateOrUpdateResourceComponent(request models.ResourceComponent) (*models.ResourceComponent, error) {
	var response models.ResourceComponent
	err := s.baseService.DoPost("CreateOrUpdateResourceComponent", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
