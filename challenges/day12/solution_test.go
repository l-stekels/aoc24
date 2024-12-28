package day12

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE"
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
			t.Errorf("CreateSolutionInput grid incorrect at %v, want %v got %v", position, string(expected.grid.Get(position)), string(value))
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
			expected: 0,
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

func createSolutionInput() SolutionInput {
	return SolutionInput{
		grid: common.NewGrid[rune]([][]rune{
			{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'F', 'F'},
			{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'C', 'F'},
			{'V', 'V', 'R', 'R', 'R', 'C', 'C', 'F', 'F', 'F'},
			{'V', 'V', 'R', 'C', 'C', 'C', 'J', 'F', 'F', 'F'},
			{'V', 'V', 'V', 'V', 'C', 'J', 'J', 'C', 'F', 'E'},
			{'V', 'V', 'I', 'V', 'C', 'C', 'J', 'J', 'E', 'E'},
			{'V', 'V', 'I', 'I', 'I', 'C', 'J', 'J', 'E', 'E'},
			{'M', 'I', 'I', 'I', 'I', 'I', 'J', 'J', 'E', 'E'},
			{'M', 'I', 'I', 'I', 'S', 'I', 'J', 'E', 'E', 'E'},
			{'M', 'M', 'M', 'I', 'S', 'S', 'J', 'E', 'E', 'E'},
		}),
	}
}
