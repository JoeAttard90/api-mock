package main

import (
	"api-mock/pkg/appbuilder"
	"api-mock/pkg/dockerutils"
	"api-mock/pkg/handlerutils"
	"api-mock/pkg/structutils"
	"flag"
	"github.com/getkin/kin-openapi/openapi3"
	"log"
	"os"
)

func main() {
	// Source openapi spec
	openAPISpecPath := flag.String("openAPISpecPath", "./exampledocs/openapi_pet_store.yaml", "the path to the open api spec file")

	// Server and Handler generator flags
	handlerFuncTemplatePath := flag.String("handlerFuncTemplatePath", "templates/handlerFunc.tpl", "the path to the template for the handler funcs")
	handlersTemplatePath := flag.String("handlersTemplatePath", "./templates/handlers.tpl", "the path to the template for the handlers")
	handlersOutputPath := flag.String("handlersOutputPath", "../api-mock-server/pkg/handlers/handlers.go", "the path to output the generated handler funcs")
	serverTemplatePath := flag.String("serverTemplatePath", "./templates/server.tpl", "the path to the template for the server")
	serverOutputPath := flag.String("serverOutputPath", "../api-mock-server/cmd/server/main.go", "the path to output the generated server")
	staticResponses := flag.String("staticResponses", "", "the path to static responses (named the same as the endpoints) to be used for specified endpoints")

	// Structs generator flags
	structsTemplatePath := flag.String("structsTemplatePath", "templates/structs.tpl", "the path to the template for the structs")
	structsOutputPath := flag.String("structsOutputPath", "../api-mock-server/pkg/structs/schemas.go", "the path to output the generated structs")

	// Docker generator flags
	dockerfileTemplatePath := flag.String("dockerfileTemplatePath", "./templates/dockerfile.tpl", "the path to the template for the dockerfile")
	dockerfileOutputPath := flag.String("dockerfileOutputPath", "../api-mock-server/Dockerfile", "the path to output the generated handler funcs")
	dockerComposeTemplatePath := flag.String("dockerComposeTemplatePath", "./templates/dockerCompose.tpl", "the path to the template for the dockerCompose")
	dockerComposeOutputPath := flag.String("dockerComposeOutputPath", "../api-mock-server/docker-compose.yml", "the path to output the generated dockerCompose")
	mockAPIPort := flag.String("mockAPIPort", "8080", "the port on which to expose the generated mock API")

	// Builder
	dirPath := flag.String("dirPath", "..", "the directory in which to generate the server")
	modName := flag.String("modName", "api-mock-server", "the name of the go module being generated")

	flag.Parse()

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(*openAPISpecPath)
	if err != nil {
		log.Printf("could not load from file: %q - %s", *openAPISpecPath, err.Error())
	}
	structsGenerator := structutils.NewStructsGenerato(
		doc,
		*structsTemplatePath,
		*structsOutputPath,
	)
	handlerGenerator := handlerutils.NewHandlersGenerator(
		doc,
		*handlerFuncTemplatePath,
		*handlersTemplatePath,
		*handlersOutputPath,
		*serverTemplatePath,
		*serverOutputPath,
		*staticResponses,
	)
	dockerGenerator := dockerutils.NewDockerFileGenerator(
		*mockAPIPort,
		*dockerfileTemplatePath,
		*dockerfileOutputPath,
		*dockerComposeTemplatePath,
		*dockerComposeOutputPath,
	)
	builder := appbuilder.NewBuilder(
		*dirPath,
		*modName,
		*staticResponses,
	)

	err = structsGenerator.Generate()
	if err != nil {
		log.Printf("failed to generate structs for spec: %q :%s", *openAPISpecPath, err.Error())
		os.Exit(1)
	}
	err = handlerGenerator.Generate()
	if err != nil {
		log.Printf("failed to generate handlers for spec: %q :%s", *openAPISpecPath, err.Error())
		os.Exit(1)
	}

	err = dockerGenerator.Generate()
	if err != nil {
		log.Printf("failed to generate docker files: %s", err.Error())
		os.Exit(1)
	}

	// TODO: update the execute command to first check if the module exists, if so delete all then re-build
	dir, err := builder.ExecuteCommands()
	if err != nil {
		log.Printf("failed to execute commands during generation: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("successfully generated a mock server based on the spec: %q, saved in location: %q", *openAPISpecPath, dir)
	log.Printf("mock server running in docker container on port %q", *mockAPIPort)
}
