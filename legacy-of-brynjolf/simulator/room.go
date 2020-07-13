package simulator

import "fmt"

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
		position = r.moveEntity(position, blockingEntities, newState, command)
		r.updateNewState(position, newState)
		positions[index] = position
	}
	return NewRoom(newState)
}

func (r Room) updateNewState(position Position, newState [][]RoomEntity) {
	if r.state[position.row][position.col] != Exit && !(position.entity == Brynjolf && newState[position.row][position.col] == Guard) {
		newState[position.row][position.col] = position.entity
	}
}

func (r Room) moveEntity(position Position, blockingEntities []RoomEntity, newState [][]RoomEntity, command Command) Position {
	for r.notAtEdgeOrBlocked(position, blockingEntities, command) {
		if position.entity == Brynjolf && (newState[position.row][position.col] == Guard || newState[position.row][position.col] == Exit) {
			break
		}
		position = position.update(command)
	}
	return position
}

func (r Room) notAtEdgeOrBlocked(position Position, blockingEntities []RoomEntity, command Command) bool {
	switch command {
	case Up:
		return position.row > 0 && !includes(blockingEntities, r.state[position.row - 1][position.col])
	case Down:
		return position.row < len(r.state) - 1 && !includes(blockingEntities, r.state[position.row + 1][position.col])
	case Left:
		return position.col > 0 && !includes(blockingEntities, r.state[position.row][position.col - 1])
	case Right:
		return position.col < len(r.state[0]) - 1 && !includes(blockingEntities, r.state[position.row][position.col + 1])
	}
	return false
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

func (r Room) display(mssg string) {
	fmt.Println(mssg)
	for _, row := range r.state {
		for _, entity := range row {
			fmt.Print(entity + ",")
		}
		fmt.Print("\n")
	}
}