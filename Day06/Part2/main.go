package main

import (
	"AdventOfCode2025/aocutil"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func GetInput() ([][]string, error) {
	data, err := os.ReadFile("inputs/day06.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	input := string(data)

	operatorsStart := strings.LastIndex(input, "\n") + 1

	operators := aocutil.Filter(strings.Split(input[operatorsStart:], " "), func(seg string) bool {
		return seg != ""
	})

	fmt.Println(operators)

	grid := input[:operatorsStart-1]

	grid = strings.ReplaceAll(grid, " ", "\t")
	grid = strings.ReplaceAll(grid, "\n", " ")
	grid = strings.ReplaceAll(grid, "\t", "\n")

	fmt.Println(grid)

	return nil, nil
}

func main() {
	input, _ := GetInput()

	fmt.Println("Part 2")
	fmt.Println("Input size:", len(input))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day06Part2")()

}
