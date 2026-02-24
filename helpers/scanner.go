package helpers

import (
	"bufio"
	"os"
	"strings"
)

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')

	return strings.TrimSpace(rawInput)
}
