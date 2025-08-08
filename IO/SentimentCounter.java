import java.io.*;
import java.lang.management.*;
import java.nio.file.*;
import java.util.*;

public class SentimentCounter {
    public static void main(String[] args) {
        // Set file paths
        String inputFile = "C:\\Users\\c380b\\OneDrive\\Desktop\\CDS\\IO\\IMDB Dataset.csv";
        String outputFile = "C:\\Users\\c380b\\OneDrive\\Desktop\\CDS\\IO\\sentiment_java.txt";

        int positiveCount = 0;
        int negativeCount = 0;

        // Start time and memory tracking
        long startWallTime = System.nanoTime();
        long startCpuTime = getCpuTime();
        // Force GC and measure initial memory
        System.gc();
        long startMemory = getUsedMemory();
        long startNonHeapMemory = getNonHeapMemory();

        try (BufferedReader reader = new BufferedReader(
                Files.newBufferedReader(Paths.get(inputFile)), 16 * 1024)) { // Increased buffer size
            String line;
            boolean firstLine = true;

            while ((line = reader.readLine()) != null) {
                if (firstLine) { // Skip header
                    firstLine = false;
                    continue;
                }

                // Minimize string operations
                if (line.endsWith(",positive")) {
                    positiveCount++;
                } else if (line.endsWith(",negative")) {
                    negativeCount++;
                }

                // Suggest GC periodically to manage temporary strings
                if ((positiveCount + negativeCount) % 10000 == 0) {
                    System.gc();
                }
            }

            // Write output efficiently
            List<String> outputLines = Arrays.asList(
                "Total positive reviews: " + positiveCount,
                "Total negative reviews: " + negativeCount
            );
            Files.write(Paths.get(outputFile), outputLines, StandardOpenOption.CREATE);
            System.out.println("Sentiment counts written to " + outputFile);

        } catch (IOException e) {
            System.err.println("Error reading or writing file: " + e.getMessage());
        }

        // End time and memory tracking
        // Force GC before final measurement
        System.gc();
        long endWallTime = System.nanoTime();
        long endCpuTime = getCpuTime();
        long endMemory = getUsedMemory();
        long endNonHeapMemory = getNonHeapMemory();

        // Display profiling report
        System.out.println("\n--- Profiling Report ---");
        System.out.printf("Wall-clock time       : %.4f seconds\n", (endWallTime - startWallTime) / 1e9);
        System.out.printf("CPU time              : %.4f seconds\n", (endCpuTime - startCpuTime) / 1e9);
        System.out.printf("Heap memory used      : %.2f MB\n", (endMemory - startMemory) / (1024.0 * 1024));
        System.out.printf("Non-heap memory used  : %.2f MB\n", (endNonHeapMemory - startNonHeapMemory) / (1024.0 * 1024));
        System.out.printf("Total memory used     : %.2f MB\n", 
            ((endMemory - startMemory) + (endNonHeapMemory - startNonHeapMemory)) / (1024.0 * 1024));
    }

    // Get CPU time using ManagementFactory
    private static long getCpuTime() {
        ThreadMXBean bean = ManagementFactory.getThreadMXBean();
        return bean.isCurrentThreadCpuTimeSupported() ? bean.getCurrentThreadCpuTime() : 0L;
    }

    // Get currently used heap memory
    private static long getUsedMemory() {
        Runtime runtime = Runtime.getRuntime();
        return runtime.totalMemory() - runtime.freeMemory();
    }

    // Get non-heap memory usage (e.g., metaspace)
    private static long getNonHeapMemory() {
        MemoryMXBean memoryBean = ManagementFactory.getMemoryMXBean();
        return memoryBean.getNonHeapMemoryUsage().getUsed();
    }
}