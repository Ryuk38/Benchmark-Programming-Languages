# 🧪 Comparative Profiling of Python, C++, Java, and Go

A performance benchmarking study of four popular programming languages — Python, C++, Java, and Go — across key software metrics using real-world computational tasks.

Author: Bineesh Mathew  
Institution: St. Xavier’s College, Mumbai  
Year: 2025

---

## 🎯 Objective

This project aims to evaluate and compare:

- Execution Time  
- Memory Usage  
- Load Handling Stability  
- Debugging Simplicity  

All benchmarks were implemented using simple baseline logic to ensure fairness, without any specialized libraries or frameworks.

---

## 🧪 Benchmark Tasks

Each task is designed to reflect common computational and I/O operations:

| Task                | Description                                                                 |
|---------------------|-----------------------------------------------------------------------------|
| Fibonacci           | Recursive calculation of F(35); tests CPU usage and function call overhead |
| File I/O            | Read/write a 100MB CSV file; tests disk performance                         |
| Matrix Multiplication | Multiply two 1000×1000 matrices; evaluates processing power               |
| Load Testing        | Send 1000 HTTP requests; tests concurrency and runtime stability            |

---

## 🛠️ Profiling Tools Used

Language-specific profiling tools were used to measure performance:

- Python: `psutil`, `tracemalloc`, `time`
- C++: `std::chrono`, `GetProcessMemoryInfo`
- Java: `ThreadMXBean`, `MemoryMXBean`
- Go: `runtime.MemStats`, `pprof`

---

## 📊 Key Results

| Metric              | Best Performer | Lowest Performer |
|---------------------|----------------|------------------|
| Execution Time      | C++            | Python           |
| Memory Efficiency   | Java           | Python           |
| Load Stability      | C++            | Python           |
| Debugging Simplicity| Python         | C++              |

### Summary

- **C++**: Fastest and most memory-efficient, but complex to debug  
- **Go**: Offers a good balance between speed, concurrency, and simplicity  
- **Java**: Very stable and memory-efficient  
- **Python**: Easiest to develop and debug with, but slowest in performance

---

## 💻 System Configuration

- OS: Windows 11 (64-bit)  
- Processor: AMD Ryzen 7 7435HS  
- RAM: 16 GB  
- Language Versions:
  - Python 3.10
  - C++20 (MSVC)
  - Java 17
  - Go 1.19

---

## 📁 Repository Structure
├── Fibonacci/ # Recursive Fibonacci in all 4 languages
├── IO/ # File I/O benchmarking scripts
├── Matrix/ # Matrix multiplication code
├── LoadTest/ # Load testing with 1000 HTTP requests
├── Fibonacci.ipynb # Jupyter notebook 
├── IO.ipynb
├── Matrix.ipynb
├── LoadTest.ipynb
└── README.md

Each `.ipynb` notebook contains all the codes together 

---

## 🚧 Limitations

- Tests were run on a single machine — results may vary with different hardware  
- No performance-specific libraries (like NumPy or Boost) were used  
- Concurrency models were kept simple for consistency across languages

---

## 📚 Citation

**Project Title**: Comparative Profiling of Python and Three High-Performance Languages for Key Software Metrics  
**Author**: Bineesh Mathew  
**Institution**: St. Xavier’s College, Mumbai (2025)

---

## 📬 Contact

For feedback or collaboration:

- Email: [c380bineesh@gmail.com](mailto:c380bineesh@gmail.com)  
- GitHub: [github.com/ryuk38](https://github.com/ryuk38)
