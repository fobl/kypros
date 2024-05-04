package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("resources/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, "data goes here")
	})

	http.HandleFunc("/resources/images", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./resources/images"))))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
