package model

const (
	WHITE   = 0
	BLACK   = 1
	VISITED = 2
	N       = -1
)

const SIZE = 10

type Position struct {
	I     int
	J     int
	Color int
}

func NewPosition(i, j, color int) Position {
	return Position{I: i, J: j, Color: color}
}

func (pos Position) Equals(position Position) bool {
	return pos.I == position.I && pos.J == position.J && pos.Color == position.Color
}
