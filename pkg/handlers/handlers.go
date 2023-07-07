package pkg

import (
	"log"
	"net/http"
)


func GetUploadEntitlement() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "GET", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func PostMetaData() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "POST", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func SearchBusinessUserDocs() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "POST", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func SearchClientUserDocs() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "POST", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func PostReadStatus() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "POST", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func GetBusinessUserDetails() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "GET", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func GetClientFundData() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "GET", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func GetClientUserDetails() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "GET", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func GetDocEntitlement() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            log.Printf("incorrect request method provided wanted: %q got: %q", "GET", r.Method)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}
