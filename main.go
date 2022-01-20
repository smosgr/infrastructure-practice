package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var indexTemplate = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<link href="/assets/main.css" rel="stylesheet">
<h1>Hello from {{ . }}!</h1>
`))

func main() {
	name := os.Getenv("NAME")
	addr := os.Getenv("ADDR")

	if name == "" {
		log.Fatal("Please set the NAME environment variable")
	}

	if addr == "" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTemplate.Execute(w, name)
	})

	// 	mux.HandleFunc("/healthz", func (w http.ResponseWriter, r *http.Request) {
	// 		indexTemplate.Execute(w, name)
	// 	})
	//
	// 	mux.HandleFunc("/assets", func (w http.ResponseWriter, r *http.Request) {
	// 		indexTemplate.Execute(w, name)
	// 	})

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", addr)
	log.Println(server.ListenAndServe())
}
