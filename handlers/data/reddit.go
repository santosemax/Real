package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/JesusIslam/tldr"
)

// Storing Reddit
type RedditResults struct {
	ResultTitle     string
	ResultBody      string
	ResultURL       string
	ResultSub       string
	ResultPermalink string
	ResultSummary   string
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
	var summarizedList []string // not apart of db

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

		// Make summary of body
		intoSentences := 3
		bag := tldr.New()
		summarizedList, _ = bag.Summarize(body, intoSentences)
		fmt.Println(summarizedList)
		// Convert from list to string
		summarizedBody := strings.Join(summarizedList, " ")

		results = append(results,
			RedditResults{
				ResultTitle:     title,
				ResultBody:      body,
				ResultURL:       url,
				ResultSub:       subreddit,
				ResultPermalink: permalink,
				ResultSummary:   summarizedBody,
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
