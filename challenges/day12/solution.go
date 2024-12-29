package day12

import (
	"advent2024/common"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		grid: common.NewGridFromElements[rune]([][]rune{}),
	}
	stringRows := strings.Split(content, "\n")
	for _, row := range stringRows {
		result.grid.AddRow([]rune(row))
	}

	return result, nil
}

type SolutionInput struct {
	grid common.Grid[rune]
}

func (s SolutionInput) Validate() error {
	return nil
}

type Region struct {
	Cells     common.Set[common.Point]
	PlantType rune
}

func SolvePart1(input SolutionInput) int {
	result := calculateTotalFencingPrice(input.grid)

	return result
}

func calculateTotalFencingPrice(grid common.Grid[rune]) int {
	totalPrice := 0
	visited := common.NewEmptyGrid[bool](grid.Rows(), grid.Cols())
	regions := make([]Region, 0)

	for grid.HasNext() {
		_, currentPosition := grid.Next()
		if visited.Get(currentPosition) {
			continue
		}
		region := findRegion(grid, currentPosition, &visited)
		regions = append(regions, region)
	}
	for _, region := range regions {
		totalPrice += calculateRegionPrice(region, grid)
	}

	return totalPrice
}

func calculateRegionPrice(region Region, garden common.Grid[rune]) int {
	area := region.Cells.Size()
	perimeter := calculatePerimeter(region, garden)

	return area * perimeter
}

func calculatePerimeter(region Region, garden common.Grid[rune]) int {
	perimeter := 0
	for cell := range region.Cells.All() {
		for _, direction := range common.Directions {
			adjacent := common.Point{X: cell.X + direction.X(), Y: cell.Y + direction.Y()}
			if !garden.IsPositionValid(adjacent) || garden.Get(adjacent) != region.PlantType {
				perimeter++
			}
		}
	}

	return perimeter
}

func findRegion(garden common.Grid[rune], start common.Point, visited *common.Grid[bool]) Region {
	plantType := garden.Get(start)
	region := Region{PlantType: plantType, Cells: common.NewSet[common.Point]()}
	queue := common.NewQueue[common.Point]()
	queue.Enqueue(start)
	for !queue.IsEmpty() {
		current := queue.Dequeue()
		if visited.Get(current) {
			continue
		}
		visited.Set(true, current)
		region.Cells.Add(current)
		for _, dir := range common.Directions {
			nextDirection := common.Point{X: current.X + dir.X(), Y: current.Y + dir.Y()}
			if !garden.IsPositionValid(nextDirection) {
				continue
			}
			if garden.Get(nextDirection) != region.PlantType {
				continue
			}
			if visited.Get(nextDirection) {
				continue
			}
			queue.Enqueue(nextDirection)
		}
	}

	return region
}

func SolvePart2(input SolutionInput) int {
	result := 0

	return result
}
