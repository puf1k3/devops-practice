package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })

    http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
        data, err := os.ReadFile("/data/message.txt")
        if err != nil {
            http.Error(w, "File not found", http.StatusNotFound)
            return
        }

        w.Write(data)
    })

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
