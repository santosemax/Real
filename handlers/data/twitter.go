package data

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
)

// Storing Twitter
type TwitterResults struct {
	ResultTitle string
	ResultBody  string
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
