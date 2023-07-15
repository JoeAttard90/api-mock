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

        {{- if ne (len .Slugs) 0 }}
        vars := mux.Vars(r)
        {{- range .Slugs }}
        {{ . }} := vars["{{ . }}"]
        log.Printf("received %q request for {{ . }}, id: %q", r.Method, {{ . }})
        {{ end }} {{ end }}

        {{- if ne .SecurityScheme ""}}
        reqToken := r.Header.Get("Authorization")
        splitToken := strings.Split(reqToken, "{{ .SecurityScheme }}")
        if len(splitToken) != 2 {
            log.Println("no auth token provided")
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        {{ end }}
        {{- if gt (len .QueryParams) 0 }}
        urlQuery := r.URL.Query()

        {{- range $key, $value := .QueryParams }}
        {{ $key }} := urlQuery.Get("{{ $value }}")
        if {{ $key }} == "" {
            log.Printf("missing query parameter %q",  "{{ $key }}")
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        {{ end }} {{ end }}
        {{- range .ReqMimeTypes }}
        if contentType == "{{ . }}" {
            log.Printf("Received {{ . }} data")
        }
        {{ end }}

        {{- if ne .ReqTypeVar "" }}
        requestBody, err := io.ReadAll(r.Body)
        if err != nil {
            log.Printf("unable to read request body: %s", err.Error())
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        var {{ .ReqTypeVar }} structs.{{ .ReqType }}
        err = json.Unmarshal(requestBody, &{{ .ReqTypeVar }})
        if err != nil {
            log.Printf("unable to unmarshal request body")
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        log.Println({{ .ReqTypeVar }})
        {{end}}

        {{- if and (ne .RespTypeVar "") (ne .RespType "") (ne .RespType .ReqType)}} var {{ .RespTypeVar }} structs.{{ .RespType }} {{end}}
        {{- if ne .StaticResponsePath "" }}
        file, err := os.Open("{{ .StaticResponsePath }}")
        if err != nil {
            log.Printf("failed to open file: %v", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }
        defer func(file *os.File) {
            err := file.Close()
            if err != nil {
                log.Printf("could not close file %q: %s", file.Name(), err.Error())
            }
        }(file)

        // Read file content
        respBytes, err := io.ReadAll(file)
        if err != nil {
            log.Printf("failed to read file: %v", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }

        // Check the example provided matches the specs (generated structs)
        err = json.Unmarshal(respBytes, &{{ .RespTypeVar }})
        if err != nil {
            log.Printf("unable to unmarshal request body")
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        {{- else if and (ne .RespTypeVar "") (eq .StaticResponsePath "") }}
        {{ .RespTypeVar }}.FakeIt()
        respBytes, err := json.Marshal({{ .RespTypeVar }})
        if err != nil {
            log.Printf("unable to marshal response body")
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        {{ end }}
        {{- if or (ne .RespTypeVar "") (ne .StaticResponsePath "") }}
        _, err = w.Write(respBytes)
        if err != nil {
            log.Printf("error writing response body")
            return
        }
        {{- end }}

        {{- if eq .RespTypeVar "" }}
        w.WriteHeader(http.StatusOK)
        {{- end }}
    }
}