package day7

import (
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20"
	expected := SolutionInput{
		equations: []Equation{
			{terms: []int{10, 19}, testValue: 190},
			{terms: []int{81, 40, 27}, testValue: 3267},
			{terms: []int{17, 5}, testValue: 83},
			{terms: []int{15, 6}, testValue: 156},
			{terms: []int{6, 8, 6, 15}, testValue: 7290},
			{terms: []int{16, 10, 13}, testValue: 161011},
			{terms: []int{17, 8, 14}, testValue: 192},
			{terms: []int{9, 7, 18, 13}, testValue: 21037},
			{terms: []int{11, 6, 16, 20}, testValue: 292},
		},
	}

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	if len(result.equations) != len(expected.equations) {
		t.Fatalf("CreateSolutionInput failed: got %d, want %d", len(result.equations), len(expected.equations))
	}
	for i, equation := range result.equations {
		if !equation.Equals(expected.equations[i]) {
			t.Fatalf("CreateSolutionInput failed: got %v, want %v", equation, expected.equations[i])
		}
	}
}

func Test_SolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		input    SolutionInput
		expected uint64
	}{
		{
			name: "Example",
			input: SolutionInput{
				equations: []Equation{
					{terms: []int{10, 19}, testValue: 190},
					{terms: []int{81, 40, 27}, testValue: 3267},
					{terms: []int{17, 5}, testValue: 83},
					{terms: []int{15, 6}, testValue: 156},
					{terms: []int{6, 8, 6, 15}, testValue: 7290},
					{terms: []int{16, 10, 13}, testValue: 161011},
					{terms: []int{17, 8, 14}, testValue: 192},
					{terms: []int{9, 7, 18, 13}, testValue: 21037},
					{terms: []int{11, 6, 16, 20}, testValue: 292},
				},
			},
			expected: 3749,
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
		expected uint64
	}{
		{
			name: "Example",
			input: SolutionInput{
				equations: []Equation{
					{terms: []int{10, 19}, testValue: 190},
					{terms: []int{81, 40, 27}, testValue: 3267},
					{terms: []int{17, 5}, testValue: 83},
					{terms: []int{15, 6}, testValue: 156},
					{terms: []int{6, 8, 6, 15}, testValue: 7290},
					{terms: []int{16, 10, 13}, testValue: 161011},
					{terms: []int{17, 8, 14}, testValue: 192},
					{terms: []int{9, 7, 18, 13}, testValue: 21037},
					{terms: []int{11, 6, 16, 20}, testValue: 292},
				},
			},
			expected: 11387,
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
