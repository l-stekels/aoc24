package day10

import (
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		grid: Grid{
			data: [][]int{},
		},
	}
	stringRows := strings.Split(content, "\n")
	for _, stringRow := range stringRows {
		strHeights := strings.Split(stringRow, "")
		var row []int
		for _, strHeight := range strHeights {
			height, err := strconv.Atoi(strHeight)
			if err != nil {
				panic(err)
			}
			row = append(row, height)
		}
		result.grid.data = append(result.grid.data, row)
	}
	result.grid.rows = len(result.grid.data)
	result.grid.cols = len(result.grid.data[0])

	return result, nil
}

type SolutionInput struct {
	grid Grid
}

type Grid struct {
	data [][]int
	rows int
	cols int
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
