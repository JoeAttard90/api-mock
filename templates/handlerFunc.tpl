func {{ .Path }}() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "{{ .Method }}" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "{{ .Method }}", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        {{ if ne (len .ReqMimeTypes) 0 }}
        contentType := r.Header.Get("Content-Type")
        {{ end }}

        {{ if ne (len .Slugs) 0 }}
        vars := mux.Vars(r)
        {{ range .Slugs }}
        {{ . }} := vars["{{ . }}"]
        log.Printf("received request for {{ . }}, id: %q", {{ . }})
        {{ end }}
        {{ end }}

        {{ if ne .SecurityScheme ""}}
        reqToken := r.Header.Get("Authorization")
        splitToken := strings.Split(reqToken, "{{ .SecurityScheme }}")
        if len(splitToken) != 2 {
            log.Println("no auth token provided")
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        {{ end }}

        {{ if gt (len .QueryParams) 0 }}
        urlQuery := r.URL.Query()

        {{ range $key, $value := .QueryParams }}
        {{ $key }} := urlQuery.Get("{{ $value }}")
        if {{ $key }} == "" {
            log.Printf("missing query parameter %q",  "{{ $key }}")
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        {{ end }}

        {{ end }}

        {{ range .ReqMimeTypes }}
        if contentType == "{{ . }}" {
            log.Printf("Received {{ . }} data")
        }
        {{ end }}

        {{if ne .ReqTypeVar "" }}
        requestBody, err := io.ReadAll(r.Body)
        if err != nil {
            log.Printf("unable to read request body: %s", err.Error())
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        var {{ .ReqTypeVar }} models.{{ .ReqType }}
        err = json.Unmarshal(requestBody, &{{ .ReqTypeVar }})
        if err != nil {
            log.Printf("unable to unmarshal request body")
            http.Error(w, "bad request", http.StatusBadRequest)
        }

        log.Println({{ .ReqTypeVar }})
        {{end}}

        {{if and (ne .ReqTypeVar .RespTypeVar) (ne .RespTypeVar "") }}
        var {{ .RespTypeVar }} models.{{ .RespType }}
        {{ end }}
        {{if ne .RespTypeVar "" }}
        {{ .RespTypeVar }}.FakeIt()
        respBytes, err := json.Marshal({{ .RespTypeVar }})
        if err != nil {
            log.Printf("unable to unmarshal request body")
            http.Error(w, "bad request", http.StatusBadRequest)
        }

        _, err = w.Write(respBytes)
        if err != nil {
            return
        }
        {{end}}
        {{if eq .RespTypeVar "" }}
        w.WriteHeader(http.StatusOK)
        {{ end }}
    }
}