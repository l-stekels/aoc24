package day6

import (
	"advent2024/common"
	"fmt"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		Grid: common.NewGrid[rune](make([][]rune, 0)),
	}
	stringRows := strings.Split(content, "\n")
	for _, row := range stringRows {
		result.Grid.AddRow([]rune(row))
	}

	return result, nil
}

type SolutionInput struct {
	Grid common.Grid[rune]
}

type FacingDirection rune

func parseRune(r rune) FacingDirection {
	return FacingDirection(r)
}

const (
	UpFacing    FacingDirection = '^'
	RightFacing FacingDirection = '>'
	DownFacing  FacingDirection = 'v'
	LeftFacing  FacingDirection = '<'
)

type Guard struct {
	pos    common.Point
	dir    common.Direction
	facing FacingDirection
}

func (g *Guard) TurnRight() {
	switch g.dir {
	case common.Up:
		g.dir = common.Right
		g.facing = RightFacing
	case common.Right:
		g.dir = common.Down
		g.facing = DownFacing
	case common.Down:
		g.dir = common.Left
		g.facing = LeftFacing
	case common.Left:
		g.dir = common.Up
		g.facing = UpFacing
	}
}

func (g *Guard) Move(pos common.Point) {
	g.pos = pos
}

func (g *Guard) copy() Guard {
	return Guard{
		pos:    g.pos,
		dir:    g.dir,
		facing: g.facing,
	}
}

func (s SolutionInput) Validate() error {
	return nil
}

func SolvePart1(input SolutionInput) int {
	visited := common.NewSet[common.Point]()
	guard, err := findGuard(input.Grid)
	if err != nil {
		panic(err)
	}

	for {
		if !input.Grid.IsPositionValid(guard.pos) {
			break
		}
		visited.Add(guard.pos)

		nextPosition := common.Point{X: guard.pos.X + guard.dir.X(), Y: guard.pos.Y + guard.dir.Y()}
		if !input.Grid.IsPositionValid(nextPosition) {
			break
		}
		if input.Grid.Get(nextPosition) == '#' {
			guard.TurnRight()
			continue
		}
		guard.Move(nextPosition)
	}

	return visited.Size()
}

func findGuard(grid common.Grid[rune]) (Guard, error) {
	for grid.HasNext() {
		cell, pos := grid.Next()
		guardFacing := parseRune(cell)
		switch guardFacing {
		case UpFacing:
			return Guard{pos, common.Up, UpFacing}, nil
		case RightFacing:
			return Guard{pos, common.Right, RightFacing}, nil
		case DownFacing:
			return Guard{pos, common.Down, DownFacing}, nil
		case LeftFacing:
			return Guard{pos, common.Left, LeftFacing}, nil
		}
	}

	return Guard{}, fmt.Errorf("guard not found")
}

type State struct {
	pos common.Point
	dir common.Direction
}

func detectLoop(grid common.Grid[rune], guard Guard) bool {
	visited := common.NewSet[State]()

	for {
		if !grid.IsPositionValid(guard.pos) {
			break
		}
		currentState := State{pos: guard.pos, dir: guard.dir}
		alreadyVisited := visited.Contains(currentState)
		if alreadyVisited {
			return true
		}
		visited.Add(currentState)
		nextPosition := common.Point{X: guard.pos.X + guard.dir.X(), Y: guard.pos.Y + guard.dir.Y()}
		if !grid.IsPositionValid(nextPosition) {
			break
		}
		if grid.Get(nextPosition) == '#' {
			guard.TurnRight()
			continue
		}
		guard.Move(nextPosition)
	}

	return false
}

func SolvePart2(input SolutionInput) int {
	result := 0
	startGuard, err := findGuard(input.Grid)
	if err != nil {
		panic(err)
	}

	for input.Grid.HasNext() {
		cell, pos := input.Grid.Next()
		if cell == '#' || cell == '^' || cell == '>' || cell == 'v' || cell == '<' {
			continue
		}
		testMap := input.Grid.Copy()
		testMap.Set('#', pos)
		if detectLoop(testMap, startGuard.copy()) {
			result++
		}
	}

	return result
}
