# 🧪 Comparative Profiling of Python, C++, Java, and Go  
*A Performance Benchmarking Study*

**📌 Author**: Bineesh Mathew  
**🏛️ Institution**: St. Xavier’s College, Mumbai  
**📅 Year**: 2025  

---

## 🎯 Objective

This project presents a **comparative performance analysis** of four widely-used programming languages — **Python, C++, Java, and Go** — by evaluating them against common software metrics using real-world computational tasks.

### Key Metrics Evaluated:
- ⏱️ **Execution Time**  
- 💾 **Memory Usage**  
- ⚙️ **Load Handling Stability**  
- 🐞 **Debugging Simplicity**

> 🔍 *Note: All implementations were kept minimal and framework-independent to maintain benchmarking fairness.*

---

## 🧪 Benchmarking Tasks

The benchmarking suite includes tasks designed to reflect realistic computational workloads:

| 🧮 Task                 | 🔍 Description                                                                 |
|------------------------|--------------------------------------------------------------------------------|
| **Fibonacci**          | Recursive computation of F(35); tests CPU load and recursion depth            |
| **File I/O**           | Read and write operations on a 100MB CSV file; tests disk interaction speed   |
| **Matrix Multiplication** | Multiplying two 1000×1000 matrices; evaluates CPU-bound computation         |
| **Load Testing**       | Sending 1000 HTTP GET requests; tests concurrency support and runtime stability|

---

## 🛠️ Tools & Profiling Methods

### 🐍 Python:
- `psutil` (system performance)
- `tracemalloc` (memory tracing)
- `time` module (execution timing)

### 💠 C++:
- `std::chrono` (execution time)
- `GetProcessMemoryInfo` (memory usage via Windows API)

### ☕ Java:
- `ThreadMXBean` (CPU time)
- `MemoryMXBean` (heap/non-heap memory stats)

### 🐹 Go:
- `runtime.MemStats` (memory profiling)
- `pprof` (advanced performance profiling)

---

## 📊 Key Findings

| **Metric**             | 🥇 **Best Performer** | 🐢 **Lowest Performer** |
|------------------------|----------------------|--------------------------|
| Execution Time         | **C++**              | **Python**               |
| Memory Efficiency      | **Java**             | **Python**               |
| Load Stability         | **C++**              | **Python**               |
| Debugging Simplicity   | **Python**           | **C++**                  |

### 📌 Summary

- **C++**: 🧠 Blazing fast and memory efficient, but difficult to debug and maintain  
- **Go**: ⚖️ Well-balanced in performance, concurrency, and ease of use  
- **Java**: 💼 Offers great stability and efficient memory management  
- **Python**: 🛠️ Developer-friendly and highly readable, but slower in raw execution

---

## 💻 System Configuration

| Specification       | Details                      |
|---------------------|------------------------------|
| **Operating System** | Windows 11 (64-bit)          |
| **Processor**        | AMD Ryzen 7 7435HS           |
| **RAM**              | 16 GB                        |
| **Language Versions**| Python 3.10, C++20 (MSVC), Java 17, Go 1.19 |

---

## 🚧 Limitations

- Benchmarks were executed on a **single hardware configuration** — results may differ across systems  
- **No performance-specific libraries** (e.g., NumPy, Boost) were used  
- **Concurrency models** were kept basic for cross-language uniformity

---

## 📚 Citation

> **Title**: *Comparative Profiling of Python and Three High-Performance Languages for Key Software Metrics*  
> **Author**: Bineesh Mathew  
> **Institution**: St. Xavier’s College, Mumbai (2025)

---

## 📬 Contact

For questions, suggestions, or collaboration opportunities:

- 📧 Email: [c380bineesh@gmail.com](mailto:c380bineesh@gmail.com)  
- 💻 GitHub: [github.com/ryuk38](https://github.com/ryuk38)
