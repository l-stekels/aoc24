package day7

import (
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		equations: []Equation{},
	}
	content = strings.TrimSpace(content)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		value := strings.TrimSpace(parts[0])
		firstNum, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		stringTerms := strings.Fields(strings.TrimSpace(parts[1]))
		terms := make([]int, 0, len(stringTerms))

		for _, term := range stringTerms {
			num, err := strconv.Atoi(term)
			if err != nil {
				continue
			}
			terms = append(terms, num)
		}

		if len(terms) > 0 {
			result.equations = append(result.equations, Equation{
				terms:     terms,
				testValue: uint64(firstNum),
			})
		}
	}

	return result, nil
}

type SolutionInput struct {
	equations []Equation
}

type Equation struct {
	terms     []int
	testValue uint64
}

func (e Equation) Equals(other Equation) bool {
	if e.testValue != other.testValue {
		return false
	}
	if len(e.terms) != len(other.terms) {
		return false
	}
	for i, term := range e.terms {
		if term != other.terms[i] {
			return false
		}
	}
	return true
}

func (s SolutionInput) Validate() error {
	return nil
}

func SolvePart1(input SolutionInput) uint64 {
	var result uint64

	for _, equation := range input.equations {
		testValue := equation.testValue
		if isSolvable(equation.terms, testValue) {
			result = result + equation.testValue
		}
	}

	return result
}

func isSolvable(numbers []int, testValue uint64) bool {
	operators := []rune{'+', '*'}
	operatorSets := generateOperatorSets(operators, len(numbers)-1)

	for _, set := range operatorSets {
		if evaluateExpression(numbers, set) == testValue {
			return true
		}
	}

	return false
}

func evaluateExpression(numbers []int, operators []rune) uint64 {
	result := uint64(numbers[0])

	for i, op := range operators {
		next := uint64(numbers[i+1])
		switch op {
		case '+':
			result = result + next
		case '*':
			result = result * next
		}
	}

	return result
}

func generateOperatorSets(operators []rune, length int) [][]rune {
	result := make([][]rune, 0)

	var generateCombos func(current []rune, position int)
	generateCombos = func(current []rune, position int) {
		if position == length {
			currentCopy := make([]rune, len(current))
			copy(currentCopy, current)
			result = append(result, currentCopy)
			return
		}

		for _, op := range operators {
			generateCombos(append(current, op), position+1)
		}
	}

	generateCombos([]rune{}, 0)
	return result
}

func SolvePart2(input SolutionInput) uint64 {
	result := uint64(0)

	return result
}
