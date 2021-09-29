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
	Retweets int64
	Likes    int64
	URL      string
	hasText  bool
}

func TwitterData(db *sql.DB) []TwitterResults {

	var results []TwitterResults

	// Run Python Scraper (TWITTER)
	cmd := exec.Command("./web/twitter.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())

	// DB STUFF
	// Read Result.db's temporary data from Python File
	rows, err := db.Query("SELECT * FROM twitterQ")
	checkErr(err)
	var username string
	var handle string
	var datepub string
	var text string
	var retweets int64
	var likes int64
	var url string
	var isText bool

	for rows.Next() {
		// Read rows and output to html page
		err := rows.Scan(
			&username,
			&handle,
			&datepub,
			&text,
			&retweets,
			&likes,
			&url,
		)
		checkErr(err)

		if text == "MEDIACONTENTONLY" {
			isText = false
		} else {
			isText = true
		}

		// Append Results
		results = append(results,
			TwitterResults{
				Username: username,
				Handle:   handle,
				DatePub:  datepub,
				Text:     text,
				Retweets: retweets,
				Likes:    likes,
				URL:      url,
				hasText:  isText,
			})

	}
	rows.Close()
	return results
}
