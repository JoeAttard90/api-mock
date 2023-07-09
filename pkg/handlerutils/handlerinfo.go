package handlerutils

import (
	"api-mock/pkg/utils"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

type HandlerInfo struct {
	Path           string
	Method         string
	Handler        string
	ReqType        string
	ReqTypeVar     string
	RespType       string
	RespTypeVar    string
	QueryParams    map[string]string
	SecurityScheme string
	ReqMimeTypes   []string
	RespMimeTypes  []string
	Slugs          []string
}

func (hi *HandlerInfo) SetReqRespTypes(reqBodyContent openapi3.Content, respBodyContent openapi3.Content) {
	var componentStatusOk *openapi3.MediaType
	if hi.ReqMimeTypes != nil {
		componentStatusOk = reqBodyContent.Get(hi.ReqMimeTypes[0])
		pathStatusOkSchema := strings.Split(componentStatusOk.Schema.Ref, "/")

		reqType := pathStatusOkSchema[len(pathStatusOkSchema)-1]
		reqTypeVar := utils.ToCamelCase(reqType)
		hi.ReqType = reqType
		hi.ReqTypeVar = reqTypeVar
	}

	if hi.RespMimeTypes != nil {
		componentStatusOk = respBodyContent.Get(hi.RespMimeTypes[0])
		pathStatusOkSchema := strings.Split(componentStatusOk.Schema.Ref, "/")

		respType := pathStatusOkSchema[len(pathStatusOkSchema)-1]
		respTypeVar := utils.ToCamelCase(respType)
		hi.RespType = respType
		hi.RespTypeVar = respTypeVar
	}
}

func (hi *HandlerInfo) SetQueryParams(parameters openapi3.Parameters) {
	if parameters != nil {
		for _, param := range parameters {
			if param.Value.In == "query" {
				hi.QueryParams[utils.ToCamelCase(param.Value.Name)] = param.Value.Name
			}
		}
	}
}
