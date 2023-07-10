package main

import (
	"api-mock-server/pkg/handlers"
	"context"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := flag.String("addr", ":8080", "the port on which to expose the server")
	flag.Parse()

	r := mux.NewRouter()
	{{ range $key, $value := .Endpoints }}
    r.Handle("{{ $key }}", handlers.{{ $value }}()){{ end }}


	srv := &http.Server{
		Addr:    *addr,
		Handler: r,
	}

	// Channel to listen for an interrupt or termination signal from the OS.
	// Use a buffered channel because the signal package requires it.
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Run our server in a goroutine so that it doesn't block.
		log.Printf("starting server on port %q", *addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stopChan // Block here until we receive the interrupt signal
	log.Println("Shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server gracefully stopped")
}
