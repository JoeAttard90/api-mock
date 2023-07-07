package pkg

import (
	"log"
	"net/http"
)

{{ range . }}
{{ .Handler }}
{{ end }}