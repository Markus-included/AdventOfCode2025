package mprofiler

import (
	"fmt"
	"math/bits"
	"runtime"
	"time"
)

func TrackExecutionTime(name string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Printf("[%s] took %s\n", name, elapsed)

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %s\n", FormatBinary(m.Alloc))
		fmt.Printf("TotalAlloc = %s\n", FormatBinary(m.TotalAlloc))
		fmt.Printf("Sys = %s\n", FormatBinary(m.Sys))
		fmt.Printf("NumGC = %d\n", m.NumGC)
	}
}

func FormatBinary(bytes uint64) string {

	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	}

	suffixes := []string{"KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}

	index := int((bits.Len64(bytes)-1)/10) - 1
	if index < 0 {
		index = 0
	}
	if index >= len(suffixes) {
		index = len(suffixes) - 1
	}

	unit := uint64(1) << (10 * (index + 1)) // 1<<10 for KiB, 1<<20 for MiB, ...
	val := float64(bytes) / float64(unit)
	return fmt.Sprintf("%.2f %s", val, suffixes[index])
}
