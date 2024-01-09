package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)

func main() {
    hostname, err := os.Hostname()
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        indexContent, err := ioutil.ReadFile("static/index.html")
        if err != nil {
            http.Error(w, "Error reading index.html", http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, string(indexContent)+"\nVersion: 1.0.0\n"+hostname+"\n")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Catch-all route to serve index.html for all other routes
    http.HandleFunc("/{any}", func(w http.ResponseWriter, r *http.Request) {
        indexContent, err := ioutil.ReadFile("static/index.html")
        if err != nil {
            http.Error(w, "Error reading index.html", http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, string(indexContent)+"\nVersion: 1.0.0\n"+hostname+"\n")
    })

    http.ListenAndServe(":8080", nil)
}








