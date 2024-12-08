package day2

import (
	"advent2024/common"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	lines := strings.Split(content, "\n")
	reports := make([][]int, len(lines))
	for j, line := range lines {
		stringInputs := strings.Split(line, " ")
		intInputs := make([]int, len(stringInputs))
		for i, stringInput := range stringInputs {
			value, err := strconv.Atoi(stringInput)
			if err != nil {
				return SolutionInput{}, fmt.Errorf("invalid value in line: %s", line)
			}
			intInputs[i] = value
		}
		reports[j] = intInputs
	}

	return SolutionInput{
		Reports: reports,
	}, nil
}

type SolutionInput struct {
	Reports [][]int
}

func (s SolutionInput) Validate() error {
	// Can not really validate this input since only thing that matters
	// is that it is a 2d int array
	// and that is already guaranteed by the type of Reports
	return nil
}

type LevelChange string

const (
	increase LevelChange = "increase"
	decrease LevelChange = "decrease"
	equal    LevelChange = "equal"
)

func SolvePart1(input SolutionInput) int {
	result := 0

	for _, report := range input.Reports {
		if reportValid(report, 1) {
			result += 1
		}
	}

	return result
}

func SolvePart2(input SolutionInput) int {
	result := 0

	for _, report := range input.Reports {
		if reportValid(report, 0) {
			result += 1
		}
	}

	return result
}

func reportValid(report []int, level int) bool {
	var currentChange, previousChange LevelChange
	reportLength := len(report)

	if reportLength <= 1 {
		return true
	}

	for i, current := range report {
		nextIndex := i + 1

		if nextIndex >= reportLength {
			break
		}
		next := report[nextIndex]
		switch {
		case current == next:
			currentChange = equal
		case current > next:
			currentChange = decrease
		case current < next:
			currentChange = increase
		}
		// For the first iteration assign both current change and previous change
		if i == 0 {
			previousChange = currentChange
		}
		if currentChange != previousChange {
			if level == 0 {
				// There are two possibilities here
				// Either remove the current index
				currentRemoved := removeIndexFromSlice(i, report)
				if reportValid(currentRemoved, level+1) {
					return true
				}
				// Or remove the next index
				nextRemoved := removeIndexFromSlice(nextIndex, report)
				if reportValid(nextRemoved, level+1) {
					return true
				}
				// Additional edge case where the first element can be removed and it would pass
				firstRemoved := removeIndexFromSlice(0, report)
				if reportValid(firstRemoved, level+1) {
					return true
				}
			}

			return false
		}
		changeAmount := common.AbsDiff(current, next)
		if changeAmount < 1 || changeAmount > 3 {
			if level == 0 {
				currentRemoved := removeIndexFromSlice(i, report)
				if reportValid(currentRemoved, level+1) {
					return true
				}
				nextRemoved := removeIndexFromSlice(nextIndex, report)
				if reportValid(nextRemoved, level+1) {
					return true
				}
			}

			return false
		}
	}

	return true
}

func removeIndexFromSlice(indexToRemove int, slice []int) []int {
	newSlice := append([]int{}, slice[:indexToRemove]...)
	return append(newSlice, slice[indexToRemove+1:]...)
}
