package day8

import (
	"advent2024/common"
	"fmt"
	"strings"
	"sync"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	lines := strings.Split(content, "\n")
	result := SolutionInput{
		grid:     common.NewGrid([][]rune{}),
		antennas: map[Frequency][]Antenna{},
	}

	for x, line := range lines {
		result.grid.AddRow([]rune(line))
		symbols := strings.Split(line, "")
		for y, symbol := range symbols {
			gridElement := rune(symbol[0])
			if gridElement == '.' {
				continue
			}
			freq := Frequency(gridElement)
			if _, ok := result.antennas[freq]; !ok {
				result.antennas[freq] = []Antenna{}
			}
			result.antennas[freq] = append(
				result.antennas[freq],
				Antenna{pos: common.Point{X: x, Y: y}, freq: string(freq)},
			)
		}
	}

	return result, nil
}

type SolutionInput struct {
	grid     common.Grid[rune]
	antennas map[Frequency][]Antenna
}

type Antenna struct {
	pos  common.Point
	freq string
}

type Frequency rune

func (f Frequency) String() string {
	return string(f)
}

func (s SolutionInput) Validate() error {
	if s.grid.Rows() == 0 {
		return fmt.Errorf("grid is empty")
	}
	if len(s.antennas) == 0 {
		return fmt.Errorf("antennas are empty")
	}

	return nil
}

func SolvePart1(input SolutionInput) int {
	width := input.grid.Rows()
	length := input.grid.Cols()
	uniqueAntiNodePositions := common.NewSet[common.Point]()

	for _, antennas := range input.antennas {
		for i, a1 := range antennas {
			// check each pair of antennas
			for j, a2 := range antennas {
				if i == j {
					continue
				}
				findAntiNodesForPair(a1, a2, width, length, &uniqueAntiNodePositions)
			}
		}
	}

	return uniqueAntiNodePositions.Size()
}

func findAntiNodesForPair(a1 Antenna, a2 Antenna, width int, length int, positions *common.Set[common.Point]) {
	// calculate the distance between the two antennas
	dx := a2.pos.X - a1.pos.X
	dy := a2.pos.Y - a1.pos.Y

	p1 := common.Point{
		X: a1.pos.X - dx,
		Y: a1.pos.Y - dy,
	}
	if inBounds(p1, width, length) {
		positions.Add(p1)
	}
	p2 := common.Point{
		X: a2.pos.X + dx,
		Y: a2.pos.Y + dy,
	}
	if inBounds(p2, width, length) {
		positions.Add(p2)
	}
}

func inBounds(p1 common.Point, width, length int) bool {
	return p1.X >= 0 && p1.X < width && p1.Y >= 0 && p1.Y < length
}

func SolvePart2(input SolutionInput) int {
	width := input.grid.Rows()
	length := input.grid.Cols()
	uniqueAntiNodePositions := common.NewSet[common.Point]()

	for _, antennas := range input.antennas {
		for i, a1 := range antennas {
			// check each pair of antennas
			for j, a2 := range antennas {
				if i == j {
					continue
				}
				uniqueAntiNodePositions.Merge(findLinePoints(a1, a2, width, length))
			}
		}
	}

	return uniqueAntiNodePositions.Size()
}

func findLinePoints(a1 Antenna, a2 Antenna, width int, length int) common.Set[common.Point] {
	points := common.NewSet[common.Point]()
	dx := a2.pos.X - a1.pos.X
	dy := a2.pos.Y - a1.pos.Y
	steps := common.Gcd(common.Abs(dx), common.Abs(dy))
	stepX := dx / steps
	stepY := dy / steps

	currentX := a1.pos.X
	currentY := a1.pos.Y
	ch := make(chan common.Point)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		findPointsStartingFrom(currentX, currentY, stepX, stepY, width, length, ch)
	}()
	currentX = a2.pos.X
	currentY = a2.pos.Y
	go func() {
		defer wg.Done()
		findPointsStartingFrom(currentX, currentY, -stepX, -stepY, width, length, ch)
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	for point := range ch {
		points.Add(point)
	}

	return points
}

func findPointsStartingFrom(currentX, currentY, stepX, stepY, width, length int, ch chan common.Point) {
	for {
		if !inBounds(common.Point{X: currentX, Y: currentY}, width, length) {
			break
		}
		ch <- common.Point{X: currentX, Y: currentY}
		currentX += stepX
		currentY += stepY
	}
}
