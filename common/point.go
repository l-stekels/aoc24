package common

type Point struct {
	X, Y int
}

var (
	Directions = []Direction{Up, Right, Down, Left}
	Up         = direction{name: "Up", p: Point{X: -1}}
	Right      = direction{name: "Right", p: Point{Y: 1}}
	Down       = direction{name: "Down", p: Point{X: 1}}
	Left       = direction{name: "Left", p: Point{Y: -1}}
)

type Direction interface {
	X() int
	Y() int
}

type direction struct {
	name string
	p    Point
}

func (d direction) String() string {
	return d.name
}

func (d direction) X() int {
	return d.p.X
}

func (d direction) Y() int {
	return d.p.Y
}
