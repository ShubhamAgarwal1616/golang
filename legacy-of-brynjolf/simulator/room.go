package simulator

type Room struct {
	state [][]RoomEntity
}

func NewRoom(state [][]RoomEntity) Room {
	return Room{state: state}
}

func includes(entities []RoomEntity, entity RoomEntity) bool {
	for _, e := range entities{
		if e == entity{
			return true
		}
	}
	return false
}

func (r Room) duplicteRoomState() [][]RoomEntity {
	duplicate := make([][]RoomEntity, len(r.state))
	for index := range r.state {
		duplicate[index] = make([]RoomEntity, len(r.state[index]))
		copy(duplicate[index], r.state[index])
	}
	return duplicate
}

func (r Room) moveEntities(positions []Position, command Command) Room {
	newState := r.duplicteRoomState()
	for index, position := range positions {
		blockingEntities := position.entity.GetBlockingEntities()
		newState[position.row][position.col] = EmptySpace
		for position.row > 0 && !includes(blockingEntities, r.state[position.row - 1][position.col]) {
			position = position.update()
		}
		if r.state[position.row][position.col] != Exit && !(position.entity == Brynjolf && newState[position.row][position.col] == Guard) {
			newState[position.row][position.col] = position.entity
		}
		positions[index] = position
	}
	return NewRoom(newState)
}

func (r Room) FindEntitiesPosition(entities []RoomEntity) []Position{
	var positions []Position
	for row, entitiesInRow := range r.state {
		for col, entity := range entitiesInRow {
			if includes(entities, entity) {
				if entity == Guard {
					positions = append([]Position{NewPostion(entity, row, col)}, positions...)
				}else {
					positions = append(positions, NewPostion(entity, row, col))
				}
			}
		}
	}
	return positions
}