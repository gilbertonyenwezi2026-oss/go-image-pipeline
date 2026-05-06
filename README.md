# Go Concurrent Image Processing Pipeline

## Overview

This repository contains a Go-based image processing pipeline that demonstrates how concurrency can improve throughput for data engineering-style workloads. The project is based on Amrit Singh’s Codeheim goroutines pipeline example and extends the original implementation by adding error checking, custom input images, sequential and concurrent execution modes, unit tests, benchmark tests, executable build instructions, and documentation.

The management problem for this assignment is centered on reducing processing time in data engineering tasks. Managers at a technology startup believe that Go concurrency may reduce data pipeline throughput time. This project tests that idea by comparing an image processing pipeline that runs sequentially against one that uses goroutines and channels.

The pipeline performs the following operations:

1. Loads image files from disk.
2. Resizes each image.
3. Converts each image to grayscale.
4. Saves the processed images to an output folder.
5. Measures and compares runtime for sequential and concurrent execution.

---

## Management Problem

Managers of a technology startup are interested in reducing processing times associated with data engineering tasks. They believe Go concurrency can improve throughput by allowing different stages of a pipeline to execute at the same time.

This assignment replicates and extends an existing Go image processing pipeline to evaluate that idea. The image processing workload acts as a proxy for a broader data engineering pipeline where records, files, or data batches move through multiple transformation stages.

The key question is:

> Can a goroutine-based pipeline reduce total processing time compared with a sequential pipeline?

---

## Original Source

This project was adapted from Amrit Singh’s Codeheim example:

- GitHub Repository: `https://github.com/code-heim/go_21_goroutines_pipeline`
- Video Tutorial: `https://www.youtube.com/watch?v=8Rn8yOQH62k`
- Codeheim Website: `https://www.codeheim.io/`

The original project demonstrates the pipeline pattern in Go using goroutines and channels. This version modifies and expands the original code to meet the assignment requirements.

---

## Assignment Requirements Covered

This repository satisfies the following assignment requirements:

| Requirement | Status |
|---|---|
| Clone original GitHub repository | Completed |
| Build and run program in original form | Completed |
| Add error checking for image input and output | Completed |
| Replace original four input images | Completed |
| Add unit tests | Completed |
| Add benchmark methods | Completed |
| Run pipeline with goroutines | Completed |
| Run pipeline without goroutines | Completed |
| Compare processing times | Completed |
| Prepare complete README.md | Completed |
| Provide executable build instructions | Completed |
| Include GenAI Tools section | Completed |
| Optional aspect-ratio discussion | Included |

---

## Features Added

The original pipeline was extended with the following improvements:

### 1. Sequential and Concurrent Modes

The program can now run in two modes:

```bash
go run . -mode=sequential
go run . -mode=concurrent
```

The sequential mode processes one image at a time from start to finish.

The concurrent mode uses goroutines and channels to process images through pipeline stages.

---

### 2. Error Checking

The revised program checks for common input and output problems, including:

- Missing image files.
- Empty image path lists.
- Invalid image paths.
- Image open failures.
- Image decoding failures.
- Output directory creation errors.
- Image writing and encoding errors.
- Invalid command-line mode values.

Instead of crashing unexpectedly, the program now returns meaningful error messages.

---

### 3. Custom Input Images

The original images were replaced with selected custom image files for this assignment. The program expects four input images in the `images` folder.

Example:

```text
images/
├── image1.jpeg
├── image2.jpeg
├── image3.jpeg
├── image4.jpeg
└── output/
```

---

### 4. Unit Tests

Unit tests were added to verify important program behavior, including:

- Valid input path handling.
- Missing input file detection.
- Empty input list detection.
- Basic validation logic.

Tests can be run with:

```bash
go test ./...
```

---

### 5. Benchmark Tests

Benchmark methods were added to compare sequential and concurrent throughput.

Benchmarks can be run with:

```bash
go test -bench=. -benchmem
```

The `-benchmem` flag reports memory usage and allocation counts in addition to runtime.

---

### 6. Executable Build Support

The program can be compiled into a Windows executable:

```bash
go build -o image_pipeline.exe .
```

The executable can then be run directly.

In Git Bash:

```bash
./image_pipeline.exe -mode=sequential
./image_pipeline.exe -mode=concurrent
```

In PowerShell:

```powershell
.\image_pipeline.exe -mode=sequential
.\image_pipeline.exe -mode=concurrent
```

---

## Repository Structure

```text
go-image-pipeline/
│
├── go.mod
├── go.sum
├── main.go
├── pipeline_test.go
├── benchmark_test.go
├── README.md
│
├── image_processing/
│   └── image_processing.go
│
├── images/
│   ├── image1.jpeg
│   ├── image2.jpeg
│   ├── image3.jpeg
│   ├── image4.jpeg
│   └── output/
│
└── image_pipeline.exe
```

---

## Program Design

The application uses a simple but effective pipeline design.

The major processing stages are:

1. `loadImage`
2. `resize`
3. `convertToGrayscale`
4. `saveImage`

In concurrent mode, each stage runs as part of a goroutine-based pipeline. Data moves between stages through Go channels.

In sequential mode, each image is processed completely before the next image begins.

---

## Sequential Pipeline

The sequential pipeline follows this flow:

```text
Image 1 → Load → Resize → Grayscale → Save
Image 2 → Load → Resize → Grayscale → Save
Image 3 → Load → Resize → Grayscale → Save
Image 4 → Load → Resize → Grayscale → Save
```

This approach is simple and easy to understand, but it does not take advantage of concurrent execution.

---

## Concurrent Pipeline

The concurrent pipeline follows this flow:

```text
Input Images
    ↓
Load Image Stage
    ↓
Resize Stage
    ↓
Grayscale Stage
    ↓
Save Image Stage
    ↓
Output Images
```

Each stage communicates through channels. This allows multiple images to be in different stages of processing at the same time.

For example, while one image is being resized, another may be loading, and another may be saving.

---

## Requirements

This project requires:

- Go 1.20 or higher
- Git
- A terminal such as Git Bash, PowerShell, Command Prompt, or VS Code terminal

Check your Go version with:

```bash
go version
```

---

## Installation and Setup

### 1. Clone the Repository

```bash
git clone https://github.com/gilbertonyenwezi2026-oss/go-image-pipeline.git
cd go-image-pipeline
```

---

### 2. Install Dependencies

```bash
go mod tidy
```

---

### 3. Confirm Project Files

```bash
ls
```

You should see files such as:

```text
main.go
go.mod
README.md
pipeline_test.go
benchmark_test.go
image_processing
images
```

---

## Running the Program

### Run Concurrent Mode

```bash
go run . -mode=concurrent
```

Example output:

```text
Success: images/image1.jpeg -> images/output/image1_processed.jpeg
Success: images/image2.jpeg -> images/output/image2_processed.jpeg
Success: images/image3.jpeg -> images/output/image3_processed.jpeg
Success: images/image4.jpeg -> images/output/image4_processed.jpeg

Mode: concurrent
Processing time: 25.4102ms
```

---

### Run Sequential Mode

```bash
go run . -mode=sequential
```

Example output:

```text
Success: images/image1.jpeg -> images/output/image1_processed.jpeg
Success: images/image2.jpeg -> images/output/image2_processed.jpeg
Success: images/image3.jpeg -> images/output/image3_processed.jpeg
Success: images/image4.jpeg -> images/output/image4_processed.jpeg

Mode: sequential
Processing time: 43.8127ms
```

Actual processing times will vary depending on the computer, operating system, image size, and background processes.

---

## Command-Line Options

| Flag | Description | Example |
|---|---|---|
| `-mode=sequential` | Runs the pipeline without goroutines | `go run . -mode=sequential` |
| `-mode=concurrent` | Runs the goroutine-based pipeline | `go run . -mode=concurrent` |

If an invalid mode is entered, the program displays an error message.

Example:

```bash
go run . -mode=fast
```

Expected result:

```text
Error: invalid mode "fast": use concurrent or sequential
```

---

## Testing

Run all unit tests with:

```bash
go test ./...
```

Expected successful output:

```text
ok      goroutines_pipeline    0.123s
?       goroutines_pipeline/image_processing    [no test files]
```

The unit tests validate important components of the application, especially input validation logic.

---

## Benchmarking

Run benchmark tests with:

```bash
go test -bench=. -benchmem
```

This command reports:

- Nanoseconds per operation.
- Memory allocated per operation.
- Number of allocations per operation.

---

## Benchmark Results

Replace the example values below with your actual benchmark results.


| Mode | Command | Time/op | Memory/op | Allocations/op |
|---|---|---:|---:|---:|
| Sequential | `go test -bench=BenchmarkSequentialPipeline -benchmem` | 55700123 ns/op | 32638555 B/op | 1000367 allocs/op |
| Concurrent | `go test -bench=BenchmarkConcurrentPipeline -benchmem` | 53562200 ns/op | 32639685 B/op | 1000369 allocs/op |
PASS
ok      goroutines_pipeline   16.641s


---

## Processing Time Comparison

Replace this section with your actual runtime output.

| Run Mode | Processing Time |
|---|---:|
| Sequential | 115.9204ms |
| Concurrent | 69.8563ms |


---

## Interpretation of Results

The concurrent pipeline is expected to perform better when the workload is large enough to benefit from overlapping work across multiple goroutines. In this project, concurrency allows multiple pipeline stages to operate at the same time.

However, concurrency is not always automatically faster. For small images or small workloads, the overhead of creating goroutines and passing data through channels may reduce or eliminate performance gains. For larger image sets or more computationally expensive transformations, the concurrent pipeline is more likely to show stronger throughput improvement.

The benchmark results provide a more reliable comparison than a single program run because benchmarks repeat the operation multiple times and report average performance.

---

## Error Handling

The project includes improved error handling for both input and output operations.

Examples of errors the program now handles:

```text
Input image not found
Invalid image path
Image decode failure
Output folder creation failure
Image write failure
Invalid processing mode
```

This improves usability because users receive clear explanations instead of unexpected crashes.

Example error:

```text
Error: input image not found: images/image1.jpeg
```

---

## Image Processing Functions

The `image_processing` package includes helper functions for:

- Reading images.
- Writing images.
- Resizing images.
- Converting images to grayscale.

The main application coordinates these helper functions through either sequential control flow or concurrent pipeline stages.

---

## Aspect Ratio Note

The original helper function resizes images to a fixed `500x500` size. This can cause distortion when the input image is not square.

A possible improvement is to preserve the original aspect ratio by detecting the image dimensions before resizing.

Example logic:

```go
bounds := img.Bounds()
width := bounds.Dx()
height := bounds.Dy()
```

Then the program can calculate a proportional width and height instead of forcing the image into a square shape.

This assignment version recognizes the distortion issue and documents aspect-ratio preservation as a recommended improvement.

---

## Building the Executable

### Windows

```bash
go build -o image_pipeline.exe .
```

Run in Git Bash:

```bash
./image_pipeline.exe -mode=sequential
./image_pipeline.exe -mode=concurrent
```

Run in PowerShell:

```powershell
.\image_pipeline.exe -mode=sequential
.\image_pipeline.exe -mode=concurrent
```

---

### MacOS or Linux

```bash
go build -o image_pipeline .
```

Run:

```bash
./image_pipeline -mode=sequential
./image_pipeline -mode=concurrent
```

---

## Final Verification Commands

Before submitting the assignment, run the following commands:

```bash
go fmt ./...
go mod tidy
go test ./...
go test -bench=. -benchmem
go build -o image_pipeline.exe .
./image_pipeline.exe -mode=sequential
./image_pipeline.exe -mode=concurrent
git status
```

If using PowerShell, run:

```powershell
go fmt ./...
go mod tidy
go test ./...
go test -bench=. -benchmem
go build -o image_pipeline.exe .
.\image_pipeline.exe -mode=sequential
.\image_pipeline.exe -mode=concurrent
git status
```

---

## GitHub Submission

The cloneable GitHub repository URL is:

```text
https://github.com/gilbertonyenwezi2026-oss/go-image-pipeline.git
```

This is the URL to submit in the assignment comments form.

---

## GenAI Tools

Generative AI tools were used as programming support during this assignment.

I used ChatGPT to help:

- Add error checking logic.
- Troubleshoot Go compiler and shell command errors.

ChatGPT was used as a coding assistant and learning tool. All code was reviewed, modified, tested, and validated before being included in the final repository. The final program was built and run locally to confirm that the sequential and concurrent modes worked correctly.

Conversation logs or plain text notes documenting GenAI usage may be saved in the repository as:

```text
docs/genai-tools-log.txt
```

---

## Lessons Learned

This assignment demonstrated several important software engineering lessons:

1. Go concurrency is useful for pipeline-style workloads.
2. Goroutines and channels allow clean separation of processing stages.
3. Concurrent code should still be benchmarked because concurrency is not always faster.
4. Error handling is essential for building user-friendly applications.
5. Unit tests improve confidence that individual components work correctly.
6. Benchmarks help compare performance in a repeatable way.
7. Documentation is a major part of software engineering, not an afterthought.

---

## References

Go Authors. “Package testing.” *Go Packages*. Accessed May 5, 2026. https://pkg.go.dev/testing.

Go Authors. “Package image.” *Go Packages*. Accessed May 5, 2026. https://pkg.go.dev/image.

Go Authors. “Package image/jpeg.” *Go Packages*. Accessed May 5, 2026. https://pkg.go.dev/image/jpeg.

McConnell, Steve. *Code Complete: A Practical Handbook of Software Construction*. 2nd ed. Redmond, WA: Microsoft Press, 2004.

Singh, Amrit. “Go Goroutines Pipeline.” Codeheim. Accessed May 5, 2026. https://github.com/code-heim/go_21_goroutines_pipeline.

Singh, Amrit. “Concurrency in Go: Pipeline Pattern.” YouTube video. Codeheim. Accessed May 5, 2026. https://www.youtube.com/watch?v=8Rn8yOQH62k.

---

## Author

Gilbert Onyenwezi 
MSDS 431  
Northwestern University  
Go Assisted Programming Assignment
