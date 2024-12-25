package main

import (
	"advent2024/solution"
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

	solutionMap := solution.NewSolutionMap(baseDir)

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
		err = solutionMap.Run(day)
		if err != nil {
			return fmt.Errorf("no solution available for day %d", day)
		}
		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
