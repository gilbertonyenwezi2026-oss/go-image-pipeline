package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"
	"time"

	imageprocessing "goroutines_pipeline/image_processing"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
	Err       error
}

func outputPath(inputPath string) string {
	fileName := filepath.Base(inputPath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "_processed.jpeg"
	return filepath.Join("images", "output", fileName)
}

func validateImagePaths(paths []string) error {
	if len(paths) == 0 {
		return fmt.Errorf("no input image paths were provided")
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("input image not found: %s", path)
		}
	}

	return nil
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)

	go func() {
		defer close(out)

		for _, path := range paths {
			img, err := imageprocessing.ReadImage(path)

			out <- Job{
				InputPath: path,
				Image:     img,
				OutPath:   outputPath(path),
				Err:       err,
			}
		}
	}()

	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		defer close(out)

		for job := range input {
			if job.Err == nil {
				job.Image = imageprocessing.Resize(job.Image)
			}

			out <- job
		}
	}()

	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		defer close(out)

		for job := range input {
			if job.Err == nil {
				job.Image = imageprocessing.Grayscale(job.Image)
			}

			out <- job
		}
	}()

	return out
}

func saveImage(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		defer close(out)

		for job := range input {
			if job.Err == nil {
				job.Err = imageprocessing.WriteImage(job.OutPath, job.Image)
			}

			out <- job
		}
	}()

	return out
}

func runConcurrent(paths []string) error {
	if err := validateImagePaths(paths); err != nil {
		return err
	}

	channel1 := loadImage(paths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	results := saveImage(channel3)

	for job := range results {
		if job.Err != nil {
			return job.Err
		}

		fmt.Printf("Success: %s -> %s\n", job.InputPath, job.OutPath)
	}

	return nil
}

func runSequential(paths []string) error {
	if err := validateImagePaths(paths); err != nil {
		return err
	}

	for _, path := range paths {
		img, err := imageprocessing.ReadImage(path)
		if err != nil {
			return err
		}

		img = imageprocessing.Resize(img)
		img = imageprocessing.Grayscale(img)

		outPath := outputPath(path)

		if err := imageprocessing.WriteImage(outPath, img); err != nil {
			return err
		}

		fmt.Printf("Success: %s -> %s\n", path, outPath)
	}

	return nil
}

func runComparisonLoop(imagePaths []string, runs int) error {
	var totalSequential time.Duration
	var totalConcurrent time.Duration

	fmt.Printf("Running sequential and concurrent pipelines %d times...\n\n", runs)

	for i := 1; i <= runs; i++ {
		seqStart := time.Now()

		if err := runSequential(imagePaths); err != nil {
			return fmt.Errorf("sequential run %d failed: %w", i, err)
		}

		seqElapsed := time.Since(seqStart)
		totalSequential += seqElapsed

		conStart := time.Now()

		if err := runConcurrent(imagePaths); err != nil {
			return fmt.Errorf("concurrent run %d failed: %w", i, err)
		}

		conElapsed := time.Since(conStart)
		totalConcurrent += conElapsed

		if i%100 == 0 {
			fmt.Printf("Completed %d runs...\n", i)
		}
	}

	avgSequential := totalSequential / time.Duration(runs)
	avgConcurrent := totalConcurrent / time.Duration(runs)

	fmt.Println("\nProcessing Time Comparison")
	fmt.Println("--------------------------")
	fmt.Printf("Total runs: %d\n", runs)
	fmt.Printf("Sequential total time: %v\n", totalSequential)
	fmt.Printf("Concurrent total time: %v\n", totalConcurrent)
	fmt.Printf("Sequential average time: %v\n", avgSequential)
	fmt.Printf("Concurrent average time: %v\n", avgConcurrent)

	if totalConcurrent < totalSequential {
		improvement := float64(totalSequential-totalConcurrent) / float64(totalSequential) * 100
		fmt.Printf("Concurrent was faster by %.2f%%\n", improvement)
	} else {
		difference := float64(totalConcurrent-totalSequential) / float64(totalSequential) * 100
		fmt.Printf("Sequential was faster by %.2f%%\n", difference)
	}

	return nil
}

func main() {
	mode := flag.String("mode", "concurrent", "processing mode: concurrent, sequential, compare, or compare1000")
	flag.Parse()

	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	start := time.Now()

	var err error

	switch *mode {
	case "concurrent":
		start := time.Now()
		err = runConcurrent(imagePaths)
		elapsed := time.Since(start)

		if err == nil {
			fmt.Printf("\nMode: concurrent\n")
			fmt.Printf("Processing time: %v\n", elapsed)
		}

	case "sequential":
		start := time.Now()
		err = runSequential(imagePaths)
		elapsed := time.Since(start)

		if err == nil {
			fmt.Printf("\nMode: sequential\n")
			fmt.Printf("Processing time: %v\n", elapsed)
		}

	case "compare1000":
		err = runComparisonLoop(imagePaths, 1000)

	default:
		err = fmt.Errorf("invalid mode %q: use concurrent, sequential, or compare1000", *mode)
	}

	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nMode: %s\n", *mode)
	fmt.Printf("Processing time: %v\n", elapsed)
}
