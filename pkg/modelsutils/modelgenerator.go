package modelsutils

import (
	"api-mock/pkg/templateutils"
	"github.com/getkin/kin-openapi/openapi3"
)

type ModelGenerator struct {
	doc *openapi3.T
}

func NewModelGenerator(doc *openapi3.T) *ModelGenerator {
	return &ModelGenerator{
		doc: doc,
	}
}

func (mc *ModelGenerator) GenerateModels() error {
	schemas := GetSchemaInfo(mc.doc)
	err := templateutils.CreateTemplate(
		"templates/models.tpl",
		"./models/schemas.go",
		schemas,
	)
	if err != nil {
		return err
	}
	return nil
}
