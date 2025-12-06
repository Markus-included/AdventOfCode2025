package main

import (
	"AdventOfCode2025/aocutil"
	"AdventOfCode2025/mprofiler"
	"fmt"
	"os"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
)

func GetInput() ([][]int, []string, error) {
	data, err := os.ReadFile("inputs/day06.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	input := string(data)

	operatorsStart := strings.LastIndex(strings.TrimSpace(input), "\n") + 1

	operators := aocutil.Filter(strings.Split(strings.TrimSpace(input[operatorsStart:]), " "), func(seg string) bool {
		return seg != ""
	})

	slices.Reverse(operators)

	grid := strings.Split(input[:operatorsStart], "\n")

	groups := make([][]int, len(operators))
	groupIdx := 0

	for x := len(grid[1]) - 1; x >= 0; x-- {
		s := ""
		for y := range grid[:len(grid)-1] {
			s += string(grid[y][x])
		}
		s = strings.TrimSpace(s)
		if s == "" {
			fmt.Println("Bazinga", operators[groupIdx])
			groupIdx++
		} else {
			n, _ := strconv.Atoi(s)
			groups[groupIdx] = append(groups[groupIdx], n)
			fmt.Println(s)
		}
	}

	fmt.Println(strings.Join(grid, ","))

	return groups, operators, nil
}

func main() {
	groups, operators, _ := GetInput()

	fmt.Println("Part 2")
	fmt.Println("Input size:", len(groups))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day06Part2")()

	grandTotal := 0

	for i, group := range groups {
		result := 0
		operator := operators[i]
		if operator == "*" {
			result = 1
		}
		fmt.Print(strings.Join(aocutil.Map(group, func(v int) string { return strconv.Itoa(v) }), operator))
		for _, n := range group {
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
