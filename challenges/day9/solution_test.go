package day9

import (
	"testing"
)

func TestParser_CreateSolutionInput(t *testing.T) {
	parser := &Parser{}
	tests := []struct {
		name     string
		input    string
		expected SolutionInput
	}{
		{
			name:     "Example",
			input:    "12345",
			expected: getExampleInput(1),
		},
		{
			name:     "Example 2",
			input:    "2333133121414131402",
			expected: getExampleInput(2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.CreateSolutionInput(tt.input)
			if err != nil {
				t.Fatalf("CreateSolutionInput failed: %v", err)
			}

			if err := result.Validate(); err != nil {
				t.Fatalf("Validate failed: %v", err)
			}
			if tt.expected.String() != result.String() {
				t.Errorf("CreateSolutionInput failed: want %s, got %s", tt.expected.String(), result.String())
			}
		})
	}
}

func Test_SolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		input    SolutionInput
		expected uint64
	}{
		{
			name:     "Example",
			input:    getExampleInput(1),
			expected: uint64(60),
		},
		{
			name:     "Example",
			input:    getExampleInput(2),
			expected: uint64(1928),
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
			name:     "Example",
			input:    getExampleInput(1),
			expected: uint64(132),
		},
		{
			name:     "Example",
			input:    getExampleInput(2),
			expected: uint64(2858),
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

func getExampleInput(n int) SolutionInput {
	if n == 1 {
		return SolutionInput{
			totalLength: 15,
			fileSegment: &DiskSegment{
				start:  0,
				length: 1,
				fileId: 0,
				left:   nil,
				right: &DiskSegment{
					start:  3,
					length: 3,
					fileId: 1,
					left:   nil,
					right: &DiskSegment{
						start:  10,
						length: 5,
						fileId: 2,
						left:   nil,
						right:  nil,
					},
				},
			},
			emptySegment: &DiskSegment{
				start:  1,
				length: 2,
				fileId: -1,
				left:   nil,
				right: &DiskSegment{
					start:  6,
					length: 4,
					fileId: -1,
					left:   nil,
					right:  nil,
				},
			},
		}
	}
	// 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41
	// 0  0  .  .  .  1  1  1  .  .  .  2  .  .  .  3  3  3  .  4  4  .  5  5  5  5  .  6  6  6  6  .  7  7  7  .  8  8  8  8  9  9
	return SolutionInput{
		totalLength: 42,
		fileSegment: &DiskSegment{
			start:  0,
			length: 2,
			fileId: 0,
			left:   nil,
			right: &DiskSegment{
				start:  5,
				length: 3,
				fileId: 1,
				left:   nil,
				right: &DiskSegment{
					start:  11,
					length: 1,
					fileId: 2,
					left:   nil,
					right: &DiskSegment{
						start:  15,
						length: 3,
						fileId: 3,
						left:   nil,
						right: &DiskSegment{
							start:  19,
							length: 2,
							fileId: 4,
							left:   nil,
							right: &DiskSegment{
								start:  22,
								length: 4,
								fileId: 5,
								left:   nil,
								right: &DiskSegment{
									start:  27,
									length: 4,
									fileId: 6,
									left:   nil,
									right: &DiskSegment{
										start:  32,
										length: 3,
										fileId: 7,
										left:   nil,
										right: &DiskSegment{
											start:  36,
											length: 4,
											fileId: 8,
											left:   nil,
											right: &DiskSegment{
												start:  40,
												length: 2,
												fileId: 9,
												left:   nil,
												right:  nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		emptySegment: &DiskSegment{
			start:  2,
			length: 3,
			fileId: -1,
			left:   nil,
			right: &DiskSegment{
				start:  8,
				length: 3,
				fileId: -1,
				left:   nil,
				right: &DiskSegment{
					start:  12,
					length: 3,
					fileId: -1,
					left:   nil,
					right: &DiskSegment{
						start:  18,
						length: 1,
						fileId: -1,
						left:   nil,
						right: &DiskSegment{
							start:  21,
							length: 1,
							fileId: -1,
							left:   nil,
							right: &DiskSegment{
								start:  26,
								length: 1,
								fileId: -1,
								left:   nil,
								right: &DiskSegment{
									start:  31,
									length: 1,
									fileId: -1,
									left:   nil,
									right: &DiskSegment{
										start:  35,
										length: 1,
										fileId: -1,
										left:   nil,
										right:  nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	//00212111222
	//0 * 0 + 1 * 0 + 2 * 2 + 3 * 1 + 4 * 1 + 5 * 1 + 6 * 0 + 7 * 0 + 8 * 2 + 9 * 2 + 10 * 2
}
