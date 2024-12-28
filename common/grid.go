package common

type Grid[T int | rune] struct {
	data       [][]T
	rows       int
	cols       int
	currentRow int
	currentCol int
}

func NewGrid[T int | rune](elements [][]T) Grid[T] {
	if len(elements) == 0 {
		data := make([][]T, 0)
		for i := range data {
			data[i] = make([]T, 0)
		}
		return Grid[T]{data: data}
	}

	return Grid[T]{data: elements, rows: len(elements), cols: len(elements[0])}
}

func (g *Grid[T]) Rows() int {
	return g.rows
}

func (g *Grid[T]) Cols() int {
	return g.cols
}

// AddRow adds a new row of elements T to the grid
func (g *Grid[T]) AddRow(row []T) {
	// if grid is empty, set dimensions
	if g.rows == 0 && g.cols == 0 {
		// grid will always be square
		g.cols = len(row)
		g.rows = len(row)
	}
	// before adding new rows check if they match dimensions
	if len(row) != g.cols {
		panic("row has different length")
	}
	rowData := make([]T, g.cols)
	for i := range rowData {
		rowData[i] = row[i]
	}
	g.data = append(g.data, rowData)
}

// Get returns the value at the given position
func (g *Grid[T]) Get(pos Point) T {
	return g.data[pos.X][pos.Y]
}

// IsPositionValid checks if the given position is within the grid bounds
func (g *Grid[T]) IsPositionValid(position Point) bool {
	return position.X >= 0 && position.X < g.rows && position.Y >= 0 && position.Y < g.cols
}

// HasNext checks if there are more elements in the grid
func (g *Grid[T]) HasNext() bool {
	if g.currentCol >= g.cols {
		g.currentRow++
		g.currentCol = 0
	}
	if g.currentRow >= g.rows {
		return false
	}

	return true
}

// Next returns the next element T in the grid and its position (x, y)
func (g *Grid[T]) Next() (T, Point) {
	value := g.data[g.currentRow][g.currentCol]
	pos := Point{X: g.currentRow, Y: g.currentCol}
	g.currentCol++
	if g.currentCol > g.cols {
		g.currentRow++
		g.currentCol = 0
	}

	return value, pos
}

// Copy returns a copy of the grid
func (g *Grid[T]) Copy() Grid[T] {
	data := make([][]T, g.rows)
	for i := range data {
		data[i] = make([]T, g.cols)
		for j := range data[i] {
			data[i][j] = g.data[i][j]
		}
	}

	return NewGrid[T](data)
}

func (g *Grid[T]) Set(el T, pos Point) {
	g.data[pos.X][pos.Y] = el
}

func (g *Grid[T]) Compare(other Grid[T]) bool {
	for i, row := range g.data {
		for j, el := range row {
			if el != other.data[i][j] {
				return false
			}
		}
	}
	return true
}
