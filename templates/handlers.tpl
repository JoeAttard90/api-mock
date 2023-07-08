package pkg

import (
    {{if eq .HasPost true}}
    "encoding/json"
    "io"
    {{end}}
	"log"
	"net/http"
)

{{ range .HandlersInfo }}
{{ .Handler }}
{{ end }}