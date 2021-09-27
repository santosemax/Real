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

	"github.com/santosemax/ReAL/handlers/data"
)

// Overall Page Data
type Page struct {
	Title             string
	PageRedditResults []data.RedditResults
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

	// Initialize DB
	db, err := sql.Open("sqlite3", "./results.db")
	checkErr(err)

	// Create Tables (if needed)
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS queryQ (query text)")
	statement.Exec()
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS redditQ (title text, body text, url text, subreddit text, permalink text)")
	statement.Exec()
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS twitterQ (title text, body text, url text, subreddit text, permalink text)")
	statement.Exec()

	// Get Query from User
	fmt.Println("method:", r.Method) // console shows method
	if r.Method != "GET" {
		t, _ := template.ParseFiles("templates/whoops.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("%s = %s\n", key, value)

			// Put query request into queryQ table to get Page title from Form
			val = strings.Join(value, "")
			statement, err = db.Prepare("INSERT INTO queryQ (query) VALUES (?)")
			statement.Exec(fmt.Sprint(val))
			checkErr(err)
			statement.Close()
		}
		// SEARCH LOGIC STARTS HERE (Python->DB->GO)

		// Run Python Scraper (TWITTER)
		cmd := exec.Command("./web/twitter.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())

		// Fill out Page (using 'data' handlers)
		page.Title = val
		page.PageRedditResults = data.RedditData(db) // Call reddit data here????
		// page.PageTwitterResults = ???

		// Delete Rows and close db
		db.Exec("DELETE FROM redditQ")
		db.Exec("DELETE FROM twitterQ")
		db.Exec("DELETE FROM queryQ")

		db.Close()
	}

	// EXECUTE TEMPLATE HERE WITH STATIC VARIABLES
	t.Execute(w, page)

}

// Check DB for errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
