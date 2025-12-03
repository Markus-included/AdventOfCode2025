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

	totalJoltage := 0

	for _, line := range input {

		const DIGITS int = 12

		joltage := 0

		windowStart :=
		windowEnd := len(line) -

		for i := 0; i < DIGITS; i++ {
			requiredRest := DIGITS - 1 - i
			search_end_pos := len(line) - requiredRest
			largest := common.LargestIdxInArray()
			fmt.Println(rest)
			rest = rest[largest+1:]
			joltage += int(math.Pow10(i))
		}

		totalJoltage += int(joltage)
	}

	fmt.Println("Total joltage: ", totalJoltage)
}
