package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"database/sql"
	"fmt"
	"os"

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

func loadWailsFromDB(db *sql.DB) ([]Wail, error) {
	rows, err := db.Query("SELECT id, timestamp, date, content, stream_id FROM wails ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wails []Wail
	for rows.Next() {
		var w Wail
		err := rows.Scan(&w.ID, &w.Timestamp, &w.Date, &w.Content, &w.StreamID)
		if err != nil {
			return nil, err
		}
		wails = append(wails, w)
	}
	return wails, rows.Err()
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

func startMenu(startStatus string, masterStream *[]Wail, db *sql.DB) {
	var wrongInputStatus string = "You pressed a wrong key :("

	helpers.PrintHeader(startStatus)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	switch input[0] {
	case 'a', 'A':
		newStatusMessage := addWail(masterStream, db)
		startMenu(newStatusMessage, masterStream, db)
	case 'v', 'V':
		viewStream(masterStream)
		choice := choiceInViewStream()

		fmt.Print("Enter Stream [ID]")
		streamID := helpers.GetUserInputInt()

		if choice == 1 {

			// display all wails by date
			viewWails(masterStream, streamID)

		} else if choice == 2 {
			newStatusMessage := deleteStream(masterStream, streamID, db)
			startMenu(newStatusMessage, masterStream, db)
		} else {
			startMenu(wrongInputStatus, masterStream, db)
		}

	case 'q', 'Q':
		os.Exit(0)
	default:
		startMenu(wrongInputStatus, masterStream, db)
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

	// Load wails from DB into slice
	masterStream, err := loadWailsFromDB(db)
	if err != nil {
		panic(err)
	}

	startMenu(baseStatusMessage, &masterStream, db)
}
