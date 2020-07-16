package position

import (
	"legacy-of-brynjolf/command"
)

type Position struct{
	row    int
	col    int
}

func NewPostion(row int, col int) Position {
	return Position{
		row: row,
		col: col,
	}
}

func (p Position) Row() int { return p.row }

func (p Position) Col() int {return p.col }


func (p Position) decrementRow() Position {
	return Position{p.row - 1, p.col}
}

func (p Position) incrementRow() Position {
	return Position{p.row + 1, p.col}
}

func (p Position) decrementCol() Position {
	return Position{p.row, p.col - 1}
}

func (p Position) inclrementCol() Position {
	return Position{p.row, p.col + 1}
}

func(p Position) Update(c command.Command) Position {
	switch c {
	case command.Up:
		return p.decrementRow()
	case command.Down:
		return p.incrementRow()
	case command.Left:
		return p.decrementCol()
	default:
		return p.inclrementCol()
	}
}