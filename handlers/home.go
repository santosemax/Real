package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// handles home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// Redirect if not a real page
		http.Redirect(w, r, "/", http.StatusFound)
		fmt.Fprintf(w, "Error")
	}
	t.Execute(w, t)
}
