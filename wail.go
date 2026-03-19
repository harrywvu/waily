package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"
)

type Wail struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Date      string `json:"date"`
	Content   string `json:"content"`
	StreamID  int    `json:"stream_id"`
}

func addWail(db *sql.DB) string {
	today := time.Now().Format("2006-01-02")
	var streamID int

	row := db.QueryRow("SELECT stream_id FROM wails WHERE date = ? LIMIT 1", today)
	err := row.Scan(&streamID)
	if err != nil && err != sql.ErrNoRows {
		return "Error checking existing stream."
	}

	if err == sql.ErrNoRows {
		row = db.QueryRow("SELECT MAX(stream_id) FROM wails")
		var maxID sql.NullInt64
		if err = row.Scan(&maxID); err != nil {
			return "Error getting max stream ID."
		}
		if maxID.Valid {
			streamID = int(maxID.Int64) + 1
		} else {
			streamID = 1
		}
	}

	helpers.PrintInlinePrompt("New wail")

	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.TrimSpace(rawInput)

	if userInput == "" {
		return "Wail cannot be empty."
	}

	now := time.Now()
	newWail := Wail{
		Timestamp: now.Format(time.RFC3339Nano),
		Date:      today,
		Content:   userInput,
		StreamID:  streamID,
	}

	if err = saveWailToDB(db, newWail); err != nil {
		return "Error saving wail."
	}
	return "Wail added successfully!"
}

// viewStream renders the stream list. Returns false when there are no streams.
func viewStream(db *sql.DB) bool {
	rows, err := db.Query(
		"SELECT date, MIN(stream_id) FROM wails GROUP BY date ORDER BY date DESC",
	)
	if err != nil {
		fmt.Println("Error querying streams:", err)
		return false
	}
	defer rows.Close()

	var ids []int
	var dates []string
	for rows.Next() {
		var date string
		var streamID int
		if err := rows.Scan(&date, &streamID); err != nil {
			continue
		}
		ids = append(ids, streamID)
		dates = append(dates, date)
	}

	helpers.PrintMasterStreamFilteredByStreamID(ids, dates)
	return len(ids) > 0
}

func deleteStream(db *sql.DB, streamID int) string {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM wails WHERE stream_id = ?", streamID).Scan(&count); err != nil {
		return "Error checking stream."
	}
	if count == 0 {
		return fmt.Sprintf("Stream [%d] does not exist.", streamID)
	}
	if err := deleteStreamFromDB(db, streamID); err != nil {
		return "Error deleting stream."
	}
	return fmt.Sprintf("Stream [%d] deleted successfully.", streamID)
}

// viewWails renders all wails for a given stream using PrintStream.
func viewWails(db *sql.DB, streamID int) {
	rows, err := db.Query(
		"SELECT id, timestamp, date, content FROM wails WHERE stream_id = ? ORDER BY timestamp ASC",
		streamID,
	)
	if err != nil {
		fmt.Println("Error querying wails:", err)
		return
	}
	defer rows.Close()

	var wailIDs []int
	var contents []string
	var date string

	for rows.Next() {
		var w Wail
		if err := rows.Scan(&w.ID, &w.Timestamp, &w.Date, &w.Content); err != nil {
			continue
		}
		if date == "" {
			date = w.Date
		}
		wailIDs = append(wailIDs, w.ID)
		contents = append(contents, w.Content)
	}

	helpers.PrintStream(streamID, date, wailIDs, contents)
}

func editWail(db *sql.DB, wailID int) string {
	var oldContent string
	err := db.QueryRow("SELECT content FROM wails WHERE id = ?", wailID).Scan(&oldContent)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Sprintf("Wail [%d] does not exist.", wailID)
		}
		return "Error retrieving wail."
	}

	fmt.Printf("\n  %sCurrent:%s %s\n", helpers.Dim, helpers.Reset, oldContent)
	helpers.PrintInlinePrompt("New content")
	newContent := helpers.GetUserInputString()

	if newContent == "" {
		return "Wail content cannot be empty."
	}

	result, err := db.Exec("UPDATE wails SET content = ? WHERE id = ?", newContent, wailID)
	if err != nil {
		return "Error updating wail."
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Sprintf("Wail [%d] does not exist.", wailID)
	}
	return fmt.Sprintf("Wail [%d] edited successfully.", wailID)
}

func deleteWail(db *sql.DB, wailID int) string {
	result, err := db.Exec("DELETE FROM wails WHERE id = ?", wailID)
	if err != nil {
		return "Error deleting wail."
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Sprintf("Wail [%d] does not exist.", wailID)
	}
	return fmt.Sprintf("Wail [%d] deleted successfully.", wailID)
}