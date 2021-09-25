package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/santosemax/ReAL/handlers"
)

var port string = ":8080"
var templates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/process", handlers.SearchHandler)

	// Handle Static Files
	http.Handle("/static/images/", http.StripPrefix("/static/images/", http.FileServer(http.Dir("./static/images"))))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))

	fmt.Printf("Opening Server on Port %s:\n", port[1:5])
	log.Fatal(http.ListenAndServe(port, nil))

}
