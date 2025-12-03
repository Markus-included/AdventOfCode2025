package main

import (
	"AdventOfCode2025/Day03/common"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"math"
	"runtime/debug"
)

func main() {
	input, _ := common.GetInput()

	fmt.Println("Part 2")
	fmt.Println("Input size:", len(input))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day03Part2")()

	var totalJoltage int64 = 0

	for _, line := range input {

		const DIGITS int = 12

		var joltage int64 = 0
		windowStart := 0
		windowEnd := len(line) - DIGITS + 1
		for i := 0; i < DIGITS; i++ {
			slice := line[windowStart:min(windowEnd, len(line))]
			relIdx := common.LargestIdxInArray(slice)
			joltage += int64(math.Pow10(DIGITS-i-1)) * int64(slice[relIdx])
			windowStart += relIdx + 1
			windowEnd++
		}

		totalJoltage += int64(joltage)
	}

	fmt.Println("Total joltage: ", totalJoltage)
}
