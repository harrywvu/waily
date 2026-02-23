package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintStatus(s string) {
	if s != "" {
		fmt.Println(s)
	}
}

func PrintTitle() {
	fmt.Printf(`
        ▗▖ ▗▖ ▗▄▖ ▗▄▄▄▖▗▖ ▗▖  ▗▖
        ▐▌ ▐▌▐▌ ▐▌  █  ▐▌  ▝▚▞▘ 
        ▐▌ ▐▌▐▛▀▜▌  █  ▐▌   ▐▌  
        ▐▙█▟▌▐▌ ▐▌▗▄█▄▖▐▙▄▄▖▐▌  
        `)

	// Get today's date
	today := time.Now()
	fmt.Printf("%v\n", today.Format("Monday, January 2, 2006"))
}

func PrintNewLine() { fmt.Print("\n") }

func PrintOptions() {
	PrintNewLine()
	fmt.Println("[A] Add Daily Wail!")
	fmt.Print("[V] View Streams!")
}

func PrintShortcuts() {
	fmt.Println("[Q- Quit]")
}

func PrintHeader(statusMessage string) {
		ClearTerminal()
		PrintShortcuts()
		PrintTitle()
		PrintStatus(statusMessage)
		PrintOptions()
		PrintNewLine()
		PrintNewLine()
}