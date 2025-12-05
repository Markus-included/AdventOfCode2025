package main

import (
	"AdventOfCode2025/Day05/common"
	"AdventOfCode2025/aocutil"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"runtime/debug"
)

func main() {
	freshIdRanges, _, _ := common.GetInput()

	fmt.Println("Part 2")
	fmt.Println("Input size:", len(freshIdRanges))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day05Part2")()

	// Merging the ranges, then calculating rang.Y - rang.X + 1 to determine the size of each range
	for {
		merged := 0
		for i := range freshIdRanges {
			for j := range freshIdRanges {
				if freshIdRanges[j].X > freshIdRanges[i].X && freshIdRanges[j].X <= freshIdRanges[i].Y {
					freshIdRanges[j].X = freshIdRanges[i].X
					merged++
				}
				if freshIdRanges[j].Y >= freshIdRanges[i].X && freshIdRanges[j].Y < freshIdRanges[i].Y {
					freshIdRanges[j].Y = freshIdRanges[i].Y
					merged++
				}
			}
		}
		fmt.Println(freshIdRanges)

		if merged == 0 {
			break
		}
	}

	freshIdRanges = aocutil.Unique(freshIdRanges)

	freshIdAmount := int64(0)

	for _, rang := range freshIdRanges {
		freshIdAmount += int64(rang.Y-rang.X) + 1
	}

	fmt.Println("Fresh ingredients: ", freshIdAmount)
}
