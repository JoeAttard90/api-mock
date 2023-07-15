package handlers

import (
    {{if eq .HasPost true}}
    "encoding/json"
    "io"
    {{end}}
    {{if eq .HasSlug true}}
    "github.com/gorilla/mux"
    {{end}}
    "api-mock-server/pkg/structs"
	"log"
	"net/http"
	{{if ne .StaticResponses ""}}
	"os"
    {{end}}
    {{if ne .GlobalSecurityScheme "" }}
    "strings"
    {{end}}
)

{{ range .HandlersInfo }}
{{ .Handler }}
{{ end }}

{{ range $endpoint, $methods := .Methods }}
type {{ parseEndpoint $endpoint }}Struct struct {
	{{ range $method := $methods }}
        {{ range $k, $v := $method }}
            {{ $k }} http.HandlerFunc
        {{ end }}
    {{ end }}
}
func (h *{{parseEndpoint $endpoint}}Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    {{ range $method := $methods }}
        {{ range $k, $v := $method }}
            case "{{ $k }}":
            if h.{{ $k }} != nil {
                h.{{ $k }}.ServeHTTP(w, r)
            }
        {{ end }}{{ end }}
    default:
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    }
}
{{ end }}