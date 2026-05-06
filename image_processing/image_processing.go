package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func ReadImage(path string) (image.Image, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open input image %s: %w", path, err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not decode image %s: %w", path, err)
	}

	return img, nil
}

func WriteImage(path string, img image.Image) error {
	outputDir := filepath.Dir(path)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create output directory %s: %w", outputDir, err)
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create output image %s: %w", path, err)
	}
	defer outputFile.Close()

	if err := jpeg.Encode(outputFile, img, nil); err != nil {
		return fmt.Errorf("could not encode output image %s: %w", path, err)
	}

	return nil
}

func Grayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}

	return grayImg
}

func Resize(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	return resizedImg
}
