package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"
)

// createBenchmarkImage creates a temporary JPEG image for benchmarking.
// This avoids depending on external image files during benchmark tests.
func createBenchmarkImage(b *testing.B, dir string, fileName string) string {
	b.Helper()

	img := image.NewRGBA(image.Rect(0, 0, 1200, 800))

	// Fill the image with synthetic pixel data.
	for y := 0; y < 800; y++ {
		for x := 0; x < 1200; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(x % 255),
				G: uint8(y % 255),
				B: uint8((x + y) % 255),
				A: 255,
			})
		}
	}

	filePath := filepath.Join(dir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		b.Fatalf("failed to create benchmark image: %v", err)
	}
	defer file.Close()

	if err := jpeg.Encode(file, img, nil); err != nil {
		b.Fatalf("failed to encode benchmark image: %v", err)
	}

	return filePath
}

// createBenchmarkImages creates multiple temporary images for the benchmark.
func createBenchmarkImages(b *testing.B) ([]string, string) {
	b.Helper()

	tempDir := b.TempDir()
	outputDir := filepath.Join(tempDir, "output")

	imagePaths := []string{
		createBenchmarkImage(b, tempDir, "bench_image1.jpeg"),
		createBenchmarkImage(b, tempDir, "bench_image2.jpeg"),
		createBenchmarkImage(b, tempDir, "bench_image3.jpeg"),
		createBenchmarkImage(b, tempDir, "bench_image4.jpeg"),
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		b.Fatalf("failed to create benchmark output directory: %v", err)
	}

	return imagePaths, outputDir
}

func BenchmarkSequentialPipeline(b *testing.B) {
	imagePaths, _ := createBenchmarkImages(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := runSequential(imagePaths); err != nil {
			b.Fatalf("sequential pipeline failed: %v", err)
		}
	}
}

func BenchmarkConcurrentPipeline(b *testing.B) {
	imagePaths, _ := createBenchmarkImages(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := runConcurrent(imagePaths); err != nil {
			b.Fatalf("concurrent pipeline failed: %v", err)
		}
	}
}
