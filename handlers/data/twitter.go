package data

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
)

// Storing Twitter
type TwitterResults struct {
	Username string
	Handle   string
	DatePub  string
	Text     string
	Comments int16
	Retweets int16
	Likes    int16
	URL      string
}

func TwitterData(db *sql.DB) []TwitterResults {

	var results []TwitterResults

	// Run Python Scraper (TWITTER)
	cmd := exec.Command("./web/twitter.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())

	// More code here

	return results
}
