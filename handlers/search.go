package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/santosemax/ReAL/handlers/data"
)

// Overall Page Data
type Page struct {
	Title              string
	PageRedditResults  []data.RedditResults
	PageTwitterResults []data.TwitterResults
	PageStackResults   []data.StackResults
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
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS twitterQ (username text, handle text, datepub text, text text, retweets integer, likes integer, url text)")
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

		// Fill out Page (using 'data' handlers)
		page.Title = val
		page.PageRedditResults = data.RedditData(db) // Call reddit data here????
		page.PageTwitterResults = data.TwitterData(db)

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
