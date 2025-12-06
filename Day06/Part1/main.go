package main

import (
	"AdventOfCode2025/aocutil"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

func GetInput() ([][]string, error) {
	data, err := os.ReadFile("inputs/day06.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	input := string(data)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := aocutil.Map(lines, func(s string) []string {
		return aocutil.Filter(strings.Split(s, " "), func(seg string) bool {
			return seg != ""
		})
	})

	return grid, nil
}

func main() {
	input, _ := GetInput()

	fmt.Println("Part 1")
	fmt.Println("Input size:", len(input))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day06Part1")()

	grandTotal := 0

	for x := range input[1] {
		result := 0
		operator := input[len(input)-1][x]
		if operator == "*" {
			result = 1
		}
		for y := range input[:len(input)-1] {
			fmt.Print(input[y][x], " ", operator, " ")
			n, _ := strconv.Atoi(input[y][x])
			switch operator {
			case "*":
				result *= n
			case "+":
				result += n
			}
		}
		fmt.Println(" = ", result)
		grandTotal += result
	}
	fmt.Println("Grand total: ", grandTotal)
}
