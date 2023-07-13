package structutils

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type SchemaInfo struct {
	Name         string
	Abbreviation string
	Struct       string
	Fields       []Field
}

type Field struct {
	FieldName    string
	Type         string
	IsCustomType bool
	IsSlice      bool
}

func GetSchemaInfo(doc *openapi3.T) []SchemaInfo {
	var schemas []SchemaInfo

	for schemaName, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value
		var isCustomType bool
		var fieldDetails []Field

		var fields []string
		for propName, propRef := range schema.Properties {
			goType := "any"
			prop := propRef.Value
			var field Field

			// set to false unless it's an object => it is a custom type
			isCustomType = false
			if prop != nil && prop.Items != nil {
				field.IsSlice = true
				switch prop.Items.Value.Type {
				case "string":
					goType = "[]string"
				case "integer":
					goType = "[]int"
				case "boolean":
					goType = "[]bool"
				case "number":
					goType = "[]float32"
				case "object":
					structFieldTypeSlice := strings.Split(prop.Items.Ref, "/")
					structFieldType := structFieldTypeSlice[len(structFieldTypeSlice)-1]
					goType = fmt.Sprintf("[]%s", structFieldType)
					isCustomType = true
				}
				fieldName := cases.Title(language.English, cases.NoLower).String(propName)
				fields = append(fields, fmt.Sprintf("%s %s `json:\"%s\"`", fieldName, goType, propName))
				field.FieldName = fieldName
				field.Type = goType[2:]
				field.IsCustomType = isCustomType
				fieldDetails = append(fieldDetails, field)

				continue
			}
			if prop != nil {
				field.IsSlice = false
				switch prop.Type {
				case "string":
					goType = "string"
				case "integer":
					goType = "int"
				case "boolean":
					goType = "bool"
				case "number":
					goType = "float32"
				case "object":
					structFieldTypeSlice := strings.Split(propRef.Ref, "/")
					structFieldType := structFieldTypeSlice[len(structFieldTypeSlice)-1]
					goType = fmt.Sprintf("%s", structFieldType)
					isCustomType = true
				}
				fieldName := cases.Title(language.English, cases.NoLower).String(propName)
				fields = append(fields, fmt.Sprintf("%s %s `json:\"%s\"`", fieldName, goType, propName))
				field.FieldName = fieldName
				field.Type = goType
				field.IsCustomType = isCustomType
				fieldDetails = append(fieldDetails, field)
			}
		}

		fieldsStr := "struct {\n"
		for _, field := range fields {
			fieldsStr += "    " + field + "\n"
		}
		fieldsStr += "}"

		schemaInfo := SchemaInfo{
			Name:         schemaName,
			Abbreviation: strings.ToLower(string(schemaName[0])),
			Struct:       fieldsStr,
			Fields:       fieldDetails,
		}

		schemas = append(schemas, schemaInfo)
	}
	return schemas
}
