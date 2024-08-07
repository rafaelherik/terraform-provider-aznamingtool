package utils

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func GetStringMap(fromMap basetypes.MapValue) (attrMap map[string]string) {
	elements := fromMap.Elements()
	attrMap = make(map[string]string, len(elements))

	for key, value := range elements {
		attrMap[key] = value.(types.String).ValueString()
	}
	return attrMap
}

func GetMapFromMatrix(components [][]string) map[string]string {
	result := make(map[string]string)
	for _, component := range components {
		if len(component) == 2 {
			result[component[0]] = component[1]
		}
	}
	return result
}

func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

func CamelToSnake(s string) string {
	var re = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}
