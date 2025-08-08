package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func matrixMultiply() {
	const n = 1000
	const blockSize = 64 // Block size for tiling, adjustable (32, 64, 128)

	// Initialize matrices
	A := make([][]float64, n)
	B := make([][]float64, n)
	C := make([][]float64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		B[i] = make([]float64, n)
		C[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			A[i][j] = rand.Float64()
			B[i][j] = rand.Float64()
		}
	}

	// Memory measurement before
	var memStats runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	memBefore := float64(memStats.Alloc) / (1024 * 1024)

	// Start timing
	start := time.Now()

	// Tiled matrix multiplication
	for i := 0; i < n; i += blockSize {
		for j := 0; j < n; j += blockSize {
			for k := 0; k < n; k += blockSize {
				// Process block
				for ii := i; ii < min(i+blockSize, n); ii++ {
					for jj := j; jj < min(j+blockSize, n); jj++ {
						for kk := k; kk < min(k+blockSize, n); kk++ {
							C[ii][jj] += A[ii][kk] * B[kk][jj]
						}
					}
				}
			}
		}
	}

	// End timing and memory measurement
	duration := time.Since(start)
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	memAfter := float64(memStats.Alloc) / (1024 * 1024)

	// Output results
	fmt.Printf("Execution time: %.3f s\n", duration.Seconds())
	fmt.Printf("Memory used: %.2f MB\n", max(memBefore, memAfter))
	fmt.Println("CPU time: See 'cpu.prof' file using `go tool pprof`")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	matrixMultiply()
}
