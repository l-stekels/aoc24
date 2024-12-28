package solution

import (
	"advent2024/challenges/day1"
	"advent2024/challenges/day10"
	"advent2024/challenges/day11"
	"advent2024/challenges/day12"
	"advent2024/challenges/day13"
	"advent2024/challenges/day14"
	"advent2024/challenges/day15"
	"advent2024/challenges/day2"
	"advent2024/challenges/day3"
	"advent2024/challenges/day4"
	"advent2024/challenges/day5"
	"advent2024/challenges/day6"
	"advent2024/challenges/day7"
	"advent2024/challenges/day8"
	"advent2024/challenges/day9"
	"advent2024/common"
	"fmt"
	"strings"
	"time"
)

type Map struct {
	daySolutions map[int]func() error
}

func NewSolutionMap(baseDir string) Map {
	return Map{
		daySolutions: map[int]func() error{
			1:  createSolutionFunc[day1.SolutionInput, int](baseDir, common.Day1, common.Input, &day1.Parser{}, day1.SolvePart1, day1.SolvePart2),
			2:  createSolutionFunc[day2.SolutionInput, int](baseDir, common.Day2, common.Input, &day2.Parser{}, day2.SolvePart1, day2.SolvePart2),
			3:  createSolutionFunc[day3.SolutionInput, int](baseDir, common.Day3, common.Input, &day3.Parser{}, day3.SolvePart1, day3.SolvePart2),
			4:  createSolutionFunc[day4.SolutionInput, int](baseDir, common.Day4, common.Input, &day4.Parser{}, day4.SolvePart1, day4.SolvePart2),
			5:  createSolutionFunc[day5.SolutionInput, int](baseDir, common.Day5, common.Input, &day5.Parser{}, day5.SolvePart1, day5.SolvePart2),
			6:  createSolutionFunc[day6.SolutionInput, int](baseDir, common.Day6, common.Input, &day6.Parser{}, day6.SolvePart1, day6.SolvePart2),
			7:  createSolutionFunc[day7.SolutionInput, uint64](baseDir, common.Day7, common.Input, &day7.Parser{}, day7.SolvePart1, day7.SolvePart2),
			8:  createSolutionFunc[day8.SolutionInput, int](baseDir, common.Day8, common.Input, &day8.Parser{}, day8.SolvePart1, day8.SolvePart2),
			9:  createSolutionFunc[day9.SolutionInput, uint64](baseDir, common.Day9, common.Input, &day9.Parser{}, day9.SolvePart1, day9.SolvePart2),
			10: createSolutionFunc[day10.SolutionInput, int](baseDir, common.Day10, common.Input, &day10.Parser{}, day10.SolvePart1, day10.SolvePart2),
			11: createSolutionFunc[day11.SolutionInput, int](baseDir, common.Day11, common.Input, &day11.Parser{}, day11.SolvePart1, day11.SolvePart2),
			12: createSolutionFunc[day12.SolutionInput, int](baseDir, common.Day12, common.Input, &day12.Parser{}, day12.SolvePart1, day12.SolvePart2),
			13: createSolutionFunc[day13.SolutionInput, int](baseDir, common.Day13, common.Input, &day13.Parser{}, day13.SolvePart1, day13.SolvePart2),
			14: createSolutionFunc[day14.SolutionInput, int](baseDir, common.Day14, common.Input, &day14.Parser{}, day14.SolvePart1, day14.SolvePart2),
			15: createSolutionFunc[day15.SolutionInput, int](baseDir, common.Day15, common.Input, &day15.Parser{}, day15.SolvePart1, day15.SolvePart2),
		},
	}
}

func (s Map) Run(day int) error {
	solution, ok := s.daySolutions[day]
	if !ok {
		return fmt.Errorf("no solution for day %d", day)
	}
	return solution()
}

type Number interface {
	~int | ~uint64
}

func createSolutionFunc[T common.SolutionInput, R Number](
	baseDir string,
	day common.ChallengeDay,
	input common.ChallengeInput,
	parser common.SolutionParser[T],
	solvePart1 func(T) R,
	solvePart2 func(T) R,
) func() error {
	return func() error {
		fmt.Printf("Day %s answers are:\n", day.String())
		input, err := common.ReadInput[T](baseDir, day, input, parser)
		if err != nil {
			return err
		}
		fmt.Println("Part 1:")
		part1Start := time.Now()
		fmt.Printf("%v\n", solvePart1(input))
		duration := time.Since(part1Start)
		fmt.Printf("Time taken: %s\n", formatDuration(duration))
		fmt.Println("Part 2:")
		part2Start := time.Now()
		fmt.Printf("%v\n", solvePart2(input))
		duration = time.Since(part2Start)
		fmt.Printf("Time taken: %s\n", formatDuration(duration))
		return nil
	}
}

func formatDuration(d time.Duration) string {
	nanoseconds := d.Nanoseconds()

	if nanoseconds == 0 {
		return "0ns"
	}

	seconds := nanoseconds / 1e9
	nanoseconds = nanoseconds % 1e9
	milliseconds := nanoseconds / 1e6
	nanoseconds = nanoseconds % 1e6
	microseconds := nanoseconds / 1e3
	nanoseconds = nanoseconds % 1e3

	var parts []string
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%ds", seconds))
	}
	if milliseconds > 0 {
		parts = append(parts, fmt.Sprintf("%dms", milliseconds))
	}
	if microseconds > 0 {
		parts = append(parts, fmt.Sprintf("%dµs", microseconds))
	}
	if nanoseconds > 0 {
		parts = append(parts, fmt.Sprintf("%dns", nanoseconds))
	}

	return strings.Join(parts, " ")
}
