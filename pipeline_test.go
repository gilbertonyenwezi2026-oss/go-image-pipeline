package main

import (
	"os"
	"testing"
)

func TestValidateImagePathsWithValidFiles(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test-image-*.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	paths := []string{tempFile.Name()}

	err = validateImagePaths(paths)
	if err != nil {
		t.Fatalf("expected valid image path, got error: %v", err)
	}
}

func TestValidateImagePathsWithMissingFile(t *testing.T) {
	paths := []string{"missing-file.jpeg"}

	err := validateImagePaths(paths)
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestValidateImagePathsWithEmptyList(t *testing.T) {
	paths := []string{}

	err := validateImagePaths(paths)
	if err == nil {
		t.Fatal("expected error for empty image path list, got nil")
	}
}
