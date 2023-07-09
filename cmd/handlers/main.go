package main

import (
	"api-mock/pkg/handlerutils"
	"api-mock/pkg/utils"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type Handlers struct {
	Endpoints            map[string]string
	HandlersInfo         []handlerutils.HandlerInfo
	HasPost              bool
	GlobalSecurityScheme string
}

func main() {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./exampledocs/openapi_pet_store.yaml")
	if err != nil {
		log.Fatalf("Could not load spec: %v", err)
	}

	var handlers Handlers
	handlers.Endpoints = make(map[string]string)

	tplFunc, err := template.ParseFiles("templates/handlerFunc.tpl")
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}

	for path, pathItem := range doc.Paths {
		var securitySchemes openapi3.SecuritySchemes
		var secSchemes []string

		securitySchemes = doc.Components.SecuritySchemes
		for _, scheme := range securitySchemes {
			secSchemes = append(secSchemes, scheme.Value.Scheme)
		}

		if len(secSchemes) == 1 {
			handlers.GlobalSecurityScheme = cases.Title(language.English, cases.NoLower).String(secSchemes[0])
		}

		for method := range pathItem.Operations() {
			handlerName := utils.PathToTitle(path)
			handlerFunction := utils.ToPascalCase(method)
			if !strings.HasPrefix(handlerName, handlerFunction) {
				handlerName = fmt.Sprintf("%s%s", handlerFunction, handlerName)
			}

			var handlerInfo handlerutils.HandlerInfo
			handlerInfo.QueryParams = make(map[string]string)
			handlerInfo.Slugs = utils.ExtractSlugs(path)

			handlers.Endpoints[path] = handlerName

			var reqBodyContent openapi3.Content
			var respBodyContent openapi3.Content
			switch method {
			case http.MethodPost:
				reqBody := pathItem.Post.RequestBody
				if reqBody != nil {
					reqBodyContent = reqBody.Value.Content
					for k := range reqBodyContent {
						if !utils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Post.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !utils.Contains(handlerInfo.RespMimeTypes, k) {
							handlerInfo.RespMimeTypes = append(handlerInfo.RespMimeTypes, k)
						}
					}
				}

				handlerInfo.SetReqRespTypes(reqBodyContent, respBodyContent)

				queryParameters := pathItem.Post.Parameters
				handlerInfo.SetQueryParams(queryParameters)
				handlers.HasPost = true
			case http.MethodGet:
				reqBody := pathItem.Get.RequestBody
				if reqBody != nil {
					reqBodyContent = reqBody.Value.Content
					for k := range reqBodyContent {
						if !utils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Get.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !utils.Contains(handlerInfo.RespMimeTypes, k) {
							handlerInfo.RespMimeTypes = append(handlerInfo.RespMimeTypes, k)
						}
					}
				}

				handlerInfo.SetReqRespTypes(reqBodyContent, respBodyContent)

				queryParameters := pathItem.Get.Parameters
				handlerInfo.SetQueryParams(queryParameters)
			case http.MethodPut:
				reqBody := pathItem.Put.RequestBody
				if reqBody != nil {
					reqBodyContent = reqBody.Value.Content
					for k := range reqBodyContent {
						if !utils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Put.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !utils.Contains(handlerInfo.RespMimeTypes, k) {
							handlerInfo.RespMimeTypes = append(handlerInfo.RespMimeTypes, k)
						}
					}
				}

				handlerInfo.SetReqRespTypes(reqBodyContent, respBodyContent)

				queryParameters := pathItem.Put.Parameters
				handlerInfo.SetQueryParams(queryParameters)
				handlers.HasPost = true
			case http.MethodDelete:
				reqBody := pathItem.Delete.RequestBody
				if reqBody != nil {
					reqBodyContent = reqBody.Value.Content
					for k := range reqBodyContent {
						if !utils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Delete.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !utils.Contains(handlerInfo.RespMimeTypes, k) {
							handlerInfo.RespMimeTypes = append(handlerInfo.RespMimeTypes, k)
						}
					}
				}

				handlerInfo.SetReqRespTypes(reqBodyContent, respBodyContent)

				queryParameters := pathItem.Delete.Parameters
				handlerInfo.SetQueryParams(queryParameters)
			}
			handlerInfo.Path = handlerName
			handlerInfo.Method = method
			//TODO: we should check the individual endpoints for security overrides
			handlerInfo.SecurityScheme = handlers.GlobalSecurityScheme

			var handlerBuilder strings.Builder
			err = tplFunc.Execute(&handlerBuilder, handlerInfo)
			if err != nil {
				log.Fatalf("Error executing template: %v", err)
			}
			handler := handlerBuilder.String()

			handlerInfo.Handler = handler

			handlers.HandlersInfo = append(handlers.HandlersInfo, handlerInfo)
		}
	}

	// Create Handlers Package
	tpl, err := template.ParseFiles("./templates/handlers.tpl")

	outputPath := "./pkg/handlers/handlers.go"
	outputDir := filepath.Dir(outputPath)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating handlers pkg file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	err = tpl.Execute(f, handlers)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	//Create server main
	serverTpl, err := template.ParseFiles("./templates/server.tpl")

	serverOutputPath := "./cmd/mockserver/main.go"
	serverOutputDir := filepath.Dir(serverOutputPath)

	if err := os.MkdirAll(serverOutputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	file, err := os.Create(serverOutputPath)
	if err != nil {
		log.Fatalf("Error creating server main file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	err = serverTpl.Execute(file, handlers)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
