package common

import (
	"fmt"
	"os"
	"strings"
)

func GetInput() ([]string, error) {
	data, err := os.ReadFile("inputs/day07.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines, nil
}
