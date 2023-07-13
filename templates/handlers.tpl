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