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

func PrintMasterStreamFilteredByStreamID() {
	// OUTPUT:
/*
	ID			Date
	1		February 14, 2026
	2		February 15, 2026
	3		February 16, 2026
	4		February 17, 2026
*/

}

func PrintStream(){
		// OUTPUT:
/*
0		[FEBRUARY 15, 2026 STREAM]

	ID			Wail
	1		Finished a Workout
	2		Code Habit: 1 Hour Go
	3		10k Steps
	4		< 2 Hours Doomscrolling

*/
}