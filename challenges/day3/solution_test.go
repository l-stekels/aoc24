package day1

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := SolutionInput{
		Commands: []string{
			"mul(2,4)",
			"don't()",
			"mul(5,5)",
			"mul(11,8)",
			"do()",
			"mul(8,5)",
		},
	}

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	common.AssertEqualSlices[string](t, result.Commands, expected.Commands)
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
				Commands: []string{
					"mul(2,4)",
					"don't()",
					"mul(5,5)",
					"mul(11,8)",
					"do()",
					"mul(8,5)",
				},
			},
			expected: 161,
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
				Commands: []string{
					"mul(2,4)",
					"don't()",
					"mul(5,5)",
					"mul(11,8)",
					"do()",
					"mul(8,5)",
				},
			},
			expected: 48,
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
