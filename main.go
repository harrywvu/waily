package main

import (
	"bufio"
	"daily-wins-cli/helpers"
	"os"
)

func main() {
	helpers.ClearTerminal()
	helpers.PrintShortcuts()
	helpers.PrintTitle()
	helpers.PrintOptions()

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	switch input[0] {
	case 'a', 'A':
		addWail()
	case 'v', 'V':
		viewStream()
	default:
	}
}
