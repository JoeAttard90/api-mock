package modelsutils

import (
	"api-mock/pkg/templateutils"
	"github.com/getkin/kin-openapi/openapi3"
)

type ModelGenerator struct {
	doc          *openapi3.T
	templatePath string
	outputPath   string
}

func NewModelGenerator(doc *openapi3.T, templatePath, outputPath string) *ModelGenerator {
	return &ModelGenerator{
		doc:          doc,
		templatePath: templatePath,
		outputPath:   outputPath,
	}
}

func (mc *ModelGenerator) GenerateModels() error {
	schemas := GetSchemaInfo(mc.doc)
	err := templateutils.CreateTemplate(
		mc.templatePath,
		mc.outputPath,
		schemas,
	)
	if err != nil {
		return err
	}
	return nil
}
