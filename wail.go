package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
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

type Wail struct {
	Date      string `json:"date"`      // e.g., "2026-02-22"
	Timestamp string `json:"timestamp"` // e.g., "15:30:12"
	Content   string `json:"content"`
}

// returns the status message
func addWail(masterStream *[]Wail) string{
	helpers.PrintNewLine()
	fmt.Print("Enter daily wail: ")

	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.TrimSpace(rawInput)

	// create a new element within the stream json file

	// Create a new Wail Struct
	newWail := Wail {
		Date: time.Now().Format("2006-01-02"),
		Timestamp: time.Now().Format("15:04:05"),
		Content: userInput,
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

func viewStream() {
	helpers.ClearTerminal()
}
