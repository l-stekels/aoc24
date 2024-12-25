package day8

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............"
	expected := createTestInput()

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	for x, row := range result.grid {
		for y, symbol := range row {
			if symbol != expected.grid[x][y] {
				t.Errorf("CreateSolutionInput grid not the same at x: %v, y: %v, got %v, want %v", x, y, symbol, expected.grid[x][y])
			}
		}
	}
	for freq, antennas := range result.antennas {
		for i, antenna := range antennas {
			if antenna != expected.antennas[freq][i] {
				t.Errorf("CreateSolutionInput antennas not the same at freq: %s, i: %v, got %v, want %v", freq, i, antenna, expected.antennas[Frequency(freq)][i])
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
			name:     "Example",
			input:    createTestInput(),
			expected: 14,
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
			input:    createTestInput(),
			expected: 34,
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

func createTestInput() SolutionInput {
	return SolutionInput{
		grid: [][]rune{
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', 'A', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		},
		antennas: map[Frequency][]Antenna{
			'0': {
				{pos: common.Point{X: 1, Y: 8}, freq: "0"},
				{pos: common.Point{X: 2, Y: 5}, freq: "0"},
				{pos: common.Point{X: 3, Y: 7}, freq: "0"},
				{pos: common.Point{X: 4, Y: 4}, freq: "0"},
			},
			'A': {
				{pos: common.Point{X: 5, Y: 6}, freq: "A"},
				{pos: common.Point{X: 8, Y: 8}, freq: "A"},
				{pos: common.Point{X: 9, Y: 9}, freq: "A"},
			},
		},
	}
}
