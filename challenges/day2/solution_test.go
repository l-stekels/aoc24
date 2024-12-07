package day1

import (
	"advent2024/common"
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"
	expected := SolutionInput{
		Reports: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		},
	}

	result, err := parser.CreateSolutionInput(input)
	if err != nil {
		t.Fatalf("CreateSolutionInput failed: %v", err)
	}

	if err := result.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	common.AssertEqual2DIntSlices(t, result.Reports, expected.Reports)
}

func Test_SolvePart1(t *testing.T) {
	tests := []struct {
		input    SolutionInput
		expected int
	}{
		{
			input: SolutionInput{
				Reports: [][]int{
					{7, 6, 4, 2, 1},     // Safe
					{1, 2, 7, 8, 9},     // Unsafe
					{9, 7, 6, 2, 1},     // Unsafe
					{1, 3, 2, 4, 5},     // Unsafe
					{8, 6, 4, 4, 1},     // Unsafe
					{1, 3, 6, 7, 9},     // Safe
					{6, 11, 12, 10, 13}, // Unsafe
				},
			},
			expected: 2,
		},
		{
			input: SolutionInput{
				Reports: [][]int{
					{1, 2, 3, 4, 5}, // Safe
					{5, 4, 3, 2, 1}, // Safe
				},
			},
			expected: 2,
		},
		{
			input: SolutionInput{
				Reports: [][]int{
					{1, 1, 1, 1, 1}, // Unsafe
					{2, 2, 2, 2, 2}, // Unsafe
				},
			},
			expected: 0,
		},
		{
			input: SolutionInput{
				Reports: [][]int{
					{1, 1, 1, 1, 1}, // Unsafe
					{2, 2, 2, 2, 2}, // Unsafe
				},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
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
				Reports: [][]int{
					{7, 6, 4, 2, 1}, // Safe
					{1, 2, 7, 8, 9}, // Unsafe
					{9, 7, 6, 2, 1}, // Unsafe
					{1, 3, 2, 4, 5}, // Safe
					{8, 6, 4, 4, 1}, // Safe
					{1, 3, 6, 7, 9}, // Safe
				},
			},
			expected: 4,
		},
		{
			name: "Both safe",
			input: SolutionInput{
				Reports: [][]int{
					{1, 2, 3, 4, 5}, // Safe
					{5, 4, 3, 2, 1}, // Safe
				},
			},
			expected: 2,
		},
		{
			name: "Both unsafe",
			input: SolutionInput{
				Reports: [][]int{
					{1, 1, 1, 1, 1}, // Unsafe
					{2, 2, 2, 2, 3}, // Unsafe
				},
			},
			expected: 0,
		},
		{
			name: "All unsafe",
			input: SolutionInput{
				Reports: [][]int{
					{92, 91, 92, 95, 95, 94},         // Unsafe
					{35, 35, 37, 39, 43, 43},         // Unsafe
					{28, 23, 22, 20, 18, 18, 11},     // Unsafe
					{38, 42, 45, 48, 47},             // Unsafe
					{5, 7, 9, 11, 17, 19, 22},        // Unsafe
					{48, 44, 42, 41, 40, 41, 39, 35}, // Unsafe
				},
			},
			expected: 0,
		},
		{
			name: "One safe, edge case at the end",
			input: SolutionInput{
				Reports: [][]int{
					{6, 8, 9, 11, 14, 12}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "All safe, multiple edge cases",
			input: SolutionInput{
				Reports: [][]int{
					{6, 8, 9, 11, 14, 12},          // Safe
					{31, 33, 36, 39, 42, 42},       // Safe
					{5, 6, 7, 9, 11, 13, 17},       // Safe
					{7, 9, 12, 15, 17, 18, 21, 27}, // Safe
				},
			},
			expected: 4,
		},
		{
			name: "One safe, edge case at the beginning",
			input: SolutionInput{
				Reports: [][]int{
					{6, 6, 8, 9, 11, 14}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "One safe, edge case at the end",
			input: SolutionInput{
				Reports: [][]int{
					{6, 8, 9, 11, 14, 14}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Safe, only one element would remain equal",
			input: SolutionInput{
				Reports: [][]int{
					{6, 6}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Safe, only one element would remain larger than 3",
			input: SolutionInput{
				Reports: [][]int{
					{6, 10}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Safe after removing middle",
			input: SolutionInput{
				Reports: [][]int{
					{6, 10, 7}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Safe after removing middle",
			input: SolutionInput{
				Reports: [][]int{
					{6, 6, 7}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Unsafe even after removal",
			input: SolutionInput{
				Reports: [][]int{
					{6, 6, 6}, // Safe
				},
			},
			expected: 0,
		},
		{
			name: "Unsafe even after removal 2",
			input: SolutionInput{
				Reports: [][]int{
					{6, 10, 6}, // Safe
				},
			},
			expected: 0,
		},
		{
			name: "Two elements 1",
			input: SolutionInput{
				Reports: [][]int{
					{59, 60}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Two elements 2",
			input: SolutionInput{
				Reports: [][]int{
					{12, 15}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Two elements 3",
			input: SolutionInput{
				Reports: [][]int{
					{10, 7}, // Safe
				},
			},
			expected: 1,
		},
		{
			name: "Repeated numbers and direction change 1",
			input: SolutionInput{
				Reports: [][]int{
					{58, 59, 60, 59, 60},
				},
			},
			expected: 0,
		},
		{
			name: "Repeated numbers and direction change 2",
			input: SolutionInput{
				Reports: [][]int{
					{40, 41, 44, 44, 45, 46, 46},
				},
			},
			expected: 0,
		},
		{
			name: "Repeated numbers and direction change 3",
			input: SolutionInput{
				Reports: [][]int{
					{77, 80, 80, 81, 88},
				},
			},
			expected: 0,
		},
		{
			name: "Large Differences Between Adjacent Numbers 1",
			input: SolutionInput{
				Reports: [][]int{
					{7, 9, 12, 15, 17, 18, 21, 27},
				},
			},
			expected: 1,
		},
		{
			name: "Large Differences Between Adjacent Numbers 2",
			input: SolutionInput{
				Reports: [][]int{
					{82, 83, 86, 92, 95},
				},
			},
			expected: 0,
		},
		{
			name: "Direction changes 1",
			input: SolutionInput{
				Reports: [][]int{
					{61, 58, 61, 58, 59, 62},
				},
			},
			expected: 0,
		},
		{
			name: "Direction changes 2",
			input: SolutionInput{
				Reports: [][]int{
					{24, 21, 22, 24, 25, 26, 27},
				},
			},
			expected: 1,
		},
		{
			name: "Long reports 1",
			input: SolutionInput{
				Reports: [][]int{
					{89, 87, 88, 92, 93, 96, 97, 95},
				},
			},
			expected: 0,
		},
		{
			name: "Long reports 2",
			input: SolutionInput{
				Reports: [][]int{
					{72, 75, 75, 78, 79, 83},
				},
			},
			expected: 0,
		},
		{
			name: "Reports That Become Valid After Removing First or Last Element 1",
			input: SolutionInput{
				Reports: [][]int{
					{6, 8, 9, 11, 14, 12},
				},
			},
			expected: 1,
		},
		{
			name: "Reports That Become Valid After Removing First or Last Element 2",
			input: SolutionInput{
				Reports: [][]int{
					{68, 70, 71, 70, 72, 75, 76, 76},
				},
			},
			expected: 0,
		},
		{
			name: "Three element reports 1",
			input: SolutionInput{
				Reports: [][]int{
					{10, 11, 15},
				},
			},
			expected: 1,
		},
		{
			name: "Three element reports 2",
			input: SolutionInput{
				Reports: [][]int{
					{10, 15, 11},
				},
			},
			expected: 1,
		},
		{
			name: "Three element reports 3",
			input: SolutionInput{
				Reports: [][]int{
					{15, 10, 11},
				},
			},
			expected: 1,
		},
		{
			name: "Three element reports 4",
			input: SolutionInput{
				Reports: [][]int{
					{11, 10, 11},
				},
			},
			expected: 1,
		},
		{
			name: "Three element reports 5",
			input: SolutionInput{
				Reports: [][]int{
					{11, 9, 11},
				},
			},
			expected: 1,
		},
		{
			name: "Reddit edge cases",
			input: SolutionInput{
				Reports: [][]int{
					{48, 46, 47, 49, 51, 54, 56},
					{1, 1, 2, 3, 4, 5},

					{1, 2, 3, 4, 5, 5},
					{5, 1, 2, 3, 4, 5},

					{1, 4, 3, 2, 1},
					{1, 6, 7, 8, 9},

					{1, 2, 3, 4, 3},
					{9, 8, 7, 6, 7},

					{7, 10, 8, 10, 11},
					{29, 28, 27, 25, 26, 25, 22, 20},
				},
			},
			expected: 10,
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

func Test_removeIndexFromSlice(t *testing.T) {
	tests := []struct {
		slice    []int
		index    int
		expected []int
	}{
		{
			slice:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			slice:    []int{1, 2, 3, 4, 5},
			index:    0,
			expected: []int{2, 3, 4, 5},
		},
		{
			slice:    []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := removeIndexFromSlice(tt.index, tt.slice)
			common.AssertEqualSlices[int](t, result, tt.expected)
		})
	}
}
