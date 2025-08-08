#include <iostream>
#include <fstream>
#include <string>
#include <chrono>
#include <windows.h>
#include <psapi.h>
#include <algorithm>
#include <cctype>

using namespace std;
using namespace std::chrono;

// Struct to hold memory info
struct MemoryUsage {
    size_t workingSetMB;
    size_t peakWorkingSetMB;
};

// Get current and peak memory usage
MemoryUsage getMemoryUsage() {
    PROCESS_MEMORY_COUNTERS memInfo;
    GetProcessMemoryInfo(GetCurrentProcess(), &memInfo, sizeof(memInfo));

    MemoryUsage usage;
    usage.workingSetMB = memInfo.WorkingSetSize / (1024 * 1024);
    usage.peakWorkingSetMB = memInfo.PeakWorkingSetSize / (1024 * 1024);
    return usage;
}

int main() {
    string inputFile = "C:\\Users\\c380b\\OneDrive\\Desktop\\CDS\\IO\\IMDB Dataset.csv";
    string outputFile = "C:\\Users\\c380b\\OneDrive\\Desktop\\CDS\\IO\\sentiment_cpp.txt";

    // Start profiling
    auto startWall = high_resolution_clock::now();
    clock_t cpuStart = clock();
    EmptyWorkingSet(GetCurrentProcess()); // Clear unused memory
    MemoryUsage memBefore = getMemoryUsage();

    ifstream inFile(inputFile);
    if (!inFile.is_open()) {
        cerr << "Failed to open file." << endl;
        return 1;
    }

    string line;
    int positiveCount = 0;
    int negativeCount = 0;

    // Skip header
    getline(inFile, line);

    while (getline(inFile, line)) {
        // Expecting format: "review","sentiment"
        size_t pos = line.rfind(',');
        if (pos != string::npos) {
            string sentiment = line.substr(pos + 1);

            // Clean up sentiment string
            sentiment.erase(remove(sentiment.begin(), sentiment.end(), '\"'), sentiment.end());
            sentiment.erase(remove_if(sentiment.begin(), sentiment.end(), ::isspace), sentiment.end());

            if (sentiment == "positive")
                ++positiveCount;
            else if (sentiment == "negative")
                ++negativeCount;
        }
    }

    inFile.close();

    // Write output
    ofstream outFile(outputFile);
    if (outFile.is_open()) {
        outFile << "Total positive reviews: " << positiveCount << "\n";
        outFile << "Total negative reviews: " << negativeCount << "\n";
        outFile.close();
    } else {
        cerr << "Failed to write to file." << endl;
        return 1;
    }

    // End profiling
    auto endWall = high_resolution_clock::now();
    clock_t cpuEnd = clock();
    EmptyWorkingSet(GetCurrentProcess());
    MemoryUsage memAfter = getMemoryUsage();

    // Calculate time and memory
    double wallTime = duration_cast<duration<double>>(endWall - startWall).count();
    double cpuTime = 1000.0 * (cpuEnd - cpuStart) / CLOCKS_PER_SEC;
    size_t usedMemoryMB = memAfter.workingSetMB > memBefore.workingSetMB
                          ? memAfter.workingSetMB - memBefore.workingSetMB
                          : 0;

    // Output report
    cout << "\nSentiment counts written to " << outputFile << endl;
    cout << "\n--- Profiling Report ---" << endl;
    cout << "Wall-clock time    : " << wallTime << " seconds" << endl;
    cout << "CPU time           : " << cpuTime << " ms" << endl;
    cout << "Used memory        : " << usedMemoryMB << " MB" << endl;
    cout << "Peak memory usage  : " << memAfter.peakWorkingSetMB << " MB" << endl;

    return 0;
}
