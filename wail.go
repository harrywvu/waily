package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"database/sql"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

// GetHighestStreamID iterates through the slice and returns the maximum StreamID value.
func GetHighestStreamID(masterStream *[]Wail) int {	
	if len(*masterStream) == 0 {
		return 0
	}

	max := 0
	for _, wail := range *masterStream {
		if wail.StreamID > max {
			max = wail.StreamID
		}
	}
	return max
}

type Wail struct {
	ID        int    `json:"id"`        // Auto-incremented ID
	Timestamp string `json:"timestamp"` // Full timestamp
	Date      string `json:"date"`      // "2026-02-22" <- For Display
	Content   string `json:"content"`   // "Went to the gym" <- For Display
	StreamID  int    `json:"stream_id"` // Stream ID as int
}

type StreamView struct {
	StreamID int
	Date     string
}

// returns the status message
func addWail(masterStream *[]Wail, db *sql.DB) string {

	maxExistingID := GetHighestStreamID(masterStream)
	newStreamID := maxExistingID + 1

	helpers.PrintNewLine()
	fmt.Print("Enter daily wail: ")

	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.TrimSpace(rawInput)

	now := time.Now()
	newWail := Wail{
		Timestamp: now.Format(time.RFC3339Nano),
		Date:      now.Format("2006-01-02"),
		Content:   userInput,
		StreamID:  newStreamID,
	}

	// Insert into DB
	err := saveWailToDB(db, newWail)
	if err != nil {
		return "Error saving wail :("
	}

	// Reload from DB
	*masterStream, err = loadWailsFromDB(db)
	if err != nil {
		return "Error reloading wails :("
	}

	return "Wail added successfully! :D"
}

// Shows Streams per dates
func viewStream(masterStream *[]Wail) {
	helpers.PrintNewLine()

	dateMap := make(map[string]int)

	// create a map of the masterStream
	for _, wail := range *masterStream {
		if _, exists := dateMap[wail.Date]; !exists {
			dateMap[wail.Date] = wail.StreamID
		}
	}

	streams := make([]StreamView, 0, len(dateMap))
	for date, id := range dateMap {
		streams = append(streams, StreamView{StreamID: id, Date: date})
	}

	sort.Slice(streams, func(i, j int) bool {
		return streams[i].Date > streams[j].Date
	})

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	for _, stream := range streams {
		fmt.Fprintf(w, "%d\t%s\n", stream.StreamID, stream.Date)
	}

	// Flush ensures the buffer is written to/ the terminal
	w.Flush()
}

func deleteStream(masterStream *[]Wail, streamID int, db *sql.DB) string {
	err := deleteStreamFromDB(db, streamID)
	if err != nil {
		return "Error deleting stream :("
	}

	// Reload from DB
	*masterStream, err = loadWailsFromDB(db)
	if err != nil {
		return "Error reloading wails :("
	}

	return "Successfully deleted a stream :("
}

func viewWails(masterStream *[]Wail, streamID int) {
	helpers.PrintNewLine()

	var matchingWails []Wail
	for _, wail := range *masterStream {
		if wail.StreamID == streamID {
			matchingWails = append(matchingWails, wail)
		}
	}

	sort.Slice(matchingWails, func(i, j int) bool {
		return matchingWails[i].Date > matchingWails[j].Date
	})

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	fmt.Fprintln(w, "Date\tContent\tID")
	for _, wail := range matchingWails {
		fmt.Fprintf(w, "%s\t%s\t%d\n", wail.Date, wail.Content, wail.ID)
	}
	w.Flush()
}
