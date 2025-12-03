package main

import (
	"AdventOfCode2025/Day03/common"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"runtime/debug"
)

func main() {
	input, _ := common.GetInput()

	fmt.Println("Part 1")
	fmt.Println("Input size:", len(input))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day03Part1")()

	totalJoltage := 0

	for _, line := range input {
		i := common.LargestIdxInArray(line[:len(line)-1])
		rest := line[i+1:]
		j := common.LargestIdxInArray(rest)

		joltage := (10 * line[i]) + rest[j]

		totalJoltage += int(joltage)
	}

	fmt.Println("Total joltage: ", totalJoltage)
}
