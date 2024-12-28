package common

type ChallengeDay string

const (
	Day1  ChallengeDay = "day1"
	Day2  ChallengeDay = "day2"
	Day3  ChallengeDay = "day3"
	Day4  ChallengeDay = "day4"
	Day5  ChallengeDay = "day5"
	Day6  ChallengeDay = "day6"
	Day7  ChallengeDay = "day7"
	Day8  ChallengeDay = "day8"
	Day9  ChallengeDay = "day9"
	Day10 ChallengeDay = "day10"
	Day11 ChallengeDay = "day11"
	Day12 ChallengeDay = "day12"
	Day13 ChallengeDay = "day13"
	Day14 ChallengeDay = "day14"
	Day15 ChallengeDay = "day15"
)

func (d ChallengeDay) String() string {
	return string(d)
}

type SolutionInput interface {
	Validate() error
}

type SolutionParser[T SolutionInput] interface {
	CreateSolutionInput(content string) (T, error)
}
