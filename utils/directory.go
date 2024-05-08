package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFileWithDirectories(filePath, content string) (*os.File, error) {
	// Ensure the directory structure exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating directories: %v", err)
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error creating file: %v", err)
	}

	// Write content to the file
	if _, err := file.WriteString(content); err != nil {
		file.Close()
		return nil, fmt.Errorf("error writing to file: %v", err)
	}

	return file, nil
}
