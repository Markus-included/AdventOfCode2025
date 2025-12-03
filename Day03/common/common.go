package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() ([][]int8, error) {
	data, err := os.ReadFile("inputs/day03.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	height := len(lines)
	width := len(lines[0])

	digits := make([][]int8, height)

	for y := 0; y < height; y++ {
		digits[y] = make([]int8, width)
		for x := 0; x < width; x++ {
			val, _ := strconv.Atoi(lines[y][x : x+1])
			digits[y][x] = int8(val)
		}
	}

	return digits, nil
}

func LargestIdxInArray(array []int8) int {
	idx := 0
	largest := array[0]
	for i, value := range array {
		if value > largest {
			largest = value
			idx = i
		}
	}
	return idx
}
