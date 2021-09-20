package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
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

	fmt.Println("method:", r.Method) // console shows method
	if r.Method != "GET" {
		t, _ := template.ParseFiles("templates/whoops.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("%s = %s\n", key, value)
			fmt.Fprintf(w, "%s\n", value)
			//fmt.Fprintf(w, title)
			//fmt.Fprintf(w, body)
		}
		// SEARCH LOGIC STARTS HERE (Python->DB->GO)

		// Initialize DB
		db, err := sql.Open("sqlite3", "./results.db")
		checkErr(err)

		// Run Python Scraper (REDDIT)
		cmd := exec.Command("./web/scrape.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())

		// Read Result.db's temporary data from Python File
		rows, err := db.Query("SELECT * FROM results")
		var title string
		var body string

		for rows.Next() {
			err = rows.Scan(&title, &body)
			checkErr(err)
			fmt.Fprintf(w, "\n%s", title)
			fmt.Fprintf(w, "\n%s\n", body)
			for i := 0; i < 70; i++ {
				fmt.Fprintf(w, "-")
			}
			fmt.Fprintf(w, "\n")
		}

		rows.Close()

		// Delete Rows that we just used (Not sure if dev/production)

		db.Close()
	}
}

// Check DB for errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
