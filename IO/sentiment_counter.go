package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

// Warm-up to stabilize CPU and runtime
func warmUp() {
	for i := 0; i < 100000; i++ {
		_ = 3 * 3
	}
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
}

// Process file and return counts
func processFile(inputFile string) (int, int, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var positiveCount, negativeCount int
	scanner := bufio.NewScanner(file)
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			firstLine = false
			continue
		}
		if strings.HasSuffix(line, ",positive") {
			positiveCount++
		} else if strings.HasSuffix(line, ",negative") {
			negativeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("error reading file: %v", err)
	}
	return positiveCount, negativeCount, nil
}

// Poll memory usage
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
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	// File paths
	inputFile := "C:\\Users\\c380b\\OneDrive\\Desktop\\CDS\\IO\\IMDB Dataset.csv"
	outputFile := "C:\\Users\\c380b\\OneDrive\\Desktop\\CDS\\IO\\sentiment_go.txt"
	cpuProfile := "C:\\Users\\c380b\\OneDrive\\Desktop\\CDS\\IO\\cpu.prof"
	const numRuns = 5 // Number of runs for averaging

	// Warm-up
	warmUp()

	// Create CPU profile file
	cpuFile, err := os.Create(cpuProfile)
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer cpuFile.Close()

	// Initialize process for CPU time
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		fmt.Println("Error getting process info:", err)
		return
	}

	var totalWallTime, totalCPUTime float64
	var maxMemory float64
	var positiveCount, negativeCount int

	for run := 0; run < numRuns; run++ {
		// Start memory polling
		memChan := make(chan float64, 100)
		done := make(chan struct{})
		go measureMemory(p, memChan, done)

		// Start CPU profiling (only for first run to avoid overwriting)
		if run == 0 {
			pprof.StartCPUProfile(cpuFile)
		}

		// Start wall time and CPU time
		startWallTime := time.Now()
		cpuBefore, err := p.Times()
		if err != nil {
			fmt.Println("Error getting CPU time:", err)
			return
		}
		var memStart runtime.MemStats
		runtime.ReadMemStats(&memStart)
		memBefore := float64(memStart.Sys) / (1024.0 * 1024.0)

		// Process file
		positive, negative, err := processFile(inputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		if run == numRuns-1 {
			positiveCount = positive
			negativeCount = negative
		}

		// End measurements
		endWallTime := time.Now()
		cpuAfter, err := p.Times()
		if err != nil {
			fmt.Println("Error getting CPU time:", err)
			return
		}
		var memEnd runtime.MemStats
		runtime.ReadMemStats(&memEnd)
		memAfter := float64(memEnd.Sys) / (1024.0 * 1024.0)

		// Stop memory polling
		close(done)

		// Stop CPU profiling after first run
		if run == 0 {
			pprof.StopCPUProfile()
		}

		// Calculate times
		totalWallTime += endWallTime.Sub(startWallTime).Seconds()
		totalCPUTime += (cpuAfter.User + cpuAfter.System) - (cpuBefore.User + cpuBefore.System)

		// Update max memory
		localMaxMemory := memBefore
		if memAfter > localMaxMemory {
			localMaxMemory = memAfter
		}
		for len(memChan) > 0 {
			if mem := <-memChan; mem > localMaxMemory {
				localMaxMemory = mem
			}
		}
		if localMaxMemory > maxMemory {
			maxMemory = localMaxMemory
		}
	}

	// Write results to file (only once)
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()
	fmt.Fprintf(out, "Total positive reviews: %d\n", positiveCount)
	fmt.Fprintf(out, "Total negative reviews: %d\n", negativeCount)
	fmt.Println("Sentiment counts written to", outputFile)

	// Calculate averages
	avgWallTime := totalWallTime / float64(numRuns)
	avgCPUTime := totalCPUTime / float64(numRuns)

	// Profiling Report
	fmt.Println("\n--- Profiling Report (Averaged over", numRuns, "runs) ---")
	fmt.Printf("Average Wall-clock time : %.4f seconds\n", avgWallTime)
	fmt.Printf("Average CPU time       : %.4f seconds\n", avgCPUTime)
	fmt.Printf("Peak memory usage      : %.2f MB\n", maxMemory)
	fmt.Printf("CPU profile saved      : %s\n", cpuProfile)
}
