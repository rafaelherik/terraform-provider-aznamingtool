package apiclient

import "github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient/models"

type ResourceDelimiterService struct {
	baseService *BaseService
}

// NewResourceDelimiterService creates a new instance of ResourceDelimiterService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created ResourceDelimiterService instance.
func NewResourceDelimiterService(client *APIClient) *ResourceDelimiterService {
	return &ResourceDelimiterService{baseService: NewBaseService(client)}
}

// GetAllResourceDelimiters retrieves all resource delimiters.
//
// Returns:
//   - A pointer to a slice of models.ResourceDelimiter containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceDelimiterService) GetAllResourceDelimiters() (*[]models.ResourceDelimiter, error) {
	var response []models.ResourceDelimiter
	err := s.baseService.DoGet("GetAllResourceDelimiters", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetResourceDelimiter retrieves a resource delimiter based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the resource delimiter.
//
// Returns:
//   - A pointer to models.ResourceDelimiter containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceDelimiterService) GetResourceDelimiter(id string) (*models.ResourceDelimiter, error) {
	var response models.ResourceDelimiter
	err := s.baseService.DoGet("GetResourceDelimiter", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateResourceDelimiter creates or updates a resource delimiter based on the provided request data.
//
// Parameters:
//   - request: An instance of models.ResourceDelimiter containing the request data.
//
// Returns:
//   - A pointer to models.ResourceDelimiter containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *ResourceDelimiterService) CreateOrUpdateResourceDelimiter(request models.ResourceDelimiter) (*models.ResourceDelimiter, error) {
	var response models.ResourceDelimiter
	err := s.baseService.DoPost("CreateOrUpdateResourceDelimiter", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
