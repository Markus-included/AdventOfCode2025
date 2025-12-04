package common

import (
	"fmt"
	"os"
	"strings"
)

func GetInput() ([][]string, error) {
	data, err := os.ReadFile("inputs/day04.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	return grid, nil
}

var DIRS = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}
