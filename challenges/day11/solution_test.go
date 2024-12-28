package day11

import (
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "125 17"
	expected := createSolutionInput()

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}
	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	if len(result.stones) != len(expected.stones) {
		t.Errorf("CreateSolutionInput failed: got %d, want %d", len(result.stones), len(expected.stones))
	}
	for i, stone := range result.stones {
		if stone != expected.stones[i] {
			t.Errorf("CreateSolutionInput failed: got %d, want %d", stone, expected.stones[i])
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
			expected: 55312,
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
			expected: 65601038650482,
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
		stones: []Stone{
			{value: uint64(125)},
			{value: uint64(17)},
		},
	}
}
