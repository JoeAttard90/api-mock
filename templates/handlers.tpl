package handlers

import (
    {{if eq .HasPost true}}
    "encoding/json"
    "io"
    {{end}}
    "github.com/gorilla/mux"
	"log"
	"net/http"
)

{{ range .HandlersInfo }}
{{ .Handler }}
{{ end }}