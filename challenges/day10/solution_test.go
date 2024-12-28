package day10

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732"
	expected := createSolutionInput()

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	for result.grid.HasNext() {
		value, position := result.grid.Next()
		if value != expected.grid.Get(position) {
			t.Errorf("CreateSolutionInput grid incorrect at %v, want %v got %v", position, expected.grid.Get(position), value)
		}
	}
}

func Test_SolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		input    SolutionInput
		expected int
	}{
		{
			name:     "Example",
			input:    createSolutionInput(),
			expected: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SolvePart1(tt.input)
			if result != tt.expected {
				t.Errorf("SolvePart1 failed: got %d, want %d", result, tt.expected)
			}
		})
	}
}

func Test_SolvePart2(t *testing.T) {
	tests := []struct {
		name     string
		input    SolutionInput
		expected int
	}{
		{
			name:     "Example",
			input:    createSolutionInput(),
			expected: 81,
		},
		{
			name:     "Example 2",
			input:    createSolutionInput(2),
			expected: 3,
		},
		{
			name:     "Example 3",
			input:    createSolutionInput(3),
			expected: 13,
		},
		{
			name:     "Example 4",
			input:    createSolutionInput(4),
			expected: 227,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SolvePart2(tt.input)
			if result != tt.expected {
				t.Errorf("SolvePart1 failed: got %d, want %d", result, tt.expected)
			}
		})
	}
}

func createSolutionInput(args ...int) SolutionInput {
	if len(args) == 0 {
		return SolutionInput{
			grid: common.NewGrid[int]([][]int{
				{8, 9, 0, 1, 0, 1, 2, 3},
				{7, 8, 1, 2, 1, 8, 7, 4},
				{8, 7, 4, 3, 0, 9, 6, 5},
				{9, 6, 5, 4, 9, 8, 7, 4},
				{4, 5, 6, 7, 8, 9, 0, 3},
				{3, 2, 0, 1, 9, 0, 1, 2},
				{0, 1, 3, 2, 9, 8, 0, 1},
				{1, 0, 4, 5, 6, 7, 3, 2},
			}),
		}
	}
	var number int
	if len(args) == 1 {
		number = args[0]
	}
	switch number {
	case 2:
		return SolutionInput{
			grid: common.NewGrid[int]([][]int{
				{-1, -1, -1, -1, -1, 0, -1},
				{-1, -1, 4, 3, 2, 1, -1},
				{-1, -1, 5, -1, -1, 2, -1},
				{-1, -1, 6, 5, 4, 3, -1},
				{-1, -1, 7, -1, -1, 4, -1},
				{-1, -1, 8, 7, 6, 5, -1},
				{-1, -1, 9, -1, -1, -1, -1},
			}),
		}
	case 3:
		return SolutionInput{
			grid: common.NewGrid[int]([][]int{
				{-1, -1, 9, 0, -1, -1, 9},
				{-1, -1, -1, 1, -1, 9, 8},
				{-1, -1, -1, 2, -1, -1, 7},
				{6, 5, 4, 3, 4, 5, 6},
				{7, 6, 5, -1, 9, 8, 7},
				{8, 7, 6, -1, -1, -1, -1},
				{9, 8, 7, -1, -1, -1, -1},
			}),
		}
	case 4:
		return SolutionInput{
			grid: common.NewGrid[int]([][]int{
				{0, 1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5, 6},
				{2, 3, 4, 5, 6, 7},
				{3, 4, 5, 6, 7, 8},
				{4, -1, 6, 7, 8, 9},
				{5, 6, 7, 8, 9, -1},
			}),
		}
	default:
		panic("Invalid number")
	}

}
