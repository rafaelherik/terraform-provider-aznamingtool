package apiclient

import "github.com/rafaelherik/terraform-provider-aznamingtool/apiclient/models"

type ResourceEnvironmentService struct {
	baseService *BaseService
}

// NewResourceEnvironmentService creates a new instance of ResourceEnvironmentService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceEnvironmentService instance.
func NewResourceEnvironmentService(client *APIClient) *ResourceEnvironmentService {
	return &ResourceEnvironmentService{baseService: NewBaseService(client)}
}

// GetAllResourceEnvironments retrieves all resource environments.
//
// Returns:
//   - A pointer to a slice of models.ResourceEnvironment containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceEnvironmentService) GetAllResourceEnvironments() (*[]models.ResourceEnvironment, error) {
	var response []models.ResourceEnvironment
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceEnvironment retrieves a resource environment based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource environment.
//
// Returns:
//   - A pointer to models.ResourceEnvironment containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceEnvironmentService) GetResourceEnvironment(id string) (*models.ResourceEnvironment, error) {
	var response models.ResourceEnvironment
	err := s.baseService.DoGet("GetResourceEnvironment", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceEnvironment creates or updates a resource environment based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceEnvironment containing the request data.
//
// Returns:
//   - A pointer to models.ResourceEnvironment containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceEnvironmentService) CreateOrUpdateResourceEnvironment(request models.ResourceEnvironment) (*models.ResourceEnvironment, error) {
	var response models.ResourceEnvironment
	err := s.baseService.DoPost("CreateOrUpdateResourceEnvironment", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteResourceEnvironment deletes a resource environment based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource environment.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceEnvironmentService) DeleteResourceEnvironment(id string) (interface{}, error) {
	var response models.ResourceEnvironment
	return s.baseService.DoDelete("DeleteResourceEnvironment", map[string]string{"id": id}, &response)
}
