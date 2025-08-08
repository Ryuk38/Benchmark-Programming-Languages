import pandas as pd
import time
import tracemalloc

# Start profiling
start_wall_time = time.time()
start_cpu_time = time.process_time()
tracemalloc.start()

# Step 1: Read the CSV file
input_file = r"C:\Users\c380b\OneDrive\Desktop\CDS\IO\IMDB Dataset.csv"
df = pd.read_csv(input_file)

# Step 2: Count the sentiment values
positive_count = (df['sentiment'].str.lower().str.strip() == 'positive').sum()
negative_count = (df['sentiment'].str.lower().str.strip() == 'negative').sum()

# Step 3: Prepare output data
output_text = f"Total positive reviews: {positive_count}\nTotal negative reviews: {negative_count}"

# Step 4: Write the result to a new file
output_file = r"C:\Users\c380b\OneDrive\Desktop\CDS\IO\sentiment_pyhton.txt"
with open(output_file, "w") as f:
    f.write(output_text)

# Stop profiling
end_wall_time = time.time()
end_cpu_time = time.process_time()
current, peak = tracemalloc.get_traced_memory()
tracemalloc.stop()

# Print profiling results
print("Sentiment counts written to", output_file)
print(f"\n--- Profiling Report ---")
print(f"Wall-clock time   : {end_wall_time - start_wall_time:.4f} seconds")
print(f"CPU time          : {end_cpu_time - start_cpu_time:.4f} seconds")
print(f"Peak memory usage : {peak / 1024 / 1024:.2f} MB")   