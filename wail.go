package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"encoding/json"
	"fmt"
	"os"
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
	Content  string `json:"content"` // "Went to the gym"  				<- For Display
	StreamID string `json:stream_id` // "1"								<- For easy stream selection
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

func viewStream(masterStream *[]Wail) {
	helpers.PrintNewLine()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

	// Iterate and print using tabs (\t) to separate columns
	for _, stream := range *masterStream {
		fmt.Fprintf(w, "%d\t%s\n", stream.StreamID, stream.Date)
	}

	// Flush ensures the buffer is written to the terminal
	w.Flush()

	fmt.Print("Select Stream : ")

	// userInput := helpers.getUserInput()



}
