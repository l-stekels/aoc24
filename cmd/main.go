package main

import (
	"advent2024/challenges/day1"
	"advent2024/challenges/day2"
	"advent2024/challenges/day3"
	"advent2024/challenges/day4"
	"advent2024/challenges/day5"
	"advent2024/challenges/day6"
	"advent2024/common"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "adventofcode"
	app.Usage = "Solutions for Advent of Code challenges 2024"
	app.Authors = []*cli.Author{
		{
			Name:  "Lauris",
			Email: "lauris@stekels.lv",
		},
	}
	baseDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	daySolutions := map[int]func() error{
		1: createSolutionFunc[day1.SolutionInput](baseDir, common.Day1, common.Input, &day1.Parser{}, day1.SolvePart1, day1.SolvePart2),
		2: createSolutionFunc[day2.SolutionInput](baseDir, common.Day2, common.Input, &day2.Parser{}, day2.SolvePart1, day2.SolvePart2),
		3: createSolutionFunc[day3.SolutionInput](baseDir, common.Day3, common.Input, &day3.Parser{}, day3.SolvePart1, day3.SolvePart2),
		4: createSolutionFunc[day4.SolutionInput](baseDir, common.Day4, common.Input, &day4.Parser{}, day4.SolvePart1, day4.SolvePart2),
		5: createSolutionFunc[day5.SolutionInput](baseDir, common.Day5, common.Input, &day5.Parser{}, day5.SolvePart1, day5.SolvePart2),
		6: createSolutionFunc[day6.SolutionInput](baseDir, common.Day6, common.Input, &day6.Parser{}, day6.SolvePart1, day6.SolvePart2),
	}

	app.Action = func(context *cli.Context) error {
		var dayNumber string
		if context.NArg() > 0 {
			dayNumber = context.Args().Get(0)
		} else {
			prompt := promptui.Prompt{
				Label: "Enter the day number",
				Validate: func(input string) error {
					_, err := strconv.Atoi(input)
					if err != nil {
						return fmt.Errorf("invalid day number")
					}
					return nil
				},
			}
			dayNumber, err = prompt.Run()
			if err != nil {
				return fmt.Errorf("prompt failed: %v", err)
			}
		}

		day, err := strconv.Atoi(dayNumber)
		if err != nil {
			return fmt.Errorf("invalid day number: %s", dayNumber)
		}

		if solution, exists := daySolutions[day]; exists {
			return solution()
		}
		return fmt.Errorf("no solution available for day %d", day)
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func createSolutionFunc[T common.SolutionInput](
	baseDir string,
	day common.ChallengeDay,
	input common.ChallengeInput,
	parser common.SolutionParser[T],
	solvePart1 func(T) int,
	solvePart2 func(T) int,
) func() error {
	return func() error {
		fmt.Printf("Day %s answers are:\n", day.String())
		input, err := common.ReadInput[T](baseDir, day, input, parser)
		if err != nil {
			return err
		}
		fmt.Println("Part 1:")
		part1Start := time.Now()
		fmt.Printf("%d\n", solvePart1(input))
		duration := time.Since(part1Start)
		fmt.Printf("Time taken: %.6f ms\n", float64(duration.Nanoseconds())/1e6)
		fmt.Println("Part 2:")
		part2Start := time.Now()
		fmt.Printf("%d\n", solvePart2(input))
		duration = time.Since(part2Start)
		fmt.Printf("Time taken: %.6f ms\n", float64(duration.Nanoseconds())/1e6)
		return nil
	}
}
