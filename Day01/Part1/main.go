package main

import (
	"AdventOfCode2025/mprofiler"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	fmt.Println("Part 1")
	fmt.Println("Lines:", len(lines))

	debug.SetGCPercent(-1)
	defer mprofiler.TrackExecutionTime("Day01Part1")()

	dial := 50
	zeroPositions := 0

	for lineNum, line := range lines {
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}
		i, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to parse line: ", lineNum, " ", line)
		}
		dial = (dial + (sign * i)) % 100

		if dial == 0 {
			zeroPositions += 1
		}
	}

	fmt.Println("Passcode: ", zeroPositions)
}
