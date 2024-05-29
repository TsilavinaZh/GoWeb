package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", HomeHandler)
    http.HandleFunc("/about", AboutHandler)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}