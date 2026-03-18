package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type Wail struct {
	ID        int    `json:"id"`        // Auto-incremented ID
	Timestamp string `json:"timestamp"` // Full timestamp
	Date      string `json:"date"`      // "2026-02-22" <- For Display
	Content   string `json:"content"`   // "Went to the gym" <- For Display
	StreamID  int    `json:"stream_id"` // Stream ID as int
}

// returns the status message
func addWail(db *sql.DB) string {

	today := time.Now().Format("2006-01-02")
	var streamID int

	// Check if a stream already exists for today
	row := db.QueryRow("SELECT stream_id FROM wails WHERE date = ? LIMIT 1", today)
	err := row.Scan(&streamID)
	if err != nil && err != sql.ErrNoRows {
		return "Error checking existing stream :("
	}

	// If no stream for today, create a new one
	if err == sql.ErrNoRows {
		row = db.QueryRow("SELECT MAX(stream_id) FROM wails")
		var maxID sql.NullInt64
		err = row.Scan(&maxID)
		if err != nil {
			return "Error getting max stream ID :("
		}
		if maxID.Valid {
			streamID = int(maxID.Int64) + 1
		} else {
			streamID = 1
		}
	}

	helpers.PrintNewLine()
	fmt.Print("Enter daily wail: ")

	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.TrimSpace(rawInput)

	now := time.Now()
	newWail := Wail{
		Timestamp: now.Format(time.RFC3339Nano),
		Date:      today,
		Content:   userInput,
		StreamID:  streamID,
	}

	// Insert into DB
	err = saveWailToDB(db, newWail)
	if err != nil {
		return "Error saving wail :("
	}

	return "Wail added successfully! :D"
}

// Shows Streams per dates
func viewStream(db *sql.DB) {
	helpers.PrintNewLine()

	rows, err := db.Query("SELECT DISTINCT date, stream_id FROM wails ORDER BY date DESC")
	if err != nil {
		fmt.Println("Error querying streams:", err)
		return
	}
	defer rows.Close()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	for rows.Next() {
		var date string
		var streamID int
		err := rows.Scan(&date, &streamID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Fprintf(w, "%d\t%s\n", streamID, date)
	}
	w.Flush()
}

func deleteStream(db *sql.DB, streamID int) string {
	err := deleteStreamFromDB(db, streamID)
	if err != nil {
		return "Error deleting stream :("
	}

	return "Successfully deleted a stream :("
}

func viewWails(db *sql.DB, streamID int) {
	helpers.PrintNewLine()

	rows, err := db.Query("SELECT id, timestamp, date, content FROM wails WHERE stream_id = ? ORDER BY date DESC", streamID)
	if err != nil {
		fmt.Println("Error querying wails:", err)
		return
	}
	defer rows.Close()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	fmt.Fprintln(w, "Date\tContent\tID")
	for rows.Next() {
		var wail Wail
		err := rows.Scan(&wail.ID, &wail.Timestamp, &wail.Date, &wail.Content)
		if err != nil {
			fmt.Println("Error scanning wail:", err)
			continue
		}
		fmt.Fprintf(w, "%s\t%s\t%d\n", wail.Date, wail.Content, wail.ID)
	}
	w.Flush()
}

func deleteWail(db *sql.DB, wailID int) {
	
}