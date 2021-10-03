package data

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
)

// Storing Reddit
type RedditResults struct {
	ResultTitle     string
	ResultBody      string
	ResultURL       string
	ResultSub       string
	ResultPermalink string
	hasContent      bool
}

func RedditData(db *sql.DB) []RedditResults {

	var results []RedditResults

	// START HERE
	// Run Python Scraper (REDDIT)
	cmd := exec.Command("./web/reddit.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())

	// DB STUFF
	// Read Result.db's temporary data from Python File
	rows, err := db.Query("SELECT * FROM redditQ")
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
			RedditResults{
				ResultTitle:     title,
				ResultBody:      body,
				ResultURL:       url,
				ResultSub:       subreddit,
				ResultPermalink: permalink,
				hasContent:      content,
			})

	}
	rows.Close()
	// END HERE

	// DEBUG
	// fmt.Printf("Results Here %v", results)
	return results
}

// Check DB for errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
