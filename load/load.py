import requests
import time
import psutil
import os
from concurrent.futures import ThreadPoolExecutor
from multiprocessing import cpu_count

NUM_REQUESTS = 1000
THREADS = min(100, cpu_count() * 2)  # Cap threads to avoid resource exhaustion
TARGET_URL = "https://jsonplaceholder.typicode.com/posts/1"

def send_request():
    try:
        response = requests.get(TARGET_URL, timeout=5)
        return response.status_code == 200
    except requests.RequestException:
        return False

def main():
    # Measure initial memory
    process = psutil.Process(os.getpid())
    mem_before = process.memory_info().rss / 1024 / 1024  # MB
    
    # Send requests concurrently
    start_time = time.time()
    start_cpu = time.process_time()
    with ThreadPoolExecutor(max_workers=THREADS) as executor:
        results = list(executor.map(lambda _: send_request(), range(NUM_REQUESTS)))
    end_cpu = time.process_time()
    end_time = time.time()
    
    # Measure final memory
    mem_after = process.memory_info().rss / 1024 / 1024  # MB
    max_memory = max(mem_before, mem_after)
    
    # Calculate metrics
    success_count = sum(1 for r in results if r)
    failure_count = NUM_REQUESTS - success_count
    wall_time = end_time - start_time
    throughput = NUM_REQUESTS / wall_time if wall_time > 0 else 0
    
    # Output metrics
    print("\n--- Load Test Report ---")
    print(f"Total requests      : {NUM_REQUESTS}")
    print(f"Successful requests : {success_count}")
    print(f"Failed requests     : {failure_count}")
    print(f"Wall-clock time    : {wall_time:.2f} seconds")
    print(f"CPU time           : {end_cpu - start_cpu:.2f} seconds")
    print(f"Max memory         : {max_memory:.2f} MB")
    print(f"Throughput         : {throughput:.2f} requests/second")

if __name__ == "__main__":
    main()