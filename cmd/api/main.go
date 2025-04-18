package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Println("api")

    // Start a simple HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "API Service")
    })

    log.Println("Starting API service on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}