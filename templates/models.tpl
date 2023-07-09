package models

{{ range . }}
type {{ .Name }} {{ .Fields }}
{{ end }}
