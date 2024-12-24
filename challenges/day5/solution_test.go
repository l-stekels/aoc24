package day5

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47"
	expected := SolutionInput{
		Update: []Update{
			{
				Pages:      []int{75, 47, 61, 53, 29},
				MiddlePage: 61,
			},
			{
				Pages:      []int{97, 61, 53, 29, 13},
				MiddlePage: 53,
			},
			{
				Pages:      []int{75, 29, 13},
				MiddlePage: 29,
			},
			{
				Pages:      []int{75, 97, 47, 61, 53},
				MiddlePage: 47,
			},
			{
				Pages:      []int{61, 13, 29},
				MiddlePage: 13,
			},
			{
				Pages:      []int{97, 13, 75, 29, 47},
				MiddlePage: 75,
			},
		},
		Rules: map[int][]int{
			47: {53, 13, 61, 29},
			53: {29, 13},
			97: {13, 61, 47, 29, 53, 75},
			13: {},
			61: {13, 53, 29},
			75: {29, 53, 47, 61, 13},
			29: {13},
		},
	}

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	for key, value := range expected.Rules {
		common.AssertEqualSlices[int](t, result.Rules[key], value)
	}
	for key, value := range expected.Update {
		common.AssertEqualSlices[int](t, result.Update[key].Pages, value.Pages)
		if result.Update[key].MiddlePage != value.MiddlePage {
			t.Errorf("CreateSolution input: update middle page incorrect for %v, want %v got %v", value, value.MiddlePage, result.Update[key].MiddlePage)
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
				Update: []Update{
					{
						Pages:      []int{75, 47, 61, 53, 29},
						MiddlePage: 61,
					},
					{
						Pages:      []int{97, 61, 53, 29, 13},
						MiddlePage: 53,
					},
					{
						Pages:      []int{75, 29, 13},
						MiddlePage: 29,
					},
					{
						Pages:      []int{75, 97, 47, 61, 53},
						MiddlePage: 47,
					},
					{
						Pages:      []int{61, 13, 29},
						MiddlePage: 13,
					},
					{
						Pages:      []int{97, 13, 75, 29, 47},
						MiddlePage: 75,
					},
				},
				Rules: map[int][]int{
					47: {53, 13, 61, 29},
					53: {29, 13},
					97: {13, 61, 47, 29, 53, 75},
					13: {},
					61: {13, 53, 29},
					75: {29, 53, 47, 61, 13},
					29: {13},
				},
			},
			expected: 143,
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
				Update: []Update{
					{
						Pages:      []int{75, 47, 61, 53, 29},
						MiddlePage: 61,
					},
					{
						Pages:      []int{97, 61, 53, 29, 13},
						MiddlePage: 53,
					},
					{
						Pages:      []int{75, 29, 13},
						MiddlePage: 29,
					},
					{
						Pages:      []int{75, 97, 47, 61, 53},
						MiddlePage: 47,
					},
					{
						Pages:      []int{61, 13, 29},
						MiddlePage: 13,
					},
					{
						Pages:      []int{97, 13, 75, 29, 47},
						MiddlePage: 75,
					},
				},
				Rules: map[int][]int{
					47: {53, 13, 61, 29},
					53: {29, 13},
					97: {13, 61, 47, 29, 53, 75},
					13: {},
					61: {13, 53, 29},
					75: {29, 53, 47, 61, 13},
					29: {13},
				},
			},
			expected: 123,
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
