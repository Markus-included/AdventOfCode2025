package main

import (
	"AdventOfCode2025/mprofiler"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

func Atoi64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func main() {
	data, err := os.ReadFile("inputs/day02.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(data)
	rows := strings.Split(strings.TrimSpace(input), ",")

	var ids []int64

	for _, row := range rows {
		rang := strings.Split(row, "-")

		start, _ := Atoi64(rang[0])
		end, _ := Atoi64(rang[1])

		for i := start; i <= end; i++ {
			// I am well aware that this is hella slow, TODO: optimize
			ids = append(ids, i)
		}

	}

	fmt.Println("Part 1")
	fmt.Println("IDs:", len(ids))

	debug.SetGCPercent(-1)
	defer mprofiler.TrackExecutionTime("Day02Part1")()

	var sum int64 = 0

	for _, id := range ids {
		strId := strconv.FormatInt(id, 10)
		if len(strId)%2 != 0 {
			continue
		}
		if strId[:len(strId)/2] == strId[(len(strId)/2):] {
			fmt.Println(strId[:len(strId)/2], strId[(len(strId)/2):])
			sum += id
		}
	}

	fmt.Println("Passcode: ", sum)
}
