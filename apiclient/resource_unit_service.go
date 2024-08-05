package apiclient

import "github.com/rafaelherik/azure-naming-provider/tf/apiclient/models"

type ResourceUnitService struct {
	baseService *BaseService
}

// NewResourceUnitService creates a new instance of ResourceUnitService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceUnitService instance.
func NewResourceUnitService(client *APIClient) *ResourceUnitService {
	return &ResourceUnitService{baseService: NewBaseService(client)}
}

// GetAllResourceUnits retrieves all resource units.
//
// Returns:
//   - A pointer to a slice of models.ResourceUnit containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceUnitService) GetAllResourceUnits() (*[]models.ResourceUnit, error) {
	var response []models.ResourceUnit
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceUnit retrieves a resource unit based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource unit.
//
// Returns:
//   - A pointer to models.ResourceUnit containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceUnitService) GetResourceUnit(id string) (*models.ResourceUnit, error) {
	var response models.ResourceUnit
	err := s.baseService.DoGet("GetResourceUnit", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceUnit creates or updates a resource unit based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceUnit containing the request data.
//
// Returns:
//   - A pointer to models.ResourceUnit containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceUnitService) CreateOrUpdateResourceUnit(request models.ResourceUnit) (*models.ResourceUnit, error) {
	var response models.ResourceUnit
	err := s.baseService.DoPost("CreateOrUpdateResourceUnit", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteResourceUnit deletes a resource unit based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource unit.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceUnitService) DeleteResourceUnit(id string) (interface{}, error) {
	var response models.ResourceUnit
	return s.baseService.DoDelete("DeleteResourceUnit", map[string]string{"id": id}, &response)
}
