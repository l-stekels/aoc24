package day1

import (
	"regexp"
	"strconv"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	commandRegexp := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	matches := commandRegexp.FindAllString(content, -1)

	return SolutionInput{
		Commands: matches,
	}, nil
}

type SolutionInput struct {
	Commands []string
}

func (s SolutionInput) Validate() error {
	return nil
}

func SolvePart1(input SolutionInput) int {
	result := 0

	numberRegexp := regexp.MustCompile(`\d{1,3}`)
	for _, command := range input.Commands {
		matches := numberRegexp.FindAllString(command, -1)
		// This should solve issues when the command is "do()" or "don't()"
		if len(matches) != 2 {
			continue
		}
		first, err := strconv.Atoi(matches[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		result += first * second
	}

	return result
}

func SolvePart2(input SolutionInput) int {
	result := 0

	return result
}
