package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/process"
)

const (
	totalRequests = 1000
	concurrency   = 100 // Run 100 concurrent workers
	url           = "https://jsonplaceholder.typicode.com/posts/1"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores

	var wg sync.WaitGroup
	requests := make(chan struct{}, totalRequests)
	successCount := 0
	failureCount := 0
	var mu sync.Mutex

	// Track wall and CPU time
	start := time.Now()
	processInfo, _ := process.NewProcess(int32(os.Getpid()))
	cpuStart, _ := processInfo.Times()

	// Fill the channel with total requests
	for i := 0; i < totalRequests; i++ {
		requests <- struct{}{}
	}
	close(requests)

	// Worker pool
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{}

			for range requests {
				resp, err := client.Get(url)
				if err != nil {
					mu.Lock()
					failureCount++
					mu.Unlock()
					continue
				}

				// Do NOT read the entire body unless needed
				// Just close it immediately
				resp.Body.Close()

				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	cpuEnd, _ := processInfo.Times()
	memoryInfo, _ := processInfo.MemoryInfo()

	fmt.Println("--- Load Test Report ---")
	fmt.Println("Total requests     :", totalRequests)
	fmt.Println("Successful requests:", successCount)
	fmt.Println("Failed requests    :", failureCount)
	fmt.Printf("Wall-clock time    : %.3f seconds\n", elapsed.Seconds())
	fmt.Printf("CPU time           : %.3f seconds\n", cpuEnd.User-cpuStart.User)
	fmt.Printf("Memory used        : %.2f MB\n", float64(memoryInfo.RSS)/(1024*1024))
}
