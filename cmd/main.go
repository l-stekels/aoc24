package main

import (
	"advent2024/challenges/day1"
	"advent2024/common"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
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
		1: func() error {
			fmt.Println("Day 1 answers are:")
			parser := &day1.Parser{}
			input, err := common.ReadInput[day1.SolutionInput](baseDir, common.Day1, common.Input1, parser)
			if err != nil {
				return err
			}
			fmt.Println("Part 1:")
			fmt.Printf("%d\n", day1.SolvePart1(input))
			fmt.Println("Part 2:")
			fmt.Printf("%d\n", day1.SolvePart2(input))
			return nil
		},
	}

	app.Action = func(context *cli.Context) error {
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

		dayNumber, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed: %v", err)
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
