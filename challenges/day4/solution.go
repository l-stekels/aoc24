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

	for i := 0; i < len(input.Grid); i++ {
		for j := 0; j < len(input.Grid[0]); j++ {
			// Technically, I only need to check neighbours of A characters
			if input.Grid[i][j] != 'A' {
				continue
			}
			// Continue if the 'A' character is on the edge of the grid
			if i-1 < 0 || i+1 >= len(input.Grid) || j-1 < 0 || j+1 >= len(input.Grid[0]) {
				continue
			}
			// Otherwise, build the 3x3 grid with 'A' in the middle
			// and check if the pattern is found
			xmasGrid := [][]rune{
				{input.Grid[i-1][j-1], input.Grid[i-1][j], input.Grid[i-1][j+1]},
				{input.Grid[i][j-1], input.Grid[i][j], input.Grid[i][j+1]},
				{input.Grid[i+1][j-1], input.Grid[i+1][j], input.Grid[i+1][j+1]},
			}
			if checkXMASPattern(xmasGrid) {
				result++
			}

		}
	}

	return result
}

func checkXMASPattern(grid [][]rune) bool {
	// I do not care about all the runes, 'A' will always be in the middle and I already know that,
	//so I only care about the corner runes
	topLeftCorner := grid[0][0]
	topRightCorner := grid[0][2]
	bottomLeftCorner := grid[2][0]
	bottomRightCorner := grid[2][2]
	// There are four possible patterns
	topLeftM := topLeftCorner == 'M'
	topLeftS := topLeftCorner == 'S'
	topRightS := topRightCorner == 'S'
	topRightM := topRightCorner == 'M'
	bottomRightS := bottomRightCorner == 'S'
	bottomRightM := bottomRightCorner == 'M'
	bottomLeftM := bottomLeftCorner == 'M'
	bottomLeftS := bottomLeftCorner == 'S'

	// M . S
	// . A .
	// M . S
	if topLeftM && topRightS && bottomLeftM && bottomRightS {
		return true
	}
	// S . M
	// . A .
	// S . M
	if topLeftS && topRightM && bottomLeftS && bottomRightM {
		return true
	}
	// S . S
	// . A .
	// M . M
	if topLeftS && topRightS && bottomLeftM && bottomRightM {
		return true
	}
	// M . M
	// . A .
	// S . S
	if topLeftM && topRightM && bottomLeftS && bottomRightS {
		return true
	}

	return false
}
