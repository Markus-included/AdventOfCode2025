package main

import (
	"AdventOfCode2025/Day04/common"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"runtime/debug"
)

func main() {
	input, _ := common.GetInput()

	fmt.Println("Part 1")
	fmt.Println("Input size:", len(input[0]), "x", len(input))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day04Part1")()
	accessibleRolls := 0
	for y, col := range input {
		_ = y
		for x, row := range col {
			_ = x
			if row != "@" {
				//fmt.Print(".")
				continue
			}

			adjRolls := 0
			for _, dir := range common.DIRS {
				ny := y + dir[0]
				nx := x + dir[1]
				if ny >= 0 && ny < len(input) && nx >= 0 && nx < len(col) {
					if input[ny][nx] == "@" {
						adjRolls++
					}
				}
			}
			if adjRolls < 4 {
				//fmt.Print("x")
				accessibleRolls++
			} else {
				//fmt.Print("@")
			}

		}
		//fmt.Println()
	}
	fmt.Println("Accessible rolls:", accessibleRolls)
}
