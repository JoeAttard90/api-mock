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

type HandlerInfo struct {
	Path    string
	Method  string
	Handler string
}

func main() {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./exampledocs/scap_api_spec.yaml")
	if err != nil {
		log.Fatalf("Could not load spec: %v", err)
	}

	var handlers []HandlerInfo

	tplFunc, err := template.ParseFiles("templates/handlerFunc.tpl")
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}

	for path, pathItem := range doc.Paths {
		for method, _ := range pathItem.Operations() {
			handlerName := fmt.Sprintf("%s%s", cases.Title(language.English, cases.NoLower).String(method), cases.Title(language.English, cases.NoLower).String(path))
			handlerName = cases.Title(language.English, cases.NoLower).String(handlerName)
			path = strings.Replace(path, "/", "", -1)

			handlerInfo := HandlerInfo{
				Path:   path,
				Method: method,
			}

			var handlerBuilder strings.Builder
			err = tplFunc.Execute(&handlerBuilder, handlerInfo)
			if err != nil {
				log.Fatalf("Error executing template: %v", err)
			}
			handler := handlerBuilder.String()

			handlerInfo.Handler = handler

			handlers = append(handlers, handlerInfo)
		}
	}

	tpl, err := template.ParseFiles("./templates/handlers.tpl")

	outputPath := "./pkg/handlers/handlers.go"
	outputDir := filepath.Dir(outputPath)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer f.Close()

	err = tpl.Execute(f, handlers)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
