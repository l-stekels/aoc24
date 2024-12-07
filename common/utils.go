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
