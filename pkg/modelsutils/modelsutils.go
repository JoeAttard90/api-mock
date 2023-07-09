package modelsutils

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type SchemaInfo struct {
	Name   string
	Fields string
}

func GetSchemaInfo(doc *openapi3.T) []SchemaInfo {
	var schemas []SchemaInfo
	for schemaName, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value

		var fields []string
		for propName, propRef := range schema.Properties {
			goType := "any"
			prop := propRef.Value

			if prop != nil && prop.Items != nil {
				switch prop.Items.Value.Type {
				case "string":
					goType = "[]string"
				case "integer":
					goType = "[]int"
				case "boolean":
					goType = "[]bool"
				case "object":
					structFieldTypeSlice := strings.Split(prop.Items.Ref, "/")
					structFieldType := structFieldTypeSlice[len(structFieldTypeSlice)-1]
					goType = fmt.Sprintf("[]%s", structFieldType)
				}
				title := cases.Title(language.English, cases.NoLower)
				fieldName := title.String(propName)
				fields = append(fields, fmt.Sprintf("%s %s `json:\"%s\"`", fieldName, goType, propName))

				continue
			}
			if prop != nil {
				switch prop.Type {
				case "string":
					goType = "string"
				case "integer":
					goType = "int"
				case "boolean":
					goType = "bool"
				}
				title := cases.Title(language.English, cases.NoLower)
				fieldName := title.String(propName)
				fields = append(fields, fmt.Sprintf("%s %s `json:\"%s\"`", fieldName, goType, propName))
			}
		}

		fieldsStr := "struct {\n"
		for _, field := range fields {
			fieldsStr += "    " + field + "\n"
		}
		fieldsStr += "}"

		schemaInfo := SchemaInfo{
			Name:   schemaName,
			Fields: fieldsStr,
		}

		schemas = append(schemas, schemaInfo)
	}
	return schemas
}
