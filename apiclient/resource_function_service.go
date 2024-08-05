package apiclient

import "github.com/rafaelherik/azure-naming-provider/tf/apiclient/models"

type ResourceFunctionService struct {
	baseService *BaseService
}

// NewResourceFunctionService creates a new instance of ResourceFunctionService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceFunctionService instance.
func NewResourceFunctionService(client *APIClient) *ResourceFunctionService {
	return &ResourceFunctionService{baseService: NewBaseService(client)}
}

// GetAllResourceFunctions retrieves all resource functions.
//
// Returns:
//   - A pointer to a slice of models.ResourceFunction containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceFunctionService) GetAllResourceFunctions() (*[]models.ResourceFunction, error) {
	var response []models.ResourceFunction
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceFunction retrieves a resource function based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource function.
//
// Returns:
//   - A pointer to models.ResourceFunction containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceFunctionService) GetResourceFunction(id string) (*models.ResourceFunction, error) {
	var response models.ResourceFunction
	err := s.baseService.DoGet("GetResourceFunction", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceFunction creates or updates a resource function based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceFunction containing the request data.
//
// Returns:
//   - A pointer to models.ResourceFunction containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceFunctionService) CreateOrUpdateResourceFunction(request models.ResourceFunction) (*models.ResourceFunction, error) {
	var response models.ResourceFunction
	err := s.baseService.DoPost("CreateOrUpdateResourceFunction", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteResourceFunction deletes a resource function based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource function.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceFunctionService) DeleteResourceFunction(id string) (interface{}, error) {
	var response models.ResourceFunction
	return s.baseService.DoDelete("DeleteResourceFunction", map[string]string{"id": id}, &response)
}
