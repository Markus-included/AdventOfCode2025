package main

import (
	"AdventOfCode2025/Day07/common"
	"AdventOfCode2025/aocutil"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"runtime/debug"
	"strings"
)

func main() {
	input, _ := common.GetInput()

	fmt.Println("Part 1")
	fmt.Println("Input size:", len(input))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day07Part1")()

	_ = CountSplits(input, strings.Index(input[0], "S"), 0)

	for _, line := range input {
		fmt.Println(line)
	}

	fmt.Println("Splits: ", splits)
}

var (
	cache      = make(map[aocutil.Vector2[int]]int64)
	splits int = 0
)

func CountSplits(lines []string, x int, y int) (res int64) {
	//defer func() {
	//	for _, line := range lines {
	//		fmt.Println(line)
	//	}
	//}()

	pos := aocutil.Vector2[int]{X: x, Y: y}
	i, ok := cache[pos]

	if ok {

		fmt.Println()
		fmt.Println("CACHE HIT")
		fmt.Println()

		return i
	}
	if y == len(lines) {
		res = 1
		cache[pos] = res
		return
	}
	if lines[y][x] == '|' {
		res = 0
		return
	}
	if lines[y][x] == '^' {
		splits++
		res = CountSplits(lines, x-1, y) + CountSplits(lines, x+1, y)
		cache[pos] = res
		return
	} else {
		lines[y] = lines[y][:x] + "|" + lines[y][x+1:]
		res = CountSplits(lines, x, y+1)
		cache[pos] = res
		return
	}
}
