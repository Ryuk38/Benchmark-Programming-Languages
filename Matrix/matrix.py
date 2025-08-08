import time
import os
import psutil
import random

def matrix_multiply(n=1000):
    process = psutil.Process(os.getpid())

    # Generate 1000×1000 matrices A and B with random values
    A = [[random.random() for _ in range(n)] for _ in range(n)]
    B = [[random.random() for _ in range(n)] for _ in range(n)]
    C = [[0.0] * n for _ in range(n)]

    # Record initial memory (MB)
    memory_before = process.memory_info().rss / (1024 * 1024)

    # Start timing
    wall_start = time.time()
    cpu_start = time.process_time()

    # Optimized matrix multiplication using i-k-j loop order
    for i in range(n):
        for k in range(n):
            temp = A[i][k]
            for j in range(n):
                C[i][j] += temp * B[k][j]

    # End timing
    wall_end = time.time()
    cpu_end = time.process_time()

    # Record final memory (MB)
    memory_after = process.memory_info().rss / (1024 * 1024)
    peak_memory = max(memory_before, memory_after)

    # Output results
    print("\n--- Matrix Multiplication Report ---")
    print(f"Matrix size        : {n} x {n}")
    print(f"Execution time     : {wall_end - wall_start:.3f} seconds")
    print(f"CPU time           : {cpu_end - cpu_start:.3f} seconds")
    print(f"Peak memory usage  : {peak_memory:.2f} MB")

if __name__ == "__main__":
    matrix_multiply()
