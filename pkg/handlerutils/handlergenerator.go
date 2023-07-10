package handlerutils

import (
	"api-mock/pkg/templateutils"
	"api-mock/pkg/utils"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strings"
	"text/template"
)

type HandlersGenerator struct {
	Endpoints               map[string]string
	HandlersInfo            []HandlerInfo
	HasPost                 bool
	HasSlug                 bool
	GlobalSecurityScheme    string
	doc                     *openapi3.T
	handlerFuncTemplatePath string
	handlerFuncOutputPath   string
	handlersTemplatePath    string
	handlersOutputPath      string
	serverTemplatePath      string
	serverOutputPath        string
}

func NewHandlersGenerator(
	doc *openapi3.T,
	handlerFuncTemplatePath,
	handlersTemplatePath,
	handlersOutputPath,
	serverTemplatePath,
	serverOutputPath string,
) *HandlersGenerator {
	endpointsMap := make(map[string]string)
	return &HandlersGenerator{
		Endpoints:               endpointsMap,
		doc:                     doc,
		handlerFuncTemplatePath: handlerFuncTemplatePath,
		handlersTemplatePath:    handlersTemplatePath,
		handlersOutputPath:      handlersOutputPath,
		serverTemplatePath:      serverTemplatePath,
		serverOutputPath:        serverOutputPath,
	}
}

func (hg *HandlersGenerator) GenerateHandlers() error {
	for path, pathItem := range hg.doc.Paths {
		var securitySchemes openapi3.SecuritySchemes
		var secSchemes []string

		securitySchemes = hg.doc.Components.SecuritySchemes
		for _, scheme := range securitySchemes {
			secSchemes = append(secSchemes, scheme.Value.Scheme)
		}

		if len(secSchemes) == 1 {
			hg.GlobalSecurityScheme = cases.Title(language.English, cases.NoLower).String(secSchemes[0])
		}

		for method := range pathItem.Operations() {
			handlerName := utils.PathToTitle(path)
			handlerFunction := utils.ToPascalCase(method)
			if !strings.HasPrefix(handlerName, handlerFunction) {
				handlerName = fmt.Sprintf("%s%s", handlerFunction, handlerName)
			}

			var handlerInfo HandlerInfo
			handlerInfo.QueryParams = make(map[string]string)
			handlerInfo.Slugs = utils.ExtractSlugs(path)
			if len(handlerInfo.Slugs) > 0 {
				hg.HasSlug = true
			}

			hg.Endpoints[path] = handlerName

			var reqBodyContent openapi3.Content
			var respBodyContent openapi3.Content
			switch method {
			case http.MethodPost:
				reqBody := pathItem.Post.RequestBody
				if reqBody != nil {
					reqBodyContent = reqBody.Value.Content
					for k := range reqBodyContent {
						if !templateutils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Post.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !templateutils.Contains(handlerInfo.RespMimeTypes, k) {
							handlerInfo.RespMimeTypes = append(handlerInfo.RespMimeTypes, k)
						}
					}
				}

				handlerInfo.SetReqRespTypes(reqBodyContent, respBodyContent)

				queryParameters := pathItem.Post.Parameters
				handlerInfo.SetQueryParams(queryParameters)
				hg.HasPost = true
			case http.MethodGet:
				reqBody := pathItem.Get.RequestBody
				if reqBody != nil {
					reqBodyContent = reqBody.Value.Content
					for k := range reqBodyContent {
						if !templateutils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Get.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !templateutils.Contains(handlerInfo.RespMimeTypes, k) {
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
						if !templateutils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Put.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !templateutils.Contains(handlerInfo.RespMimeTypes, k) {
							handlerInfo.RespMimeTypes = append(handlerInfo.RespMimeTypes, k)
						}
					}
				}

				handlerInfo.SetReqRespTypes(reqBodyContent, respBodyContent)

				queryParameters := pathItem.Put.Parameters
				handlerInfo.SetQueryParams(queryParameters)
				hg.HasPost = true
			case http.MethodDelete:
				reqBody := pathItem.Delete.RequestBody
				if reqBody != nil {
					reqBodyContent = reqBody.Value.Content
					for k := range reqBodyContent {
						if !templateutils.Contains(handlerInfo.ReqMimeTypes, k) {
							handlerInfo.ReqMimeTypes = append(handlerInfo.ReqMimeTypes, k)
						}
					}
				}

				respOkBody := pathItem.Delete.Responses.Get(http.StatusOK)
				if respOkBody != nil {
					respBodyContent = respOkBody.Value.Content
					for k := range respBodyContent {
						if !templateutils.Contains(handlerInfo.RespMimeTypes, k) {
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
			handlerInfo.SecurityScheme = hg.GlobalSecurityScheme

			// Build each handler func
			var handlerBuilder strings.Builder
			tplFunc, err := template.ParseFiles(hg.handlerFuncTemplatePath)
			if err != nil {
				return err
			}

			err = tplFunc.Execute(&handlerBuilder, handlerInfo)
			if err != nil {
				return err
			}
			handler := handlerBuilder.String()

			handlerInfo.Handler = handler

			hg.HandlersInfo = append(hg.HandlersInfo, handlerInfo)
		}
	}

	// Create Handlers with all handler funcs
	err := templateutils.CreateTemplate(
		hg.handlersTemplatePath,
		hg.handlersOutputPath,
		hg,
	)
	if err != nil {
		return err
	}
	// Create Server
	err = templateutils.CreateTemplate(
		hg.serverTemplatePath,
		hg.serverOutputPath,
		hg,
	)
	if err != nil {
		return err
	}

	return nil
}
