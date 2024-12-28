package day12

import (
	"advent2024/common"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		grid: common.NewGrid[rune]([][]rune{}),
	}
	stringRows := strings.Split(content, "\n")
	for _, row := range stringRows {
		result.grid.AddRow([]rune(row))
	}

	return result, nil
}

type SolutionInput struct {
	grid common.Grid[rune]
}

func (s SolutionInput) Validate() error {
	return nil
}

func SolvePart1(input SolutionInput) int {
	result := 0

	return result
}

func SolvePart2(input SolutionInput) int {
	result := 0

	return result
}
