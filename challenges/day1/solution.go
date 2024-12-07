package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	// Normalize spaces
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.Join(strings.Fields(line), " ")
	}

	var leftColumn, rightColumn []int
	for _, line := range lines {
		values := strings.Split(line, " ")
		if len(values) != 2 {
			return SolutionInput{}, fmt.Errorf("invalid line format: %s", line)
		}
		value1, err := strconv.Atoi(values[0])
		if err != nil {
			return SolutionInput{}, fmt.Errorf("invalid value in line: %s", line)
		}
		value2, err := strconv.Atoi(values[1])
		if err != nil {
			return SolutionInput{}, fmt.Errorf("invalid value in line: %s", line)
		}
		leftColumn = append(leftColumn, value1)
		rightColumn = append(rightColumn, value2)
	}

	return SolutionInput{
		LeftColumn:  leftColumn,
		RightColumn: rightColumn,
	}, nil
}

type SolutionInput struct {
	LeftColumn  []int
	RightColumn []int
}

func (s SolutionInput) Validate() error {
	if len(s.LeftColumn) != len(s.RightColumn) {
		return fmt.Errorf("column length mismatch: %d != %d", len(s.LeftColumn), len(s.RightColumn))
	}
	return nil
}

func SolvePart1(input SolutionInput) int {
	if err := input.Validate(); err != nil {
		panic(err)
	}
	sort.Ints(input.LeftColumn)
	sort.Ints(input.RightColumn)
	var result int
	for i := 0; i < len(input.LeftColumn); i++ {
		var distance int
		distance = absDiff(input.LeftColumn[i], input.RightColumn[i])
		result += distance
	}

	return result
}

func SolvePart2(input SolutionInput) int {
	if err := input.Validate(); err != nil {
		panic(err)
	}
	similarityScore := 0

	for i := 0; i < len(input.LeftColumn); i++ {
		left := input.LeftColumn[i]
		occurrences := countOccurrences(left, input.RightColumn)
		similarityScore += left * occurrences
	}

	return similarityScore
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func countOccurrences(needle int, haystack []int) int {
	count := 0
	for _, value := range haystack {
		if value == needle {
			count++
		}
	}

	return count
}
