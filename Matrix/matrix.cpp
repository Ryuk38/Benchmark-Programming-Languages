#include <iostream>
#include <vector>
#include <random>
#include <chrono>
#include <windows.h>
#include <psapi.h>

void matrix_multiply() {
    const int n = 1000;
    std::vector<std::vector<double>> A(n, std::vector<double>(n));
    std::vector<std::vector<double>> B(n, std::vector<double>(n));
    std::vector<std::vector<double>> C(n, std::vector<double>(n, 0.0));

    std::mt19937 gen(42);
    std::uniform_real_distribution<> dis(0.0, 1.0);

    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++) {
            A[i][j] = dis(gen);
            B[i][j] = dis(gen);
        }

    PROCESS_MEMORY_COUNTERS memInfo;
    GetProcessMemoryInfo(GetCurrentProcess(), &memInfo, sizeof(memInfo));
    double memBefore = memInfo.WorkingSetSize / (1024.0 * 1024.0);

    auto start = std::chrono::high_resolution_clock::now();
    clock_t cpuStart = clock();

    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            for (int k = 0; k < n; k++)
                C[i][j] += A[i][k] * B[k][j];

    clock_t cpuEnd = clock();
    auto end = std::chrono::high_resolution_clock::now();
    GetProcessMemoryInfo(GetCurrentProcess(), &memInfo, sizeof(memInfo));
    double memAfter = memInfo.WorkingSetSize / (1024.0 * 1024.0);

    std::cout << "Execution time: " << std::chrono::duration<double>(end - start).count() << " s\n";
    std::cout << "CPU time: " << (cpuEnd - cpuStart) / (double)CLOCKS_PER_SEC << " s\n";
    std::cout << "Memory used: " << std::max(memBefore, memAfter) << " MB\n";
}

int main() {
    matrix_multiply();
    return 0;
}
