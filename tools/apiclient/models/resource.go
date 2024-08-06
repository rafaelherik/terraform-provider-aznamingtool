package models

type ResourceBaseEntity struct {
	Id        int
	Name      string
	ShortName string
	SortOrder int
}

type ResourceDelimiter struct {
	Id        int
	Name      string
	Delimiter string
	Enabled   bool
	SortOrder int
}

type ResourceEnvironment struct {
	ResourceBaseEntity
}

type ResourceFunction struct {
	ResourceBaseEntity
}

type ResourceOrganization struct {
	ResourceBaseEntity
}

type ResourceProject struct {
	ResourceBaseEntity
}

type ResourceLocation struct {
	ResourceBaseEntity
}

type ResourceUnit struct {
	ResourceBaseEntity
}

type ResourceType struct {
	Id                           int64
	Resource                     string
	Optional                     string
	Exclude                      string
	Property                     string
	ShortName                    string
	Scope                        string
	LenghtMin                    string
	LenghtMax                    string
	ValidText                    string
	InvalidText                  string
	InvalidCharacters            string
	InvalidCharactersStart       string
	InvalidCharactersEnd         string
	InvalidCharactersConsecutive string
	Regx                         string
	StaticValues                 string
	Enabled                      bool
	ApplyDelimiter               bool
}

type ResourceNameRequest struct {
	ResourceEnvironment string
	ResourceFunction    string
	ResourceInstance    string
	ResourceLocation    string
	ResourceOrg         string
	ResourceProjAppSvc  string
	ResourceType        string
	ResourceUnitDept    string
	CustomComponents    []CustomComponent
	ResourceId          int64
	CreatedBy           string
}

type ResourceNameRequestWithComponents struct {
	ResourceEnvironment ResourceEnvironment
	ResourceFunction    ResourceFunction
	ResourceDelimiter   ResourceDelimiter
	ResourceInstance    string
	ResourceLocation    ResourceLocation
	ResourceOrg         ResourceOrganization
	ResourceProjAppSvc  ResourceProject
	ResourceType        ResourceType
	ResourceUnitDept    ResourceUnit
}

type ResourceGeneratedName struct {
	Id               int64
	CreatedOn        string
	ResourceName     string
	ResourceTypeName string
	User             string
	Message          string
}

type ResourceNameResponse struct {
	ResourceName        string
	Message             string
	Success             bool
	ResourceNameDetails ResourceGeneratedName
}

type ResourceComponent struct {
	Id                   int64
	Name                 string
	DisplayName          string
	Enabled              bool
	SortOrder            int
	IsCustom             bool
	IsFreeText           bool
	MinLength            string
	MaxLength            string
	EnforceRandom        bool
	Alphanumeric         bool
	ApplyDelimiterBefore bool
	ApplyDelimiterAfter  bool
}
