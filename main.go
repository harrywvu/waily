package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"os"
)

func startMenu() {

	startStatus := "Wail Stream is empty :("

	helpers.ClearTerminal()
	helpers.PrintShortcuts()
	helpers.PrintTitle()
	helpers.PrintStatus(startStatus)
	helpers.PrintOptions()

	helpers.PrintNewLine()
	helpers.PrintNewLine()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	switch input[0] {
	case 'a', 'A':
		addWail()
	case 'v', 'V':
		viewStream()
	default:
		startMenu()
	}
}

func main() {
	startMenu()
}
