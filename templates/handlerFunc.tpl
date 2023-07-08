func {{ .Path }}() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "{{ .Method }}" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "{{ .Method }}", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        {{if eq .Method "POST"}}
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

        {{if eq .Method "GET"}}
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


        w.WriteHeader(http.StatusOK)
    }
}