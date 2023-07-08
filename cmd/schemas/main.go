package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaInfo struct {
	Name   string
	Fields string
}

func main() {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./exampledocs/scap_api_spec.yaml")
	if err != nil {
		log.Fatalf("Could not load spec: %v", err)
	}

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

	const structTpl = `
package models

{{ range . }}
type {{ .Name }} {{ .Fields }}
{{ end }}
`
	tpl, err := template.New("structs").Parse(structTpl)
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}

	outputPath := "./models/schemas.go"
	outputDir := filepath.Dir(outputPath)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer f.Close()

	err = tpl.Execute(f, schemas)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
