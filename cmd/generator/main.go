package main

import (
	"api-mock/pkg/appbuilder"
	"api-mock/pkg/handlerutils"
	"api-mock/pkg/modelsutils"
	"flag"
	"github.com/getkin/kin-openapi/openapi3"
	"log"
	"os"
)

func main() {
	// Source openapi spec
	openAPISpecPath := flag.String("openAPISpecPath", "./exampledocs/scap_api_spec.yaml", "the path to the open api spec file")

	// Server and Handler
	handlerFuncTemplatePath := flag.String("handlerFuncTemplatePath", "templates/handlerFunc.tpl", "the path to the template for the handler funcs")
	handlersTemplatePath := flag.String("handlersTemplatePath", "./templates/handlers.tpl", "the path to the template for the handlers")
	handlersOutputPath := flag.String("handlersOutputPath", "../api-mock-server/pkg/handlers/handlers.go", "the path to output the generated handler funcs")
	serverTemplatePath := flag.String("serverTemplatePath", "./templates/server.tpl", "the path to the template for the server")
	serverOutputPath := flag.String("serverOutputPath", "../api-mock-server/cmd/server/main.go", "the path to output the generated server")

	// Models
	modelsTemplatePath := flag.String("modelsTemplatePath", "templates/models.tpl", "the path to the template for the models")
	modelsOutputPath := flag.String("modelsOutputPath", "../api-mock-server/pkg/models/schemas.go", "the path to output the generated models")

	// Builder
	dirPath := flag.String("dirPath", "..", "the directory in which to generate the server")
	modName := flag.String("modName", "api-mock-server", "the name of the go module being generated")

	flag.Parse()

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(*openAPISpecPath)
	if err != nil {
		log.Printf("could not load from file: %q - %s", *openAPISpecPath, err.Error())
	}
	modelGenerator := modelsutils.NewModelGenerator(
		doc,
		*modelsTemplatePath,
		*modelsOutputPath,
	)
	handlerGenerator := handlerutils.NewHandlersGenerator(
		doc,
		*handlerFuncTemplatePath,
		*handlersTemplatePath,
		*handlersOutputPath,
		*serverTemplatePath,
		*serverOutputPath,
	)
	builder := appbuilder.NewBuilder(
		*dirPath,
		*modName,
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

	// TODO: update the execute command to first check if the module exists, if so delete all then re-build
	dir, err := builder.ExecuteCommands()
	if err != nil {
		log.Printf("failed to execute commands during generation: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("successfully generated a mock server based on the spec: %q", *openAPISpecPath)
	log.Printf("mock server location: %q", dir)
}
