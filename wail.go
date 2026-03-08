package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

func saveWails(filename string, masterStream []Wail) error {
	// Convert slice to JSON
	jsonBytes, err := json.MarshalIndent(masterStream, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	err = os.WriteFile(filename, jsonBytes, 0644)
	return err
}

// GetHighestStreamID iterates through the slice and returns the maximum StreamID value.
func GetHighestStreamID(masterStream *[]Wail) int {
	if len(*masterStream) == 0 {
		return 0
	}

	max := 0
	for _, wail := range *masterStream {
		var streamID int
		fmt.Sscanf(wail.StreamID, "%d", &streamID)
		if streamID > max {
			max = streamID
		}
	}
	return max
}

type Wail struct {
	ID       string `json:"id"`      // "2006-01-02T15:04:05.999999999Z07:0 		<- UNIQUE ID
	Date     string `json:"date"`    // "2026-02-22" 				  			<- For Display
	Content  string `json:"content"` // "Went to the		 gym"  				<- For Display
	StreamID string `json:stream_id` // "1"								<- For easy stream selection
}

type StreamView struct {
	StreamID string
	Date     string
}

// returns the status message
func addWail(masterStream *[]Wail) string {

	maxExistingID := GetHighestStreamID(masterStream)
	newStreamID := maxExistingID + 1

	helpers.PrintNewLine()
	fmt.Print("Enter daily wail: ")

	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.TrimSpace(rawInput)

	now := time.Now()
	newWail := Wail{
		ID:       now.Format(time.RFC3339Nano),
		Date:     time.Now().Format("2006-01-02"),
		Content:  userInput,
		StreamID: fmt.Sprintf("%d", newStreamID),
	}

	// Add the wail
	*masterStream = append(*masterStream, newWail)

	// Save back to JSON file
	err := saveWails("master-stream.json", *masterStream)
	if err != nil {
		return "Error saving wail :("
	}

	return "Wail added successfully! :D"
}

// Shows Streams per datez
func viewStream(masterStream *[]Wail) {
	helpers.PrintNewLine()

	dateMap := make(map[string]string)

	// create a map of the masterStream
	for _, wail := range *masterStream {
		if _, exists := dateMap[wail.Date]; 
		!exists {
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
		fmt.Fprintf(w, "%s\t%s\n", stream.StreamID, stream.Date)
	}

	// Flush ensures the buffer is written to/ the terminal
	w.Flush()
}
 
func deleteStream(masterStream *[]Wail, streamID int) {
	streamIDtoBeDeleted := fmt.Sprintf("%d", streamID)

	var toBeDeleted []int

	// iterate through the masterStream(dereferenced)
	for i, stream := range *masterStream{
		// collect the indices that are to be deleted 
		if stream.StreamID == streamIDtoBeDeleted {
			toBeDeleted = append(toBeDeleted, i)
		}
	}

	// delete the elements at the indexes collected 
	// pop and swap technique is used here, so elements will be deleted in reverse
	// 2. truncate the slice to remove the duplicate

	for j := len(toBeDeleted) - 1; j >= 0; j--{
			// use the last element in toBeDeleted as the index to delete in the masterStream
		i := toBeDeleted[j]
		lastIndexInMasterStream := len(*masterStream) - 1

			// 1. replace element to be deleted with the last element
		(*masterStream)[i] = (*masterStream)[lastIndexInMasterStream]

			// 2. truncate the slice to remove the duplicate
		*masterStream = (*masterStream)[:lastIndexInMasterStream]
	}

	saveWails("master-stream.json", *masterStream)
}