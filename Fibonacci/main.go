package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

// Pure recursive Fibonacci
func fibonacciRecursive(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// Warm-up to stabilize performance
func warmUp() {
	for i := 0; i < 100000; i++ {
		_ = 3 * 3
	}
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
}

func measureMemory(p *process.Process, memChan chan float64, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			memMB := float64(memStats.Sys) / (1024.0 * 1024.0) // Use Sys for total memory
			select {
			case memChan <- memMB:
			default:
			}
			time.Sleep(10 * time.Millisecond) // Poll every 10ms
		}
	}
}

func main() {
	n := 35
	numRuns := 5 // Number of runs to average measurements

	// Warm-up CPU and runtime
	warmUp()

	// Prepare process info
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		fmt.Println("Error getting process info:", err)
		return
	}

	var totalWallTime, totalCPUTime float64
	var maxMemory float64

	for run := 0; run < numRuns; run++ {
		// Force GC and collect initial memory stats
		var memStats runtime.MemStats
		runtime.GC()
		runtime.Gosched()
		time.Sleep(100 * time.Millisecond) // Allow GC to finish
		runtime.ReadMemStats(&memStats)
		memBefore := float64(memStats.Sys) / (1024.0 * 1024.0) // in MB

		// Start memory polling
		memChan := make(chan float64, 100)
		done := make(chan struct{})
		go measureMemory(p, memChan, done)

		// Get CPU times before
		cpuBefore, err := p.Times()
		if err != nil {
			fmt.Println("Error getting CPU time:", err)
			return
		}

		// Get wall time before
		startWall := time.Now()

		// Execute the algorithm
		result := fibonacciRecursive(n)

		// Measure elapsed wall time
		elapsedWall := time.Since(startWall)

		// Get CPU times after
		runtime.Gosched()
		cpuAfter, err := p.Times()
		if err != nil {
			fmt.Println("Error getting CPU time:", err)
			return
		}
		cpuTime := (cpuAfter.User + cpuAfter.System) - (cpuBefore.User + cpuBefore.System)

		// Stop memory polling
		close(done)

		// Collect max memory from polling
		localMaxMemory := memBefore
		for len(memChan) > 0 {
			if mem := <-memChan; mem > localMaxMemory {
				localMaxMemory = mem
			}
		}

		// Read memory after execution
		runtime.GC()
		time.Sleep(100 * time.Millisecond)
		runtime.ReadMemStats(&memStats)
		memAfter := float64(memStats.Sys) / (1024.0 * 1024.0)
		if memAfter > localMaxMemory {
			localMaxMemory = memAfter
		}

		// Update max memory across runs
		if localMaxMemory > maxMemory {
			maxMemory = localMaxMemory
		}

		// Accumulate times
		totalWallTime += elapsedWall.Seconds()
		totalCPUTime += cpuTime

		// Output result for this run
		fmt.Printf("Run %d: Fibonacci(%d) = %d\n", run+1, n, result)
	}

	// Calculate averages
	avgWallTime := totalWallTime / float64(numRuns)
	avgCPUTime := totalCPUTime / float64(numRuns)

	// Output benchmark results
	fmt.Printf("\nAverage Execution time (wall): %.6f seconds\n", avgWallTime)
	fmt.Printf("Average CPU time: %.6f seconds\n", avgCPUTime)
	fmt.Printf("Max memory used: %.2f MB\n", maxMemory)
}
