# 🧪 Comparative Profiling of Python, C++, Java, and Go

This repository contains the source code and benchmarking setup used for a comparative analysis of four programming languages — **Python**, **C++**, **Java**, and **Go** — based on key software performance metrics.

📘 **Prepared by:** Bineesh Mathew  
🏫 **Institution:** St. Xavier’s College, Mumbai

---

## 📌 Objective

To compare the performance, memory usage, load stability, and debugging ease of four popular programming languages using benchmark tasks that simulate real-world scenarios.

---

## 📁 Benchmark Tasks

Each task was implemented in all four languages to ensure fair comparison:

| Task | Description |
|------|-------------|
| 🔢 **Fibonacci** | Recursive calculation of F(35) to test CPU usage and function call overhead |
| 💾 **I/O Operations** | Reading and writing a 100MB CSV file to test disk I/O performance |
| ➗ **Matrix Multiplication** | Multiplying two 1000×1000 matrices to evaluate numerical processing |
| 🌐 **Load Testing** | Sending 1000 HTTP requests to evaluate concurrency and runtime stability |

---

## 🧰 Tools Used

Each language used native profiling tools to capture:

- Execution Time
- CPU Time
- Peak Memory Usage
- Stability Under Load

Examples include:  
- `psutil`, `tracemalloc`, `time` (Python)  
- `std::chrono`, `GetProcessMemoryInfo` (C++)  
- `ThreadMXBean`, `MemoryMXBean` (Java)  
- `runtime.MemStats`, `pprof` (Go)

---

## ✅ Key Findings

| Metric | Best Performer | Worst Performer |
|--------|----------------|------------------|
| Execution Time | C++ | Python |
| Memory Usage | Java | Python |
| Load Stability | C++ | Python |
| Debugging Ease | Python | C++ |

- **C++**: Fastest and most efficient, but complex to debug.
- **Go**: Balanced in performance, concurrency, and simplicity.
- **Java**: Most memory-efficient and stable across tasks.
- **Python**: Simplest to debug and develop with, but least performant.

---

## 🖥️ System Configuration

- **Operating System**: Windows 11 (64-bit)  
- **Processor**: AMD Ryzen 7 7435HS  
- **RAM**: 16 GB  
- **Languages Used**:
  - Python 3.10
  - C++20 (MSVC)
  - Java 17
  - Go 1.19

---

## 📂 Repository Structure
/Fibonacci -> Recursive Fibonacci in all languages
/IO -> I/O benchmarking scripts
/Matrix -> Matrix multiplication programs
/LoadTest -> Load testing code for HTTP requests
README.md -> This file


---

## 🚧 Limitations

- Benchmarks were performed on a single system and may vary across different hardware.
- Naive implementations were used for fairness; no language-specific libraries (e.g., NumPy, OpenBLAS) were applied.
- Multi-threading and concurrency models were tested, but full optimization was not the focus.

---

## 📚 Citation

**Project Title**: *Comparative Profiling of Python and Three High-Performance Languages for Key Software Metrics*  
**Author**: Bineesh Mathew  
**Institution**: St. Xavier’s College, Mumbai (2025)

---

## 📬 Contact

For queries or feedback:  
📧 Email: [c380bineesh@gmail.com]  
🔗 GitHub: [github.com/ryuk38](https://github.com/ryuk38)

---


