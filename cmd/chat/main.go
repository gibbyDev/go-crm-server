package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Println("chat")

    // Start a simple HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Chat Service")
    })

    log.Println("Starting chat service on port 8082...")
    log.Fatal(http.ListenAndServe(":8082", nil))
}