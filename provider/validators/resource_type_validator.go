package validators

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ validator.String = resourceTypeValidator{}
)

// Other methods to implement the attr.Value interface are omitted for brevity
type resourceTypeValidator struct {
	basetypes.StringValue
	availableTypes []string
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
	if v.IsNull() || v.IsUnknown() {
		return
	}

	if !v.isValid(v.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Resource Type Validation Error",
			fmt.Sprintf("Resource Type Validation Error: The informed value is not available, got: %s", v.ValueString()),
		)

		return
	}
}

func (v resourceTypeValidator) isValid(in string) bool {
	for _, valid := range v.availableTypes {
		if v.ValueString() == valid {
			return true
		}
	}
	return false
}
