package apiclient

import "github.com/rafaelherik/azure-naming-provider/tf/apiclient/models"

type ResourceLocationService struct {
	baseService *BaseService
}

// NewResourceLocationService creates a new instance of ResourceLocationService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceLocationService instance.
func NewResourceLocationService(client *APIClient) *ResourceLocationService {
	return &ResourceLocationService{baseService: NewBaseService(client)}
}

// GetAllResourceLocations retrieves all resource locations.
//
// Returns:
//   - A pointer to a slice of models.ResourceLocation containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceLocationService) GetAllResourceLocations() (*[]models.ResourceLocation, error) {
	var response []models.ResourceLocation
	err := s.baseService.DoGet("RequestName", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceLocation retrieves a resource location based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource location.
//
// Returns:
//   - A pointer to models.ResourceLocation containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceLocationService) GetResourceLocation(id string) (*models.ResourceLocation, error) {
	var response models.ResourceLocation
	err := s.baseService.DoGet("GetResourceLocation", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceLocation creates or updates a resource location based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceLocation containing the request data.
//
// Returns:
//   - A pointer to models.ResourceLocation containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceLocationService) CreateOrUpdateResourceLocation(request models.ResourceLocation) (*models.ResourceLocation, error) {
	var response models.ResourceLocation
	err := s.baseService.DoPost("CreateOrUpdateResourceLocation", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteResourceLocation deletes a resource location based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource location.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceLocationService) DeleteResourceLocation(id string) (interface{}, error) {
	var response models.ResourceLocation
	return s.baseService.DoDelete("DeleteResourceLocation", map[string]string{"id": id}, &response)
}
