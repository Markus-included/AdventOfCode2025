package mprofiler

import (
	"fmt"
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
		fmt.Printf("Alloc = %d bytes\n", m.Alloc)
		fmt.Printf("TotalAlloc = %d bytes\n", m.TotalAlloc)
		fmt.Printf("Sys = %d bytes\n", m.Sys)
		fmt.Printf("NumGC = %d\n", m.NumGC)
	}
}
