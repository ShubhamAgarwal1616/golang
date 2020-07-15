package simulator

import "legacy-of-brynjolf/command"

type Position struct{
	entity RoomEntity
	row int
	col int
}

func NewPostion(entity RoomEntity, row int, col int) Position {
	return Position{
		entity: entity,
		row: row,
		col: col,
	}
}

func (p Position) decrementRow() Position {
	return Position{p.entity, p.row - 1, p.col}
}

func (p Position) incrementRow() Position {
	return Position{p.entity, p.row + 1, p.col}
}

func (p Position) decrementCol() Position {
	return Position{p.entity, p.row, p.col - 1}
}

func (p Position) inclrementCol() Position {
	return Position{p.entity, p.row, p.col + 1}
}

func(p Position) update(c command.Command) Position{
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