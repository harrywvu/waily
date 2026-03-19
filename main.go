package main

import (
	"daily-wins-cli/helpers"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var baseStatusMessage string = "Wail Stream is empty :("

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
	_, err := db.Exec("INSERT INTO wails (timestamp, date, content, stream_id) VALUES (?, ?, ?, ?)",
		wail.Timestamp, wail.Date, wail.Content, wail.StreamID)
	return err
}

func deleteStreamFromDB(db *sql.DB, streamID int) error {
	_, err := db.Exec("DELETE FROM wails WHERE stream_id = ?", streamID)
	return err
}

func startMenu(startStatus string, db *sql.DB) {
	var wrongInputStatus string = "You pressed a wrong key :("

	helpers.PrintHeader(startStatus)

	input := strings.ToLower(helpers.GetUserInputString())

	switch input {
	case "a":
		newStatusMessage := addWail(db)
		startMenu(newStatusMessage, db)
	case "v":
		viewStream(db)
		choice := choiceInViewStream()

		fmt.Print("Enter Stream [ID]")
		streamID := helpers.GetUserInputInt()

		if choice == 1 {

			// display all wails by date
			viewWails(db, streamID)

			// give them the option to either edit or delete a wail
			fmt.Print("EDIT [1]		DELETE [2]")
			var key int = helpers.GetUserInputInt()

			if key == 1 {
				fmt.Print("Enter Wail to Edit [ID]: ")
				var editKey int = helpers.GetUserInputInt()
				var newStatusMessage string = editWail(db, editKey)
				startMenu(newStatusMessage, db)

			} else if key == 2 {
				
				viewWails(db, streamID)
				// get wailID to be deleted
				fmt.Print("Enter Wail to Delete [ID]: ")
				var deleteKey int = helpers.GetUserInputInt()
				var newStatusMessage string = deleteWail(db, deleteKey)
				startMenu(newStatusMessage, db)

			} else {
				startMenu(wrongInputStatus, db)
			}

		} else if choice == 2 {
			newStatusMessage := deleteStream(db, streamID)
			startMenu(newStatusMessage, db)
		} else {
			startMenu(wrongInputStatus, db)
		}

	case "q":
		os.Exit(0)
	default:
		startMenu(wrongInputStatus, db)
	}
}

// choose either to edit a wail or delete the stream
func choiceInViewStream() int {

	fmt.Print("[1] Select 		[2] Delete Stream\n")
	var key int = helpers.GetUserInputInt()

	return key
}

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	startMenu(baseStatusMessage, db)
}
