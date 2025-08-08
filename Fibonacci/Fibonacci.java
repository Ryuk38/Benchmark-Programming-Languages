
import java.lang.management.*;

public class Fibonacci {
    public static long fibonacciRecursive(int n) {
        if (n <= 1) return n;
        return fibonacciRecursive(n-1) + fibonacciRecursive(n-2);
    }
    
    public static void main(String[] args) {
        int n = 35;
        
        // Measure initial memory
        Runtime runtime = Runtime.getRuntime();
        MemoryMXBean memoryBean = ManagementFactory.getMemoryMXBean();
        runtime.gc();
        long mem_before = runtime.totalMemory() - runtime.freeMemory();
        long non_heap_before = memoryBean.getNonHeapMemoryUsage().getUsed();
        
        // Compute Fibonacci with naive recursion
        ThreadMXBean bean = ManagementFactory.getThreadMXBean();
        long start_cpu = bean.isCurrentThreadCpuTimeSupported() ? bean.getCurrentThreadCpuTime() : 0;
        long start_time = System.nanoTime();
        long result = fibonacciRecursive(n);
        long end_time = System.nanoTime();
        long end_cpu = bean.isCurrentThreadCpuTimeSupported() ? bean.getCurrentThreadCpuTime() : 0;
        
        // Measure final memory
        runtime.gc();
        long mem_after = runtime.totalMemory() - runtime.freeMemory();
        long non_heap_after = memoryBean.getNonHeapMemoryUsage().getUsed();
        
        // Calculate memory usage
        long heap_memory = Math.max(mem_before, mem_after) / (1024 * 1024); // Heap in MB
        long non_heap_memory = Math.max(non_heap_before, non_heap_after) / (1024 * 1024); // Non-heap in MB
        long total_memory = (Math.max(mem_before, mem_after) + Math.max(non_heap_before, non_heap_after)) / (1024 * 1024); // Total in MB
        
        // Output metrics
        System.out.printf("Execution time: %.4f seconds\n", (end_time - start_time) / 1e9);
        System.out.printf("CPU time: %.4f seconds\n", (end_cpu - start_cpu) / 1e9);
        System.out.printf("Heap memory used: %d MB\n", heap_memory);
        System.out.printf("Non-heap memory used: %d MB\n", non_heap_memory);
        System.out.printf("Total memory used: %d MB\n", total_memory);
        System.out.println("Fibonacci(" + n + ") = " + result);
    }
}
