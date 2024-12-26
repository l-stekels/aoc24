package day10

import (
	"advent2024/common"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		grid: Grid{
			data: [][]int{},
		},
	}
	stringRows := strings.Split(content, "\n")
	for _, stringRow := range stringRows {
		strHeights := strings.Split(stringRow, "")
		var row []int
		for _, strHeight := range strHeights {
			height, err := strconv.Atoi(strHeight)
			if err != nil {
				panic(err)
			}
			row = append(row, height)
		}
		result.grid.data = append(result.grid.data, row)
	}
	result.grid.rows = len(result.grid.data)
	result.grid.cols = len(result.grid.data[0])

	return result, nil
}

type SolutionInput struct {
	grid Grid
}

func (s SolutionInput) Validate() error {
	if s.grid.rows != s.grid.cols {
		return fmt.Errorf("grid is not square")
	}

	return nil
}

type Grid struct {
	data [][]int
	rows int
	cols int
}

func (g Grid) Get(pos Position) int {
	return g.data[pos.X][pos.Y]
}

func (g Grid) IsPositionValid(position Position) bool {
	return position.X >= 0 && position.X < g.rows && position.Y >= 0 && position.Y < g.cols
}

// Position represents a point in the grid
type Position common.Point

var (
	Directions = []Position{Up, Right, Down, Left}
	Up         = Position{X: -1}
	Right      = Position{Y: 1}
	Down       = Position{X: 1}
	Left       = Position{Y: -1}
)

type QueueEntry struct {
	pos  Position
	path Path
}

func newQueue() Queue {
	return Queue{items: []QueueEntry{}}
}

type Queue struct {
	items []QueueEntry
	size  int
}

func (q *Queue) Enqueue(entry QueueEntry) {
	q.items = append([]QueueEntry{entry}, q.items...)
	q.size++
}

func (q *Queue) Dequeue() QueueEntry {
	if q.size == 0 {
		panic("queue is empty")
	}
	entry := q.items[0]
	q.items = q.items[1:]
	q.size--

	return entry
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

type Path struct {
	set common.Set[Position]
}

func NewPath() Path {
	return Path{set: common.NewSet[Position]()}
}

func (p *Path) Add(position Position) {
	p.set.Add(position)
}

func (p *Path) Copy() Path {
	newPath := NewPath()
	newPath.set.Merge(p.set)

	return newPath
}

func (p *Path) Contains(position Position) bool {
	return p.set.Contains(position)
}

func calculateTrailhead(grid Grid, startPosition Position, callback func(position Position)) {
	queue := newQueue()
	visited := common.NewSet[string]()

	initialPath := NewPath()
	initialPath.Add(startPosition)
	queue.Enqueue(QueueEntry{pos: startPosition, path: initialPath})

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		currentHeight := grid.Get(current.pos)
		if currentHeight == 9 {
			callback(current.pos)
			continue
		}
		for _, direction := range Directions {
			nextPosition := Position{X: current.pos.X + direction.X, Y: current.pos.Y + direction.Y}
			if !grid.IsPositionValid(nextPosition) {
				continue
			}
			nextHeight := grid.Get(nextPosition)
			if nextHeight != currentHeight+1 {
				continue
			}
			if current.path.Contains(nextPosition) {
				continue
			}
			newPath := current.path.Copy()
			newPath.Add(nextPosition)
			state := fmt.Sprintf("%v-%v", nextPosition, newPath)
			if visited.Contains(state) {
				continue
			}
			visited.Add(state)
			queue.Enqueue(QueueEntry{pos: nextPosition, path: newPath})
		}
	}
}

func findTrailheads(grid Grid) []Position {
	trailheads := make([]Position, 0)
	for x, row := range grid.data {
		for y, cell := range row {
			if cell == 0 {
				trailheads = append(trailheads, Position{X: x, Y: y})
			}
		}
	}

	return trailheads
}

func SolvePart1(input SolutionInput) int {
	result := 0
	trailheads := findTrailheads(input.grid)

	for _, trailhead := range trailheads {
		reachableNines := common.NewSet[Position]()
		scoreCallback := func(position Position) {
			reachableNines.Add(position)
		}
		calculateTrailhead(input.grid, trailhead, scoreCallback)
		result += reachableNines.Size()
	}

	return result
}

func SolvePart2(input SolutionInput) int {
	result := 0
	trailheads := findTrailheads(input.grid)

	for _, trailhead := range trailheads {
		uniquePaths := 0
		countUniquePaths := func(position Position) {
			uniquePaths++
		}
		calculateTrailhead(input.grid, trailhead, countUniquePaths)
		result += uniquePaths
	}

	return result
}
