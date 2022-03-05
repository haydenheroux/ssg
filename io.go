package main

import (
	"bufio"
	"os"
)

func Read(filePath string) ([]string, error) {
	// Open the specified file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// Ensure the file will close after an error occurs or all lines are read
	defer file.Close()

	// Initialize a slice to store the lines of the file
	lines := make([]string, 0)

	// Initialize a scanner to read through the file
	scanner := bufio.NewScanner(file)

	// Consume each line of the file
	for scanner.Scan() {
		// Append the consumed line to the lines slice
		lines = append(lines, scanner.Text())
	}

	// Provide the error to the caller if an error occurred
	if err := scanner.Err(); err != nil {
		return lines, err
	}

	// Provide the caller the lines of the file
	return lines, nil
}
