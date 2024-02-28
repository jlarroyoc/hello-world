package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    title := os.Getenv("TITLE")
    message := os.Getenv("MESSAGE")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Title: %s\nMessage: %s",title, message)
    })

    fmt.Println("Starting server on port 8080.")
    http.ListenAndServe(":8080", nil)
}
