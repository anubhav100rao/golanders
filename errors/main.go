package main

import (
	"errors"
	"fmt"
	"os"
)

// The function openFile returns a custom error message if opening the file fails
func openFile(filename string) error {
	if _, err := os.ReadFile(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}

	return nil
}

func main() {
	err := openFile("new_file.txt")

	if err != nil {
		fmt.Printf("error running main.go: %s\n", err) // Print the wrapped error message

		unwrappedErr := errors.Unwrap(err)                // This line unwraps the error
		fmt.Printf("unwrapped error: %s\n", unwrappedErr) // Print the original error message
	}
}
