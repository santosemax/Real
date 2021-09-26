package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// Overall Page Data
type Page struct {
	Title       string
	PageResults []Results
}

// Storing Results (ADD URL SECTION)
type Results struct {
	ResultTitle     string
	ResultBody      string
	ResultURL       string
	ResultSub       string
	ResultPermalink string
	hasContent      bool
}

// handles search results
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	// Render Page
	t, err := template.ParseFiles("templates/results.html")
	if err != nil {
		// Redirect if not a real page
		http.Redirect(w, r, "/", http.StatusFound)
		fmt.Fprintf(w, "Error")
	}

	var val string
	var page Page
	var results []Results

	// Initialize DB
	db, err := sql.Open("sqlite3", "./results.db")
	checkErr(err)

	// Create Tables (if needed)
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS queryQ (query text)")
	statement.Exec()
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS results (title text, body text, url text, subreddit text, permalink text)")
	statement.Exec()

	fmt.Println("method:", r.Method) // console shows method
	if r.Method != "GET" {
		t, _ := template.ParseFiles("templates/whoops.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("%s = %s\n", key, value)

			// Put query request into queryQ table
			val = strings.Join(value, "")
			statement, err = db.Prepare("INSERT INTO queryQ (query) VALUES (?)")
			statement.Exec(fmt.Sprint(val))
			checkErr(err)
			statement.Close()
		}
		// SEARCH LOGIC STARTS HERE (Python->DB->GO)

		// Run Python Scraper (REDDIT)
		cmd := exec.Command("./web/reddit.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())

		// DB STUFF
		// Read Result.db's temporary data from Python File
		rows, err := db.Query("SELECT * FROM results")
		checkErr(err)
		var title string
		var body string
		var url string
		var subreddit string
		var permalink string
		var content bool

		// Read the Rows
		for rows.Next() {

			// Read rows and output to html page
			err := rows.Scan(&title, &body, &url, &subreddit, &permalink)
			checkErr(err)

			if body == "Click to see video.image" {
				content = true
			} else {
				content = false
			}

			results = append(results,
				Results{
					ResultTitle:     title,
					ResultBody:      body,
					ResultURL:       url,
					ResultSub:       subreddit,
					ResultPermalink: permalink,
					hasContent:      content,
				})

		}
		rows.Close()

		// Delete Rows and close db
		db.Exec("DELETE FROM results")
		db.Exec("DELETE FROM queryQ")

		db.Close()
	}

	// Fill out Page
	page.Title = val
	page.PageResults = results

	// EXECUTE TEMPLATE HERE WITH STATIC VARIABLES
	t.Execute(w, page)

}

// Check DB for errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
