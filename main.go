package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var port string = ":8080"
var templates = template.Must(template.ParseFiles("templates/index.html"))

// Should use a struct for handling results on a page
type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/process", searchHandler)

	// Handle Static Files
	http.Handle("/static/images/", http.StripPrefix("/static/images/", http.FileServer(http.Dir("./static/images"))))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))

	fmt.Printf("Opening Server on Port %s:\n", port[1:5])
	log.Fatal(http.ListenAndServe(port, nil))

}

// handles home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// Redirect if not a real page
		http.Redirect(w, r, "/", http.StatusFound)
		fmt.Fprintf(w, "Error")
	}
	t.Execute(w, t)
}

// handles search results
func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // gets request method
	if r.Method != "GET" {
		t, _ := template.ParseFiles("templates/whoops.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("%s = %s\n", key, value)
			fmt.Fprintf(w, "%s\n", value)
		}
		// logic part here
	}
}
