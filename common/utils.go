package common

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReadInput[T SolutionInput](
	baseDir string,
	day ChallengeDay,
	input ChallengeInput,
	parser SolutionParser[T],
) (T, error) {
	filePath := filepath.Join(
		baseDir,
		"challenges",
		day.String(),
		input.String()+".txt",
	)

	file, err := os.Open(filePath)
	if err != nil {
		var zero T
		return zero, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	fileContent, err := io.ReadAll(file)
	// Normalize line endings
	normalizedContent := strings.ReplaceAll(string(fileContent), "\r\n", "\n")
	// Trim any trailing newline characters
	normalizedContent = strings.TrimRight(normalizedContent, "\n")

	return parser.CreateSolutionInput(normalizedContent)
}

// AbsDiff returns the absolute difference between two integers.
// For example, AbsDiff(3, 5) returns 2, and AbsDiff(5, 3) also returns 2.
func AbsDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
