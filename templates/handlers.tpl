package handlers

import (
    {{if eq .HasPost true}}
    "encoding/json"
    "io"
    {{end}}
    {{if eq .HasSlug true}}
    "github.com/gorilla/mux"
    {{end}}
    "api-mock-server/pkg/models"
    {{if ne .GlobalSecurityScheme "" }}
    "strings"
    {{end}}
	"log"
	"net/http"
)

{{ range .HandlersInfo }}
{{ .Handler }}
{{ end }}