package apiclient

import "github.com/rafaelherik/terraform-provider-aznamingtool/apiclient/models"

type ResourceProjectService struct {
	baseService *BaseService
}

// NewResourceProjectService creates a new instance of ResourceProjectService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceProjectService instance.
func NewResourceProjectService(client *APIClient) *ResourceProjectService {
	return &ResourceProjectService{baseService: NewBaseService(client)}
}

// GetAllResourceProjects retrieves all resource projects.
//
// Returns:
//   - A pointer to a slice of models.ResourceProject containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceProjectService) GetAllResourceProjects() (*[]models.ResourceProject, error) {
	var response []models.ResourceProject
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceProject retrieves a resource project based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource project.
//
// Returns:
//   - A pointer to models.ResourceProject containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceProjectService) GetResourceProject(id string) (*models.ResourceProject, error) {
	var response models.ResourceProject
	err := s.baseService.DoGet("GetResourceProject", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceProject creates or updates a resource project based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceProject containing the request data.
//
// Returns:
//   - A pointer to models.ResourceProject containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceProjectService) CreateOrUpdateResourceProject(request models.ResourceProject) (*models.ResourceProject, error) {
	var response models.ResourceProject
	err := s.baseService.DoPost("CreateOrUpdateResourceProject", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteResourceProject deletes a resource project based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource project.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceProjectService) DeleteResourceProject(id string) (interface{}, error) {
	var response models.ResourceProject
	return s.baseService.DoDelete("DeleteResourceProject", map[string]string{"id": id}, &response)
}
