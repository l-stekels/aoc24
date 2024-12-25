package day6

import (
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#..."
	expected := SolutionInput{
		Map: [][]rune{
			{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
			{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		},
	}

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	for y, row := range expected.Map {
		for x, cell := range row {
			if result.Map[y][x] != cell {
				t.Errorf("CreateSolution input: map cell incorrect at %v, want %v got %v", NewPoint(x, y), cell, result.Map[y][x])
			}
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
			name: "Example",
			input: SolutionInput{
				Map: [][]rune{
					{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
					{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				},
			},
			expected: 41,
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
			name: "Example",
			input: SolutionInput{
				Map: [][]rune{
					{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
					{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				},
			},
			expected: 6,
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
