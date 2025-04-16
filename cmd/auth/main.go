package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Println("auth")

    // Start a simple HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Auth Service")
    })

    log.Println("Starting auth service on port 8081...")
    log.Fatal(http.ListenAndServe(":8081", nil))
}