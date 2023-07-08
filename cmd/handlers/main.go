package main

import (
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

type HandlerInfo struct {
	Path        string
	Method      string
	Handler     string
	ReqType     string
	ReqTypeVar  string
	QueryParams []string
}

type Handlers struct {
	HandlersInfo []HandlerInfo
	HasPost      bool
}

func main() {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./exampledocs/scap_api_spec.yaml")
	if err != nil {
		log.Fatalf("Could not load spec: %v", err)
	}

	var handlers Handlers

	tplFunc, err := template.ParseFiles("templates/handlerFunc.tpl")
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}

	for path, pathItem := range doc.Paths {
		for method, _ := range pathItem.Operations() {
			handlerName := fmt.Sprintf("%s%s", cases.Title(language.English, cases.NoLower).String(method), cases.Title(language.English, cases.NoLower).String(path))
			handlerName = cases.Title(language.English, cases.NoLower).String(handlerName)
			cleanPath := strings.Replace(path, "/", "", -1)

			var handlerInfo HandlerInfo
			// TODO: Here if we know it's a post, we just need to add logic to a template to handle posts
			// => unmarshalling to a known object (model) we have already extracted
			var componentStatusOk *openapi3.MediaType

			switch method {
			case http.MethodPost:
				componentStatusOk = pathItem.Post.RequestBody.Value.Content.Get("application/json")
				handlers.HasPost = true
			case http.MethodGet:
				componentStatusOk = pathItem.Get.Responses.Get(http.StatusOK).Value.Content.Get("application/json")
				for _, queryParam := range pathItem.Get.Parameters {
					handlerInfo.QueryParams = append(handlerInfo.QueryParams, queryParam.Value.Name)
				}
			}

			pathStatusOkSchema := strings.Split(componentStatusOk.Schema.Ref, "/")

			reqType := pathStatusOkSchema[len(pathStatusOkSchema)-1]
			reqTypeVar := utils.ToCamelCase(reqType)
			handlerInfo.ReqType = reqType
			handlerInfo.ReqTypeVar = reqTypeVar

			handlerInfo.Path = cleanPath
			handlerInfo.Method = method

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
