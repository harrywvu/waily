package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInputString() string {
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')

	return strings.TrimSpace(rawInput)
}

func GetUserInputInt() int {
	var key int
	_, err := fmt.Scanln(&key)

	if err != nil {
        fmt.Println("Error reading input:", err)
        return 0
    }

	return key
}
