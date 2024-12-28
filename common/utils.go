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
	parser SolutionParser[T],
) (T, error) {
	filePath := filepath.Join(
		baseDir,
		"challenges",
		day.String(),
		"input.txt",
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
// For example, both AbsDiff(3, 5) and AbsDiff(5, 3) return 2.
func AbsDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// Gcd returns the greatest common divisor of x and y.
func Gcd(x int, y int) int {
	a := Abs(x)
	b := Abs(y)
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
