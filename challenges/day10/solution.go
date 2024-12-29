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
		grid: common.NewGridFromElements[int]([][]int{}),
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
		result.grid.AddRow(row)
	}

	return result, nil
}

type SolutionInput struct {
	grid common.Grid[int]
}

func (s SolutionInput) Validate() error {
	return nil
}

type QueueEntry struct {
	pos  common.Point
	path Path
}

type Path struct {
	set common.Set[common.Point]
}

func NewPath() Path {
	return Path{set: common.NewPointSet()}
}

func (p *Path) Add(position common.Point) {
	p.set.Add(position)
}

func (p *Path) Copy() Path {
	newPath := NewPath()
	newPath.set.Merge(p.set)

	return newPath
}

func (p *Path) Contains(position common.Point) bool {
	return p.set.Contains(position)
}

func calculateTrailhead(grid common.Grid[int], startPosition common.Point, callback func(position common.Point)) {
	queue := common.NewQueue[QueueEntry]()
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
		for _, direction := range common.Directions {
			nextPosition := common.Point{X: current.pos.X + direction.X(), Y: current.pos.Y + direction.Y()}
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

func findTrailheads(grid common.Grid[int]) []common.Point {
	trailheads := make([]common.Point, 0)
	for grid.HasNext() {
		cell, pos := grid.Next()
		if cell == 0 {
			trailheads = append(trailheads, pos)
		}
	}

	return trailheads
}

func SolvePart1(input SolutionInput) int {
	result := 0
	trailheads := findTrailheads(input.grid)

	for _, trailhead := range trailheads {
		reachableNines := common.NewPointSet()
		scoreCallback := func(position common.Point) {
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
		countUniquePaths := func(position common.Point) {
			uniquePaths++
		}
		calculateTrailhead(input.grid, trailhead, countUniquePaths)
		result += uniquePaths
	}

	return result
}
