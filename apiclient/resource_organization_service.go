package apiclient

import "github.com/rafaelherik/azure-naming-provider/tf/apiclient/models"

type ResourceOrganizationService struct {
	baseService *BaseService
}

// NewResourceOrganizationService creates a new instance of ResourceOrganizationService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceOrganizationService instance.
func NewResourceOrganizationService(client *APIClient) *ResourceOrganizationService {
	return &ResourceOrganizationService{baseService: NewBaseService(client)}
}

// GetAllResourceOrganizations retrieves all resource organizations.
//
// Returns:
//   - A pointer to a slice of models.ResourceOrganization containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceOrganizationService) GetAllResourceOrganizations() (*[]models.ResourceOrganization, error) {
	var response []models.ResourceOrganization
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceOrganization retrieves a resource organization based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource organization.
//
// Returns:
//   - A pointer to models.ResourceOrganization containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceOrganizationService) GetResourceOrganization(id string) (*models.ResourceOrganization, error) {
	var response models.ResourceOrganization
	err := s.baseService.DoGet("GetResourceOrganization", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceOrganization creates or updates a resource organization based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceOrganization containing the request data.
//
// Returns:
//   - A pointer to models.ResourceOrganization containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceOrganizationService) CreateOrUpdateResourceOrganization(request models.ResourceOrganization) (*models.ResourceOrganization, error) {
	var response models.ResourceOrganization
	err := s.baseService.DoPost("CreateOrUpdateResourceOrganization", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteResourceOrganization deletes a resource organization based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource organization.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceOrganizationService) DeleteResourceOrganization(id string) (interface{}, error) {
	var response models.ResourceOrganization
	return s.baseService.DoDelete("DeleteResourceOrganization", map[string]string{"id": id}, &response)
}
