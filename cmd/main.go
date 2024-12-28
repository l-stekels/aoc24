package main

import (
	"advent2024/solution"
	"fmt"
	"os"
)

func main() {
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	solutionMap := solution.NewSolutionMap(baseDir)

	switch len(os.Args) {
	case 1: // No arguments, run all solutions
		solutionMap.RunAll()
	case 2: // One argument, run specific day
		err := solutionMap.Run(os.Args[1])
		if err != nil {
			fmt.Println("Error running solution:", err)
		}
	default: // More than one argument is not supported
		fmt.Println("Too many arguments.")
		fmt.Println("Usage: adventofcode [day]")
		fmt.Println("If no day is specified, all solutions will be run.")
		os.Exit(1)
	}
}
