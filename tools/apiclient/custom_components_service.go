package apiclient

import (
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient/models"
)

type CustomComponentService struct {
	baseService *BaseService
}

// NewCustomComponentService creates a new instance of CustomComponentService with the provided API client.
//
// Parameters:
//   - client: A pointer to the APIClient instance.
//
// Returns:
//   - A pointer to the newly created CustomComponentService instance.
func NewCustomComponentService(client *APIClient) *CustomComponentService {
	return &CustomComponentService{baseService: NewBaseService(client)}
}

// GetAllCustomComponents retrieves all custom components.
//
// Returns:
//   - A pointer to a slice of models.CustomComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetAllCustomComponents() (*[]models.CustomComponent, error) {
	var response []models.CustomComponent
	err := s.baseService.DoGet("GetAllCustomComponents", nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCustomComponent retrieves a custom component based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the custom component.
//
// Returns:
//   - A pointer to models.CustomComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetCustomComponent(id string) (*models.CustomComponent, error) {
	var response models.CustomComponent
	err := s.baseService.DoGet("GetCustomComponent", map[string]string{"id": id}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCustomComponentByParentId retrieves a custom component based on the provided parent component ID.
//
// Parameters:
//   - parentComponentId: A string representing the ID of the parent custom component.
//
// Returns:
//   - A pointer to models.CustomComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetCustomComponentByParentId(parentComponentId string) (*models.CustomComponent, error) {
	var response models.CustomComponent
	err := s.baseService.DoGet("GetCustomComponentByParentId", map[string]string{"parentComponentId": parentComponentId}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCustomComponentByParentType retrieves custom components based on the provided parent type.
//
// Parameters:
//   - parentType: A string representing the type of the parent custom component.
//
// Returns:
//   - A pointer to a slice of models.CustomComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) GetCustomComponentByParentType(parentType string) (*[]models.CustomComponent, error) {
	var response []models.CustomComponent
	err := s.baseService.DoGet("GetCustomComponentByParentType", map[string]string{"parentType": parentType}, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateCustomComponent creates or updates a custom component based on the provided request data.
//
// Parameters:
//   - request: An instance of models.CustomComponent containing the request data.
//
// Returns:
//   - A pointer to models.CustomComponent containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) CreateOrUpdateCustomComponent(request models.CustomComponent) (*models.CustomComponent, error) {
	var response models.CustomComponent
	err := s.baseService.DoPost("CreateOrUpdateCustomComponent", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteCustomComponent deletes a custom component based on the provided ID.
//
// Parameters:
//   - id: A string representing the ID of the custom component.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) DeleteCustomComponent(id string) (interface{}, error) {
	var response models.ResourceOrganization
	return s.baseService.DoDelete("DeleteCustomComponent", map[string]string{"id": id}, &response)
}

// DeleteCustomComponentByParentId deletes custom components based on the provided parent component ID.
//
// Parameters:
//   - parentComponentId: A string representing the ID of the parent custom component.
//
// Returns:
//   - An interface containing the response data.
//   - An error if the request fails or the response indicates failure.
func (s *CustomComponentService) DeleteCustomComponentByParentId(parentComponentId string) (interface{}, error) {
	var response models.ResourceOrganization
	return s.baseService.DoDelete("DeleteCustomComponentByParentId", map[string]string{"parentComponentId": parentComponentId}, &response)
}
