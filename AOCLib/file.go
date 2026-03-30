package aoclib

import (
	"fmt"
	"os"
)

// ReadFile read in the file and return the contents, prints an error message and returns empty if failed to read
func ReadFile(inFile string) string {
	content, err := os.ReadFile(inFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read in %s, %v\n", inFile, err)
		return ""
	}

	return string(content)
}
