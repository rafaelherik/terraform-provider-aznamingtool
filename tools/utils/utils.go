package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func GetStringMap(fromMap basetypes.MapValue) (attrMap map[string]string) {
	elements := fromMap.Elements()
	attrMap = make(map[string]string, len(elements))

	for key, value := range elements {
		attrMap[key] = value.(types.String).ValueString()
	}
}
