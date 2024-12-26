package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		stones: make([]Stone, 0),
	}
	stringValues := strings.Split(content, " ")
	for _, stringValue := range stringValues {
		val := strings.TrimSpace(stringValue)
		value, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return SolutionInput{}, fmt.Errorf("failed to parse stone value %s: %v", stringValue, err)
		}
		result.stones = append(result.stones, Stone{value: value})
	}

	return result, nil
}

type SolutionInput struct {
	stones []Stone
}

type Stone struct {
	value uint64
}

func (s Stone) String() string {
	return fmt.Sprintf("%d", s.value)
}

func (s SolutionInput) Validate() error {
	return nil
}

func SolvePart1(input SolutionInput) int {
	patterns := make([]StonePattern, len(input.stones))
	for i, stone := range input.stones {
		patterns[i] = StonePattern{value: stone.value, count: 1}
	}
	for i := 0; i < 25; i++ {
		patterns = transformStones(patterns)
	}

	var totalCount uint64
	for _, pattern := range patterns {
		totalCount += pattern.count
	}

	return int(totalCount)
}

type StonePattern struct {
	value uint64
	count uint64
}

func transformStones(patterns []StonePattern) []StonePattern {
	newPatterns := make([]StonePattern, 0, len(patterns)*2)
	for _, pattern := range patterns {
		if pattern.value == 0 {
			newPatterns = append(newPatterns, StonePattern{
				value: 1,
				count: pattern.count,
			})
			continue
		}

		digitCount := 1
		temp := pattern.value
		for temp >= 10 {
			digitCount++
			temp /= 10
		}
		if digitCount%2 == 1 {
			newPatterns = append(newPatterns, StonePattern{
				value: pattern.value * 2024,
				count: pattern.count,
			})
			continue
		}
		divisor := uint64(1)
		for i := 0; i < digitCount/2; i++ {
			divisor *= 10
		}
		newPatterns = append(newPatterns,
			StonePattern{value: pattern.value / divisor, count: pattern.count},
			StonePattern{value: pattern.value % divisor, count: pattern.count},
		)
	}
	patterns = mergePatterns(newPatterns)

	return patterns
}

func mergePatterns(patterns []StonePattern) []StonePattern {
	valueMap := make(map[uint64]uint64)

	for _, pattern := range patterns {
		valueMap[pattern.value] += pattern.count
	}

	result := make([]StonePattern, 0, len(valueMap))
	for value, count := range valueMap {
		result = append(result, StonePattern{value: value, count: count})
	}

	return result
}

func SolvePart2(input SolutionInput) int {
	patterns := make([]StonePattern, len(input.stones))
	for i, stone := range input.stones {
		patterns[i] = StonePattern{value: stone.value, count: 1}
	}

	for i := 0; i < 75; i++ {
		patterns = transformStones(patterns)
	}

	var totalCount uint64
	for _, pattern := range patterns {
		totalCount += pattern.count
	}

	return int(totalCount)
}
