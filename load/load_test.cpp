#include <iostream>
#include <vector>
#include <chrono>
#include <windows.h>
#include <winhttp.h>
#include <psapi.h>

#pragma comment(lib, "winhttp.lib")
#pragma comment(lib, "psapi.lib")

const int NUM_REQUESTS = 1000;
const int THREADS = 8; // Match typical CPU core count
const std::wstring TARGET_URL = L"https://jsonplaceholder.typicode.com/posts/1";
const DWORD STACK_SIZE = 128 * 1024; // 128 KB stack size per thread

// Global WinHTTP session and connection
HINTERNET hSession = nullptr;
HINTERNET hConnect = nullptr;
std::wstring hostName;
std::wstring urlPath;
LONG volatile success_count = 0; // Thread-safe success counter

bool initialize_winhttp() {
    hSession = WinHttpOpen(L"LoadTest/1.0", WINHTTP_ACCESS_TYPE_DEFAULT_PROXY,
                           WINHTTP_NO_PROXY_NAME, WINHTTP_NO_PROXY_BYPASS, 0);
    if (!hSession) return false;

    URL_COMPONENTS urlComp = { sizeof(URL_COMPONENTS) };
    urlComp.dwHostNameLength = -1;
    urlComp.dwUrlPathLength = -1;

    if (!WinHttpCrackUrl(TARGET_URL.c_str(), 0, 0, &urlComp)) {
        WinHttpCloseHandle(hSession);
        return false;
    }

    hostName = std::wstring(urlComp.lpszHostName, urlComp.dwHostNameLength);
    urlPath = std::wstring(urlComp.lpszUrlPath, urlComp.dwUrlPathLength);

    hConnect = WinHttpConnect(hSession, hostName.c_str(), INTERNET_DEFAULT_HTTPS_PORT, 0);
    if (!hConnect) {
        WinHttpCloseHandle(hSession);
        return false;
    }

    return true;
}

bool send_request() {
    HINTERNET hRequest = WinHttpOpenRequest(hConnect, L"GET", urlPath.c_str(),
                                            NULL, WINHTTP_NO_REFERER,
                                            WINHTTP_DEFAULT_ACCEPT_TYPES,
                                            WINHTTP_FLAG_SECURE);
    if (!hRequest) return false;

    WinHttpSetTimeouts(hRequest, 5000, 5000, 5000, 5000);

    bool success = false;

    if (WinHttpSendRequest(hRequest, WINHTTP_NO_ADDITIONAL_HEADERS, 0,
                           WINHTTP_NO_REQUEST_DATA, 0, 0, 0)) {
        if (WinHttpReceiveResponse(hRequest, NULL)) {
            DWORD statusCode = 0;
            DWORD size = sizeof(statusCode);
            WinHttpQueryHeaders(hRequest,
                                WINHTTP_QUERY_STATUS_CODE | WINHTTP_QUERY_FLAG_NUMBER,
                                NULL, &statusCode, &size, NULL);
            success = (statusCode == 200);
        }
    }

    WinHttpCloseHandle(hRequest);
    return success;
}

struct ThreadData {
    int num_requests;
};

DWORD WINAPI thread_proc(LPVOID param) {
    ThreadData* data = static_cast<ThreadData*>(param);
    for (int i = 0; i < data->num_requests; ++i) {
        if (send_request()) {
            InterlockedIncrement(&success_count);
        }
    }
    return 0;
}

long long get_memory_usage() {
    PROCESS_MEMORY_COUNTERS pmc;
    GetProcessMemoryInfo(GetCurrentProcess(), &pmc, sizeof(pmc));
    return pmc.PeakWorkingSetSize;
}

long long get_cpu_time() {
    FILETIME creationTime, exitTime, kernelTime, userTime;
    GetProcessTimes(GetCurrentProcess(), &creationTime, &exitTime, &kernelTime, &userTime);

    ULARGE_INTEGER kernel, user;
    kernel.LowPart = kernelTime.dwLowDateTime;
    kernel.HighPart = kernelTime.dwHighDateTime;
    user.LowPart = userTime.dwLowDateTime;
    user.HighPart = userTime.dwHighDateTime;

    // Return time in seconds
    return (kernel.QuadPart + user.QuadPart) / 10000000;
}

int main() {
    if (!initialize_winhttp()) {
        std::cerr << "Failed to initialize WinHTTP\n";
        return 1;
    }

    // Measure memory before
    EmptyWorkingSet(GetCurrentProcess());
    long long mem_before = get_memory_usage();

    // Start time
    auto start_time = std::chrono::high_resolution_clock::now();
    long long start_cpu = get_cpu_time();

    // Thread setup
    std::vector<HANDLE> threads(THREADS);
    std::vector<ThreadData> thread_data(THREADS);
    int requests_per_thread = NUM_REQUESTS / THREADS;
    int extra_requests = NUM_REQUESTS % THREADS;

    for (int i = 0; i < THREADS; ++i) {
        thread_data[i].num_requests = requests_per_thread + (i < extra_requests ? 1 : 0);
        threads[i] = CreateThread(NULL, STACK_SIZE, thread_proc, &thread_data[i], 0, NULL);
        if (!threads[i]) {
            std::cerr << "Failed to create thread " << i << "\n";
        }
    }

    WaitForMultipleObjects(THREADS, threads.data(), TRUE, INFINITE);
    for (int i = 0; i < THREADS; ++i) {
        CloseHandle(threads[i]);
    }

    // End time
    auto end_time = std::chrono::high_resolution_clock::now();
    long long end_cpu = get_cpu_time();
    EmptyWorkingSet(GetCurrentProcess());
    long long mem_after = get_memory_usage();

    WinHttpCloseHandle(hConnect);
    WinHttpCloseHandle(hSession);

    // Metrics
    double wall_time = std::chrono::duration<double>(end_time - start_time).count();
    double cpu_time = static_cast<double>(end_cpu - start_cpu);
    double max_memory = std::max(mem_before, mem_after) / (1024.0 * 1024.0);
    double throughput = wall_time > 0 ? NUM_REQUESTS / wall_time : 0;

    std::cout.setf(std::ios::fixed);
    std::cout.precision(6);
    std::cout << "\n--- Load Test Report ---\n";
    std::cout << "Total requests      : " << NUM_REQUESTS << "\n";
    std::cout << "Successful requests : " << success_count << "\n";
    std::cout << "Failed requests     : " << (NUM_REQUESTS - success_count) << "\n";
    std::cout << "Wall-clock time     : " << wall_time << " seconds\n";
    std::cout << "CPU time            : " << cpu_time << " seconds\n";
    std::cout << "Max memory          : " << max_memory << " MB\n";
    std::cout << "Throughput          : " << throughput << " requests/second\n";

    return 0;
}
