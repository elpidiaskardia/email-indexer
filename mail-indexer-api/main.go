package main

import (
    "log"
    "net/http"
    "Api_Go/api"
)

func main() {

    server := http.Server{
		Addr:    ":8080",
		Handler: api.NewRouter(),
	}

    log.Printf("Server listening on port %s...", "8080")
    if err :=	server.ListenAndServe()   ; err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
