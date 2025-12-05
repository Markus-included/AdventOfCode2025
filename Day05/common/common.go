package common

import (
	"AdventOfCode2025/aocutil"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() (freshIds []aocutil.Vector2[int], idsToSearch []int, outerr error) {
	data, err := os.ReadFile("inputs/day05.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		outerr = err
		return
	}

	input := string(data)
	sections := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, rang := range strings.Split(sections[0], "\n") {
		splitRange := strings.Split(rang, "-")
		start, _ := strconv.Atoi(splitRange[0])
		end, _ := strconv.Atoi(splitRange[1])

		freshIds = append(freshIds, aocutil.Vector2[int]{X: start, Y: end})
	}

	//fmt.Println(freshIds)

	idsToSearch = []int{}
	idsToSearch = aocutil.Map(strings.Split(sections[1], "\n"), func(s string) int {
		x, err := strconv.Atoi(s)
		//fmt.Println(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Inval", s, err)
		}
		return x
	})

	//fmt.Println(idsToSearch)

	return
}
