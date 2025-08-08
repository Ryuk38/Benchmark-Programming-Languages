#include <iostream>
#include <chrono>
#include <windows.h>
#include <psapi.h>

long long fibonacci_recursive(int n) {
    if (n <= 1) return n;
    return fibonacci_recursive(n-1) + fibonacci_recursive(n-2);
}

int main() {
    const int n = 35;
    
    // Measure initial memory
    PROCESS_MEMORY_COUNTERS memInfo;
    GetProcessMemoryInfo(GetCurrentProcess(), &memInfo, sizeof(memInfo));
    double mem_before = memInfo.PeakWorkingSetSize / (1024.0 * 1024.0); // MB
    
    // Compute Fibonacci with naive recursion
    auto start = std::chrono::high_resolution_clock::now();
    clock_t start_cpu = clock();
    long long result = fibonacci_recursive(n);
    clock_t end_cpu = clock();
    auto end = std::chrono::high_resolution_clock::now();
    
    // Measure final memory
    GetProcessMemoryInfo(GetCurrentProcess(), &memInfo, sizeof(memInfo));
    double mem_after = memInfo.PeakWorkingSetSize / (1024.0 * 1024.0); // MB
    double max_memory = std::max(mem_before, mem_after);
    
    // Output metrics
    std::chrono::duration<double> diff = end - start;
    std::cout << "Execution time: " << diff.count() << " seconds\n";
    std::cout << "CPU time: " << (end_cpu - start_cpu) / (double)CLOCKS_PER_SEC << " seconds\n";
    std::cout << "Max memory: " << max_memory << " MB\n";
    std::cout << "Fibonacci(" << n << ") = " << result << "\n";
    return 0;
}