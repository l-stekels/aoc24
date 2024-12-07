package day1

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "3 4\n4 3\n2 5\n1 3\n3 9\n3 3"
	expected := SolutionInput{
		LeftColumn:  []int{3, 4, 2, 1, 3, 3},
		RightColumn: []int{4, 3, 5, 3, 9, 3},
	}

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	common.AssertEqualIntSlices(t, result.LeftColumn, expected.LeftColumn)
	common.AssertEqualIntSlices(t, result.RightColumn, expected.RightColumn)
}

func Test_SolvePart1(t *testing.T) {
	input := SolutionInput{
		LeftColumn:  []int{3, 4, 2, 1, 3, 3},
		RightColumn: []int{4, 3, 5, 3, 9, 3},
	}
	expected := 11

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("SolvePart1 failed: got %d, want %d", result, expected)
	}
}

func Test_SolvePart2(t *testing.T) {
	input := SolutionInput{
		LeftColumn:  []int{3, 4, 2, 1, 3, 3},
		RightColumn: []int{4, 3, 5, 3, 9, 3},
	}
	expected := 31

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("SolvePart1 failed: got %d, want %d", result, expected)
	}
}
