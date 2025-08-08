import java.util.Random;
import java.lang.management.ManagementFactory;
import java.lang.management.ThreadMXBean;

public class MatrixMultiply {
    public static void main(String[] args) {
        int n = 1000;
        double[][] A = new double[n][n];
        double[][] B = new double[n][n];
        double[][] C = new double[n][n];

        Random rand = new Random(42);
        for (int i = 0; i < n; i++)
            for (int j = 0; j < n; j++) {
                A[i][j] = rand.nextDouble();
                B[i][j] = rand.nextDouble();
            }

        Runtime runtime = Runtime.getRuntime();
        runtime.gc();
        long memBefore = runtime.totalMemory() - runtime.freeMemory();

        ThreadMXBean bean = ManagementFactory.getThreadMXBean();
        long cpuStart = bean.getCurrentThreadCpuTime();
        long wallStart = System.nanoTime();

        for (int i = 0; i < n; i++)
            for (int j = 0; j < n; j++)
                for (int k = 0; k < n; k++)
                    C[i][j] += A[i][k] * B[k][j];

        long wallEnd = System.nanoTime();
        long cpuEnd = bean.getCurrentThreadCpuTime();
        runtime.gc();
        long memAfter = runtime.totalMemory() - runtime.freeMemory();

        System.out.printf("Execution time: %.3f s\n", (wallEnd - wallStart) / 1e9);
        System.out.printf("CPU time: %.3f s\n", (cpuEnd - cpuStart) / 1e9);
        System.out.printf("Memory used: %.2f MB\n", Math.max(memBefore, memAfter) / (1024.0 * 1024.0));
    }
}
