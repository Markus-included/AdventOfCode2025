package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	fmt.Println("Part 2")
	fmt.Println("Lines:", len(lines))

	dial := 50
	zeroPositions := 0

	for lineNum, line := range lines {
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}
		i, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to parse line: ", lineNum, " ", line)
		}
		oldDial := dial
		dial = (dial + (sign * i))
		for ; oldDial != dial; oldDial += sign {
			if oldDial%100 == 0 {
				zeroPositions += 1
			}
		}
		dial %= 100
	}

	fmt.Println("Passcode: ", zeroPositions)
}
