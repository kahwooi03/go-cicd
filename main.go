package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Server started at :8080")
}
