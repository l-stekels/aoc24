package day6

import (
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

type Set[T comparable] struct {
	data map[T]bool
}

func (s *Set[T]) Add(element T) {
	_, ok := s.data[element]
	if !ok {
		s.data[element] = true
	}
}

func (s *Set[T]) contains(element T) bool {
	_, ok := s.data[element]

	return ok
}

type Position struct {
	x, y uint
}

func NewPosition(x int, y int) *Position {
	if x < 0 || y < 0 {
		panic(fmt.Sprintf("Invalid position: %d, %d", x, y))
	}
	return &Position{x: uint(x), y: uint(y)}
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
	pos    Position
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

func (g *Guard) Move(pos Position) {
	g.pos = pos
}

func clearScreen() {
	print("\033[H\033[2J")
}

func printGrid(grid [][]rune, visited map[Position]bool, guard Guard) {
	clearScreen()
	for x := range grid {
		for y := range grid[x] {
			pos := *NewPosition(x, y)
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
	visited := Set[Position]{data: map[Position]bool{}}
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

		nextPosition := Position{x: guard.pos.x + uint(guard.dir.dx), y: guard.pos.y + uint(guard.dir.dy)}
		if !isInBounds(nextPosition, input.Map) {
			break
		}
		if input.Map[nextPosition.x][nextPosition.y] == '#' {
			guard.TurnRight()
			continue
		}
		guard.Move(nextPosition)
		//printGrid(input.Map, visited.data, guard)
	}
	//printGrid(input.Map, visited.data, guard)

	return len(visited.data)
}

func isInBounds(pos Position, m [][]rune) bool {
	return pos.x < uint(len(m[0])) && pos.y < uint(len(m))
}

func findGuard(m [][]rune) (Guard, error) {
	for x, row := range m {
		for y, cell := range row {
			switch cell {
			case '^':
				return Guard{*NewPosition(x, y), Up, UpFacing}, nil
			case '>':
				return Guard{*NewPosition(x, y), Right, RightFacing}, nil
			case 'v':
				return Guard{*NewPosition(x, y), Down, DownFacing}, nil
			case '<':
				return Guard{*NewPosition(x, y), Left, LeftFacing}, nil
			}
		}
	}

	return Guard{}, fmt.Errorf("guard not found")
}

type State struct {
	pos Position
	dir Direction
}

func detectLoop(grid [][]rune, guard Guard) bool {
	visited := Set[State]{data: map[State]bool{}}

	for {
		if !isInBounds(guard.pos, grid) {
			break
		}
		currentState := State{pos: guard.pos, dir: guard.dir}
		alreadyVisited := visited.contains(currentState)
		if alreadyVisited {
			return true
		}
		visited.Add(currentState)
		nextPosition := Position{x: guard.pos.x + uint(guard.dir.dx), y: guard.pos.y + uint(guard.dir.dy)}
		if !isInBounds(nextPosition, grid) {
			break
		}
		if grid[nextPosition.x][nextPosition.y] == '#' {
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
