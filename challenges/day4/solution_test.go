package day4

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"
	expected := SolutionInput{
		Grid: [][]rune{
			{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
			{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
			{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
			{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
			{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
			{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
			{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
			{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
			{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
			{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
		},
	}

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	common.AssertEqual2DSlices[rune](t, result.Grid, expected.Grid)
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
				Grid: [][]rune{
					{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
					{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
					{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
					{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
					{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
					{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
					{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
					{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
					{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
					{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
				},
			},
			expected: 18,
		},
		{
			name: "Example 2",
			input: SolutionInput{
				Grid: [][]rune{
					{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
					{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
					{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
					{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
					{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
					{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
					{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
					{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
					{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
					{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
				},
			},
			expected: 18,
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
			input:    SolutionInput{},
			expected: 0,
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
