package data

import "database/sql"

// Storing Twitter
type TwitterResults struct {
	ResultTitle string
	ResultBody  string
}

func TwitterData(db *sql.DB) []TwitterResults {

	var results []TwitterResults

	// Code Here

	return results
}
