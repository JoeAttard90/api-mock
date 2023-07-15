package templateutils

import (
	"api-mock/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func CreateTemplate(templatePath, outputPath string, data any) error {
	tpl := template.New(filepath.Base(templatePath)).Funcs(
		template.FuncMap{
			"parseEndpoint": ParseEndpoint,
			"pathToHandler": utils.PathToTitle,
			"toPascal":      utils.ToPascalCase,
		})
	tpl, err := tpl.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	outputDir := filepath.Dir(outputPath)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}
