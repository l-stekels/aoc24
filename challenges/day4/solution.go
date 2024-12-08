package day4

import (
	"fmt"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	stringRows := strings.Split(content, "\n")

	result := NewSolutionInput(len(stringRows[0]), len(stringRows))

	for i, row := range stringRows {
		for j, char := range row {
			result.Grid[i][j] = char
		}
	}

	return *result, nil
}

type SolutionInput struct {
	Grid [][]rune
}

func NewSolutionInput(nCols, nRows int) *SolutionInput {
	grid := make([][]rune, nRows)
	for i := range grid {
		grid[i] = make([]rune, nCols)
	}

	return &SolutionInput{Grid: grid}
}

func (s SolutionInput) Validate() error {
	var rowLength int
	for i, row := range s.Grid {
		if i == 0 {
			rowLength = len(row)
			continue
		}
		if rowLength != len(row) {
			return fmt.Errorf("row %d has different length: got %d, want %d", i, len(row), rowLength)
		}
		rowLength = len(row)
	}

	return nil
}

func SolvePart1(input SolutionInput) int {

	trie := NewTrie()
	trie.Insert("XMAS")

	visited := make([][]bool, len(input.Grid))
	for i := range visited {
		visited[i] = make([]bool, len(input.Grid[0]))
	}
	var foundWords []string
	for i := 0; i < len(input.Grid); i++ {
		for j := 0; j < len(input.Grid[0]); j++ {
			dfs(input.Grid, trie.root, i, j, visited, "", &foundWords, [2]int{0, 0})
		}
	}

	return len(foundWords)
}

func SolvePart2(input SolutionInput) int {
	result := 0

	return result
}
