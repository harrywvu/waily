package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"encoding/json"
	"os"
)

var baseStatusMessage string = "Wail Stream is empty :("


func loadOrCreateJSON(filename string, data interface{}) error {
    // Try to read the file
    fileData, err := os.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            // File doesn't exist, create it with empty data
            file, err := os.Create(filename)
            if err != nil {
                return err
            }
            defer file.Close()

            // Write initial JSON structure
            jsonBytes, _ := json.MarshalIndent(data, "", "  ")
            _, err = file.Write(jsonBytes)
            return err
        }
        return err
    }

    // File exists, unmarshal into data
    return json.Unmarshal(fileData, data)
}

func startMenu(startStatus string, masterStream *[]Wail) {
	var wrongInputStatus string = "You pressed a wrong key :("

	helpers.PrintHeader(startStatus)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	switch input[0] {
	case 'a', 'A':
		newStatusMessage := addWail(masterStream)
		startMenu(newStatusMessage, masterStream)
	case 'v', 'V':
		viewStream()
	default:
		startMenu(wrongInputStatus, masterStream)
	}
}

func main() {
	// master-stream.json
	// Load the JSON file, if it doesn't exist create it.
	var masterStream []Wail
	err := loadOrCreateJSON("master-stream.json", &masterStream)	
	if err != nil {
		panic(err)
	}

	jsonFile, _ := os.Open("master_stream.json")
	json.NewDecoder(jsonFile).Decode(&masterStream)
	jsonFile.Close()

	startMenu(baseStatusMessage, &masterStream)
}