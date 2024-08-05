package models

type ValidateNameRequest struct {
	ResourceTypeId int
	ResourceType   string
	Name           string
}

type ValidateNameResponse struct {
	Valid   bool
	Name    string
	Message string
}
