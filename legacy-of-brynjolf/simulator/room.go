package simulator

import (
	"fmt"
	"legacy-of-brynjolf/command"
	"legacy-of-brynjolf/entities"
	"strings"
)

type Room struct {
	state [][]entities.RoomEntity
}

func buildRow(cols []string) ([]entities.RoomEntity, error) {
	var entityRow []entities.RoomEntity
	for _, entity := range cols {
		entity, err := entities.BuildEntity(entity)
		if err != nil {
			return nil, err
		}
		entityRow = append(entityRow, entity)
	}
	return entityRow, nil
}

func NewRoom(data string) (Room, error) {
	var state [][]entities.RoomEntity
	rows := strings.Split(strings.TrimSpace(data), "\n")
	for _, row := range rows {
		entityRow, err := buildRow(strings.Split(strings.TrimSpace(row), ","))
		if err != nil{
			return Room{}, err
		}
		state = append(state, entityRow)
	}
	return Room{state: state}, nil
}

func includes(entities []entities.RoomEntity, entity entities.RoomEntity) bool {
	for _, e := range entities{
		if e == entity{
			return true
		}
	}
	return false
}

func (r Room) duplicteRoomState() [][]entities.RoomEntity {
	duplicate := make([][]entities.RoomEntity, len(r.state))
	for index := range r.state {
		duplicate[index] = make([]entities.RoomEntity, len(r.state[index]))
		copy(duplicate[index], r.state[index])
	}
	return duplicate
}

func (r Room) moveEntities(positions []Position, command command.Command) Room {
	newState := r.duplicteRoomState()
	for index, position := range positions {
		blockingEntities := position.entity.GetBlockingEntities()
		newState[position.row][position.col] = entities.EmptySpace
		position = r.moveEntity(position, blockingEntities, newState, command)
		r.updateNewState(position, newState)
		positions[index] = position
	}
	return Room{newState}
}

func (r Room) updateNewState(position Position, newState [][]entities.RoomEntity) {
	if r.state[position.row][position.col] != entities.Exit && !(position.entity == entities.Brynjolf && newState[position.row][position.col] == entities.Guard) {
		newState[position.row][position.col] = position.entity
	}
}

func (r Room) moveEntity(position Position, blockingEntities []entities.RoomEntity, newState [][]entities.RoomEntity, command command.Command) Position {
	for r.notAtEdgeOrBlocked(position, blockingEntities, command) {
		if position.entity == entities.Brynjolf && (newState[position.row][position.col] == entities.Guard || newState[position.row][position.col] == entities.Exit) {
			break
		}
		position = position.update(command)
	}
	return position
}

func (r Room) notAtEdgeOrBlocked(position Position, blockingEntities []entities.RoomEntity, c command.Command) bool {
	switch c {
	case command.Up:
		return position.row > 0 && !includes(blockingEntities, r.state[position.row - 1][position.col])
	case command.Down:
		return position.row < len(r.state) - 1 && !includes(blockingEntities, r.state[position.row + 1][position.col])
	case command.Left:
		return position.col > 0 && !includes(blockingEntities, r.state[position.row][position.col - 1])
	case command.Right:
		return position.col < len(r.state[0]) - 1 && !includes(blockingEntities, r.state[position.row][position.col + 1])
	}
	return false
}

func (r Room) FindEntitiesPosition(e []entities.RoomEntity) []Position{
	var positions []Position
	for row, entitiesInRow := range r.state {
		for col, entity := range entitiesInRow {
			if includes(e, entity) {
				if entity == entities.Guard {
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