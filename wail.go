package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"fmt"
	"os"
	"strings"
)

type Wail struct {
	Date      string `json:"date"`      // e.g., "2026-02-22"
	Timestamp string `json:"timestamp"` // e.g., "15:30:12"
	Content   string `json:"content"`
}

func addWail() {
	helpers.PrintNewLine()
	fmt.Print("Enter daily wail: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	
}

func viewStream() {
	helpers.ClearTerminal()
}
