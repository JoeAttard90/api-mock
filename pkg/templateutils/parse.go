package templateutils

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func CreateTemplate(templatePath, outputPath string, data any) error {
	tpl, err := template.ParseFiles(templatePath)
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
