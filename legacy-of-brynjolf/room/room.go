package room

import (
	"fmt"
	"legacy-of-brynjolf/command"
	"legacy-of-brynjolf/position"
	"legacy-of-brynjolf/room/entities"
	"strings"
)

type Room struct {
	state [][]entities.Entity
}

func buildRow(cols []string) ([]entities.Entity, error) {
	var entityRow []entities.Entity
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
	var state [][]entities.Entity
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

func includes(entities []entities.Entity, entity entities.Entity) bool {
	for _, e := range entities{
		if e == entity{
			return true
		}
	}
	return false
}

func (r Room) Size() int {
	return len(r.state)
}

func (r Room) duplicteRoomState() [][]entities.Entity {
	duplicate := make([][]entities.Entity, len(r.state))
	for index := range r.state {
		duplicate[index] = make([]entities.Entity, len(r.state[index]))
		copy(duplicate[index], r.state[index])
	}
	return duplicate
}

func (r Room) MoveEntities(positions []position.Position, command command.Command) Room {
	newState := r.duplicteRoomState()
	for index, pos := range positions {
		blockingEntities := pos.Entity().GetBlockingEntities()
		newState[pos.Row()][pos.Col()] = entities.EmptySpace
		pos = r.moveEntity(pos, blockingEntities, newState, command)
		r.updateNewState(pos, newState)
		positions[index] = pos
	}
	return Room{newState}
}

func (r Room) updateNewState(pos position.Position, newState [][]entities.Entity) {
	if r.state[pos.Row()][pos.Col()] != entities.Exit && !(pos.Entity() == entities.Brynjolf && newState[pos.Row()][pos.Col()] == entities.Guard) {
		newState[pos.Row()][pos.Col()] = pos.Entity()
	}
}

func (r Room) moveEntity(pos position.Position, blockingEntities []entities.Entity, newState [][]entities.Entity, command command.Command) position.Position {
	for r.NotAtEdgeOrBlocked(pos, blockingEntities, command) {
		if pos.Entity() == entities.Brynjolf && (newState[pos.Row()][pos.Col()] == entities.Guard || newState[pos.Row()][pos.Col()] == entities.Exit) {
			break
		}
		pos = pos.Update(command)
	}
	return pos
}

func (r Room) NotAtEdgeOrBlocked(pos position.Position, blockingEntities []entities.Entity, c command.Command) bool {
	switch c {
	case command.Up:
		return pos.Row() > 0 && !includes(blockingEntities, r.state[pos.Row() - 1][pos.Col()])
	case command.Down:
		return pos.Row() < len(r.state) - 1 && !includes(blockingEntities, r.state[pos.Row() + 1][pos.Col()])
	case command.Left:
		return pos.Col() > 0 && !includes(blockingEntities, r.state[pos.Row()][pos.Col() - 1])
	case command.Right:
		return pos.Col() < len(r.state[0]) - 1 && !includes(blockingEntities, r.state[pos.Row()][pos.Col() + 1])
	}
	return false
}

func (r Room) FindEntitiesPosition(e []entities.Entity) []position.Position {
	var positions []position.Position
	for row, entitiesInRow := range r.state {
		for col, entity := range entitiesInRow {
			if includes(e, entity) {
				if entity == entities.Guard {
					positions = append([]position.Position{position.NewPostion(entity, row, col)}, positions...)
				}else {
					positions = append(positions, position.NewPostion(entity, row, col))
				}
			}
		}
	}
	return positions
}

func (r Room) Display(mssg string) {
	fmt.Println(mssg)
	for _, row := range r.state {
		for _, entity := range row {
			fmt.Print(entity + ",")
		}
		fmt.Print("\n")
	}
}