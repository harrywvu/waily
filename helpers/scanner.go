package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// GetUserInputString reads a full line (Enter to confirm). Visible input — used
// for wail content where the user needs to see and edit what they type.
func GetUserInputString() string {
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(rawInput)
}

// GetUserInputInt reads a full line and parses an integer (Enter to confirm).
// Used for IDs where multi-digit input is needed.
func GetUserInputInt() int {
	var key int
	_, err := fmt.Scanln(&key)
	if err != nil {
		return 0
	}
	return key
}

// GetKeyPress reads exactly one byte from stdin without requiring Enter and
// without echoing the character to the terminal. Used for all menu navigation.
// Returns the lowercase character as a string. Returns "" on error.
func GetKeyPress() string {
	// Switch stdin to raw mode: no echo, no line buffering.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		// Fallback: normal line read (e.g. piped / non-TTY input).
		reader := bufio.NewReader(os.Stdin)
		raw, _ := reader.ReadString('\n')
		return strings.ToLower(strings.TrimSpace(raw))
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	buf := make([]byte, 1)
	os.Stdin.Read(buf)

	return strings.ToLower(string(buf))
}