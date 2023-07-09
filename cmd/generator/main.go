package main

import (
	"api-mock/pkg/handlerutils"
	"api-mock/pkg/modelsutils"
	"flag"
	"github.com/getkin/kin-openapi/openapi3"
	"log"
	"os"
)

func main() {
	handlerFuncTemplatePath := flag.String("handlerFuncTemplatePath", "templates/handlerFunc.tpl", "the path to the template for the handler funcs")
	handlersTemplatePath := flag.String("handlersTemplatePath", "./templates/handlers.tpl", "the path to the template for the handlers")
	handlersOutputPath := flag.String("handlersOutputPath", "./pkg/handlers/handlers.go", "the path to output the generated handler funcs")
	serverTemplatePath := flag.String("serverTemplatePath", "./templates/server.tpl", "the path to the template for the server")
	serverOutputPath := flag.String("serverOutputPath", "./cmd/mockserver/main.go", "the path to output the generated server")
	openAPISpecPath := flag.String("openAPISpecPath", "./exampledocs/scap_api_spec.yaml", "the path to the open api spec file")

	flag.Parse()

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(*openAPISpecPath)
	if err != nil {
		log.Printf("could not load from file: %q - %s", *openAPISpecPath, err.Error())
	}
	modelGenerator := modelsutils.NewModelGenerator(doc)
	handlerGenerator := handlerutils.NewHandlersGenerator(
		doc,
		*handlerFuncTemplatePath,
		*handlersTemplatePath,
		*handlersOutputPath,
		*serverTemplatePath,
		*serverOutputPath,
	)

	err = modelGenerator.GenerateModels()
	if err != nil {
		log.Printf("failed to generate models for spec: %q", *openAPISpecPath)
		os.Exit(1)
	}
	err = handlerGenerator.GenerateHandlers()
	if err != nil {
		log.Printf("failed to generate handlers for spec: %q", *openAPISpecPath)
		os.Exit(1)
	}

	log.Printf("successfully generated a mock server based on the spec: %q", *openAPISpecPath)
}
