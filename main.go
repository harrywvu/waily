package main

import (
	"daily-wins-cli/helpers"
	"database/sql"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var baseStatusMessage string = ""

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./wails.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS wails (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        timestamp TEXT,
        date TEXT,
        content TEXT,
        stream_id INTEGER
    )`)
	return db, err
}

func saveWailToDB(db *sql.DB, wail Wail) error {
	_, err := db.Exec(
		"INSERT INTO wails (timestamp, date, content, stream_id) VALUES (?, ?, ?, ?)",
		wail.Timestamp, wail.Date, wail.Content, wail.StreamID,
	)
	return err
}

func deleteStreamFromDB(db *sql.DB, streamID int) error {
	_, err := db.Exec("DELETE FROM wails WHERE stream_id = ?", streamID)
	return err
}

func startMenu(statusMsg string, db *sql.DB) {
	helpers.PrintHeader(statusMsg)

	input := helpers.GetKeyPress() // ← instant, no echo, no Enter

	switch strings.ToLower(input) {
	case "a":
		newStatus := addWail(db)
		startMenu(newStatus, db)

	case "v":
		hasStreams := viewStream(db)
		if !hasStreams {
			startMenu("No streams yet — add a wail first.", db)
			return
		}

		helpers.PrintActionBar("[1] Open stream   [2] Delete stream   [0] Back")
		choice := helpers.GetKeyPress() // ← single keypress

		switch choice {
		case "0":
			startMenu(baseStatusMessage, db)

		case "1":
			helpers.PrintInlinePrompt("Stream ID")
			streamID := helpers.GetUserInputInt() // ← typed, multi-digit

			viewWails(db, streamID)

			helpers.PrintActionBar("[1] Edit wail   [2] Delete wail   [0] Back")
			action := helpers.GetKeyPress() // ← single keypress

			switch action {
			case "0":
				startMenu(baseStatusMessage, db)
			case "1":
				helpers.PrintInlinePrompt("Wail ID to edit")
				startMenu(editWail(db, helpers.GetUserInputInt()), db)
			case "2":
				helpers.PrintInlinePrompt("Wail ID to delete")
				startMenu(deleteWail(db, helpers.GetUserInputInt()), db)
			default:
				startMenu("Invalid key — try again.", db)
			}

		case "2":
			helpers.PrintInlinePrompt("Stream ID to delete")
			startMenu(deleteStream(db, helpers.GetUserInputInt()), db)

		default:
			startMenu("Invalid key — try again.", db)
		}

	case "q":
		helpers.PrintNewLine()
		os.Exit(0)

	default:
		startMenu("Invalid key — try again.", db)
	}
}

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	startMenu(baseStatusMessage, db)
}