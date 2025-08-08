package load;

import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.net.URI;
import java.time.Duration;
import java.util.concurrent.*;
import java.util.*;
import java.lang.management.ManagementFactory;
import java.lang.management.ThreadMXBean;

public class LoadTest {
    private static final int NUM_REQUESTS = 1000;
    private static final int THREADS = 100;
    private static final String TARGET_URL = "https://jsonplaceholder.typicode.com/posts/1";

    public static void main(String[] args) throws Exception {
        // Initialize HttpClient
        HttpClient client = HttpClient.newBuilder()
                .connectTimeout(Duration.ofSeconds(5))
                .followRedirects(HttpClient.Redirect.ALWAYS)
                .build();

        // Memory before
        Runtime runtime = Runtime.getRuntime();
        runtime.gc();
        long memBefore = runtime.totalMemory() - runtime.freeMemory();

        // CPU time before
        ThreadMXBean bean = ManagementFactory.getThreadMXBean();
        long startCpuTime = 0;
        long[] startThreadIds = bean.getAllThreadIds();
        for (long id : startThreadIds) {
            long time = bean.getThreadCpuTime(id);
            if (time != -1) startCpuTime += time;
        }

        // Wall-clock start
        long startTime = System.nanoTime();

        // Send requests
        ExecutorService executor = Executors.newFixedThreadPool(THREADS);
        List<Future<Boolean>> futures = new ArrayList<>();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(new URI(TARGET_URL))
                .timeout(Duration.ofSeconds(5))
                .GET()
                .build();

        for (int i = 0; i < NUM_REQUESTS; i++) {
            if (i % THREADS == 0) Thread.sleep(1); // Light pacing
            futures.add(executor.submit(() -> {
                try {
                    HttpResponse<Void> response = client.send(request, HttpResponse.BodyHandlers.discarding());
                    return response.statusCode() == 200;
                } catch (Exception e) {
                    return false;
                }
            }));
        }

        int successCount = 0;
        for (Future<Boolean> fut : futures) {
            if (fut.get()) successCount++;
        }
        executor.shutdown();

        // Wall-clock end
        long endTime = System.nanoTime();

        // CPU time after
        long endCpuTime = 0;
        long[] endThreadIds = bean.getAllThreadIds();
        for (long id : endThreadIds) {
            long time = bean.getThreadCpuTime(id);
            if (time != -1) endCpuTime += time;
        }

        // Memory after
        runtime.gc();
        long memAfter = runtime.totalMemory() - runtime.freeMemory();

        // Metrics
        double wallTime = (endTime - startTime) / 1e9;
        double cpuTime = (endCpuTime - startCpuTime) / 1e9;
        double maxMemory = Math.max(memBefore, memAfter) / (1024.0 * 1024.0);
        double throughput = wallTime > 0 ? NUM_REQUESTS / wallTime : 0;

        // Report
        System.out.println("\n--- Load Test Report ---");
        System.out.println("Total requests      : " + NUM_REQUESTS);
        System.out.println("Successful requests : " + successCount);
        System.out.println("Failed requests     : " + (NUM_REQUESTS - successCount));
        System.out.println("Wall-clock time     : " + wallTime + " seconds");
        System.out.println("Total CPU time      : " + cpuTime + " seconds");
        System.out.println("Max memory usage    : " + maxMemory + " MB");
        System.out.println("Throughput          : " + throughput + " requests/second");
    }
}
