package common

type ChallengeDay string

type ChallengeInput string

const (
	Day1 ChallengeDay = "day1"
	Day2 ChallengeDay = "day2"
	Day3 ChallengeDay = "day3"
)

const (
	Input1 ChallengeInput = "input1"
)

func (d ChallengeDay) String() string {
	return string(d)
}

func (i ChallengeInput) String() string {
	return string(i)
}

type SolutionInput interface {
	Validate() error
}

type SolutionParser[T SolutionInput] interface {
	CreateSolutionInput(content string) (T, error)
}
