package main

import (
	"AdventOfCode2025/Day05/common"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"runtime/debug"
)

func main() {
	freshIds, idsToSearch, _ := common.GetInput()

	fmt.Println("Part 1")
	fmt.Println("Input size:", len(freshIds), len(idsToSearch))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day05Part1")()

	freshIngredients := 0

	for _, id := range idsToSearch {
		for _, rang := range freshIds {
			if id >= rang.X && id <= rang.Y {
				freshIngredients++
				break
			}
		}
	}

	fmt.Println("Fresh ingredients: ", freshIngredients)

}
