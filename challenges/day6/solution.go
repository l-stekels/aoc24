package day6

import (
	"advent2024/common"
	"fmt"
	"strings"
	"time"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		Map: [][]rune{},
	}
	stringRows := strings.Split(content, "\n")
	for x, row := range stringRows {
		result.Map = append(result.Map, []rune{})
		for _, cell := range row {
			result.Map[x] = append(result.Map[x], cell)
		}
	}

	return result, nil
}

type SolutionInput struct {
	Map [][]rune
}

type Direction struct {
	dx, dy int
}

var (
	Up    = Direction{-1, 0}
	Right = Direction{0, 1}
	Down  = Direction{1, 0}
	Left  = Direction{0, -1}
)

type FacingDirection rune

const (
	UpFacing    FacingDirection = '^'
	RightFacing FacingDirection = '>'
	DownFacing  FacingDirection = 'v'
	LeftFacing  FacingDirection = '<'
)

type Guard struct {
	pos    common.Point
	dir    Direction
	facing FacingDirection
}

func (g *Guard) TurnRight() {
	switch g.dir {
	case Up:
		g.dir = Right
		g.facing = RightFacing
	case Right:
		g.dir = Down
		g.facing = DownFacing
	case Down:
		g.dir = Left
		g.facing = LeftFacing
	case Left:
		g.dir = Up
		g.facing = UpFacing
	}
}

func (g *Guard) Move(pos common.Point) {
	g.pos = pos
}

func clearScreen() {
	print("\033[H\033[2J")
}

func printGrid(grid [][]rune, visited map[common.Point]bool, guard Guard) {
	clearScreen()
	for x := range grid {
		for y := range grid[x] {
			pos := common.Point{X: x, Y: y}
			if visited[pos] {
				fmt.Print("X")
			} else if guard.pos == pos {
				fmt.Print(string(guard.facing))
			} else {
				fmt.Print(string(grid[x][y]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(100 * time.Millisecond) // Add delay to make movement visible
}

func (s SolutionInput) Validate() error {
	return nil
}

func SolvePart1(input SolutionInput) int {
	visited := common.NewSet[common.Point]()
	guard, err := findGuard(input.Map)
	if err != nil {
		panic(err)
	}
	//printGrid(input.Map, visited.data, guard)

	for {
		if !isInBounds(guard.pos, input.Map) {
			break
		}
		visited.Add(guard.pos)

		nextPosition := common.Point{X: guard.pos.X + guard.dir.dx, Y: guard.pos.Y + guard.dir.dy}
		if !isInBounds(nextPosition, input.Map) {
			break
		}
		if input.Map[nextPosition.X][nextPosition.Y] == '#' {
			guard.TurnRight()
			continue
		}
		guard.Move(nextPosition)
		//printGrid(input.Map, visited.data, guard)
	}
	//printGrid(input.Map, visited.data, guard)

	return visited.Length()
}

func isInBounds(pos common.Point, m [][]rune) bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}
	return pos.X < len(m[0]) && pos.Y < len(m)
}

func findGuard(m [][]rune) (Guard, error) {
	for x, row := range m {
		for y, cell := range row {
			point := common.Point{X: x, Y: y}
			switch cell {
			case '^':
				return Guard{pos: point, dir: Up, facing: UpFacing}, nil
			case '>':
				return Guard{point, Right, RightFacing}, nil
			case 'v':
				return Guard{point, Down, DownFacing}, nil
			case '<':
				return Guard{point, Left, LeftFacing}, nil
			}
		}
	}

	return Guard{}, fmt.Errorf("guard not found")
}

type State struct {
	pos common.Point
	dir Direction
}

func detectLoop(grid [][]rune, guard Guard) bool {
	visited := common.NewSet[State]()

	for {
		if !isInBounds(guard.pos, grid) {
			break
		}
		currentState := State{pos: guard.pos, dir: guard.dir}
		alreadyVisited := visited.Contains(currentState)
		if alreadyVisited {
			return true
		}
		visited.Add(currentState)
		nextPosition := common.Point{X: guard.pos.X + guard.dir.dx, Y: guard.pos.Y + guard.dir.dy}
		if !isInBounds(nextPosition, grid) {
			break
		}
		if grid[nextPosition.X][nextPosition.Y] == '#' {
			guard.TurnRight()
			continue
		}
		guard.Move(nextPosition)
	}

	return false
}

func SolvePart2(input SolutionInput) int {
	result := 0
	startGuard, err := findGuard(input.Map)
	if err != nil {
		panic(err)
	}

	for x, row := range input.Map {
		for y, cell := range row {
			if cell == '#' || cell == '^' || cell == '>' || cell == 'v' || cell == '<' {
				continue
			}
			testMap := copyMap(input.Map)
			testMap[x][y] = '#'
			if detectLoop(testMap, Guard{pos: startGuard.pos, dir: startGuard.dir, facing: startGuard.facing}) {
				result++
			}
		}
	}

	return result
}

func copyMap(m [][]rune) [][]rune {
	var result [][]rune
	for _, row := range m {
		result = append(result, append([]rune{}, row...))
	}

	return result
}
