import time
import psutil
import os

def fibonacci_recursive(n):
    if n <= 1:
        return n
    return fibonacci_recursive(n-1) + fibonacci_recursive(n-2)

def main():
    n = 35
    
    # Measure initial memory
    process = psutil.Process(os.getpid())
    mem_before = process.memory_info().rss / 1024 / 1024  # MB
    
    # Compute Fibonacci with naive recursion
    start_time = time.time()
    start_cpu = time.process_time()
    result = fibonacci_recursive(n)
    end_cpu = time.process_time()
    end_time = time.time()
    
    # Measure final memory
    mem_after = process.memory_info().rss / 1024 / 1024  # MB
    max_memory = max(mem_before, mem_after)
    
    # Output metrics
    print(f"Execution time: {end_time - start_time:.2f} seconds")
    print(f"CPU time: {end_cpu - start_cpu:.2f} seconds")
    print(f"Max memory: {max_memory:.2f} MB")
    print(f"Fibonacci({n}) = {result}")

if __name__ == "__main__":
    main()