package simulator

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

func (p Position) decrement_row() Position {
	return Position{p.entity, p.row - 1, p.col}
}

func(p Position) update() Position{
	return p.decrement_row()
}