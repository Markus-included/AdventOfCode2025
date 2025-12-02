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

	fmt.Println("Part 2")
	fmt.Println("IDs:", len(ids))

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(4 * 1024 * 1024 * 1024)
	defer mprofiler.TrackExecutionTime("Day02Part1")()

	var sum int64 = 0

	for _, id := range ids {
		strId := strconv.FormatInt(id, 10)
		//if len(strId)%2 != 0 {
		//	continue
		//}
		for chunkSz := 1; chunkSz <= len(strId)/2; chunkSz++ {
			chunked := ChunkSlice(strId, chunkSz)
			if AllSameStrings(chunked) {
				//fmt.Println(id)
				sum += id
				goto outer
			}
		}
	outer:
	}

	fmt.Println("Passcode: ", sum)
}

func ChunkSlice(slice string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func AllSameStrings(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
