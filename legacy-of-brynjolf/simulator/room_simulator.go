package simulator

import (
	"fmt"
)

type simulator interface {
	simulate(commands []Command)
}

type RoomSimulator struct {
	room Room
	status RoomStatus
}

var possibleCommands = []Command{Up, Down, Left, Right}
var smallestPossibleWayLenth = 0
var possibleWays []Command

func commonPosition(positions []Position, brynjolfPosition Position) bool {
	for _, position := range positions {
		if position.row == brynjolfPosition.row && position.col == brynjolfPosition.col{
			return true
		}
	}
	return false
}

func NewRoomSimulator(room Room) RoomSimulator {
	return RoomSimulator{room: room, status: Undecided}
}

func (rs *RoomSimulator) Room() Room {
	return rs.room
}

func (rs *RoomSimulator) Status() RoomStatus {
	return rs.status
}

func (rs *RoomSimulator) wonOrLost() bool {
	return rs.status == Won || rs.status == Lost
}

func (rs *RoomSimulator) getupdatedStatus(movableEntitiesPositions []Position, exitPosition []Position) RoomStatus {
	brynjolfPosition := movableEntitiesPositions[len(movableEntitiesPositions) - 1]
	guardsPositions := movableEntitiesPositions[0:len(movableEntitiesPositions) - 1]
	if commonPosition(guardsPositions, brynjolfPosition){
		return Lost
	}else if commonPosition(exitPosition, brynjolfPosition) {
		return Won
	}
	return Undecided
}

func (rs *RoomSimulator) executeCommand(commands []Command, commandExecuted int, movableEntitiesPositions []Position, exitPosition []Position) int {
	for index, command := range commands {
		commandExecuted = index + 1
		rs.room = rs.room.moveEntities(movableEntitiesPositions, command)
		rs.status = rs.getupdatedStatus(movableEntitiesPositions, exitPosition)
		if rs.wonOrLost() {
			break
		}
	}
	return commandExecuted
}

func (rs *RoomSimulator) DisplayRoom(mssg string) {
	rs.room.display(mssg)
	fmt.Println("\npossible ways to win")
	fmt.Println(possibleWays)
}

func (rs *RoomSimulator) Simulate(commands []Command) int{
	movableEntitiesPositions := rs.room.FindEntitiesPosition([]RoomEntity{Guard, Brynjolf})
	exitPosition := rs.room.FindEntitiesPosition([]RoomEntity{Exit})
	var commandExecuted int
	commandExecuted = rs.executeCommand(commands, commandExecuted, movableEntitiesPositions, exitPosition)
	if !rs.wonOrLost() {
		rs.findPossibleWays(movableEntitiesPositions, exitPosition)
	}
	return commandExecuted
}

//assuming length of smallest possible way is less than 2 * height of room
func (rs *RoomSimulator) findPossibleWays(movableEntitiesPositions []Position, exitPosition []Position) {
	smallestPossibleWayLenth = 2 * len(rs.room.state)
	rs.findWays(rs.room, movableEntitiesPositions, exitPosition, Command(""), 0)
	possibleWays = filterWays(possibleWays)
}

func (rs *RoomSimulator) findWays(room Room, movableEntitiesPositions []Position, exitPosition []Position, previousCommands Command, levelCount int) {
	status := rs.getupdatedStatus(movableEntitiesPositions, exitPosition)
	if status == Won && len(previousCommands) <= smallestPossibleWayLenth{
		possibleWays = append(possibleWays, previousCommands)
		smallestPossibleWayLenth = len(previousCommands)
		return
	}

	if levelCount > smallestPossibleWayLenth {return}

	for _, command := range possibleCommands {
		if necessaryMove(previousCommands, command, status) && movableEntitiesNotBlocked(room, movableEntitiesPositions, command){
			movableEntitiesPositionsCopy := makeCopy(movableEntitiesPositions)
			newRoom := room.moveEntities(movableEntitiesPositionsCopy, command)
			rs.findWays(newRoom, movableEntitiesPositionsCopy, exitPosition, previousCommands + command, levelCount + 1)
		}
	}
}

func makeCopy(positions []Position) []Position {
	duplicate := make([]Position, len(positions))
	copy(duplicate, positions)
	return duplicate
}

func filterWays(possibleWays []Command) []Command {
	filteredWays := []Command{}
	for _, way := range possibleWays {
		if len(way) == smallestPossibleWayLenth {
			filteredWays = append(filteredWays, way)
		}
	}
	return filteredWays
}