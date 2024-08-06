package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient/models"
)

var (
	_ validator.String = resourceTypeValidator{}
)

// Other methods to implement the attr.Value interface are omitted for brevity
type resourceTypeValidator struct {
	availableTypes []string
	client         *apiclient.APIClient
}

// Description implements validator.String.
func (v resourceTypeValidator) Description(context.Context) string {
	return fmt.Sprintf("Resource type must be one of the available types: %s", strings.Join(v.availableTypes, ","))
}

// MarkdownDescription implements validator.String.
func (v resourceTypeValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf("Resource type must be one of the available types: %s", strings.Join(v.availableTypes, ","))
}

// Implementation of the function.ValidateableParameter interface
func (v resourceTypeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if !v.isValid(req.ConfigValue.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Resource Type Validation Error",
			fmt.Sprintf("Resource Type Validation Error: The informed value is not available, got: %s. Available Values are:%s.", req.ConfigValue.ValueString(), strings.Join(v.availableTypes, ",")),
		)

		return
	}
}

func (v resourceTypeValidator) isValid(in string) bool {

	if v.availableTypes == nil {
		fmt.Println(v.client)
		svc := apiclient.NewResourceTypeService(v.client)
		result, err := svc.GetAllResourceTypes()

		if err != nil {
			fmt.Println("Error getting resource types: ", err)
		}

		validResourceTypeNames := func(resourceType *[]models.ResourceType) []string {
			resources := make([]string, len(*resourceType))
			for i, resourceType := range *resourceType {
				resources[i] = resourceType.Resource
			}
			return resources
		}(result)

		v.availableTypes = validResourceTypeNames
	}

	for _, valid := range v.availableTypes {
		if in == valid {
			return true
		}
	}
	return false
}
