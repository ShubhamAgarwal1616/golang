package simulator

import (
	"legacy-of-brynjolf/command"
	"legacy-of-brynjolf/status"
)

var possibleCommands = []command.Command{command.Up, command.Down, command.Left, command.Right}
var smallestPossibleWayLenth = 0
var possibleWays []command.Command

func movableEntitiesNotBlocked(room Room, positions []Position, command command.Command) bool {
	for _, position := range positions {
		blockingEntities := position.entity.GetBlockingEntities()
		if room.notAtEdgeOrBlocked(position, blockingEntities, command){
			return true
		}
	}
	return false
}

func previousMove(previousCommands command.Command, comand command.Command) bool {
	if len(previousCommands) > 0 {
		return previousCommands[len(previousCommands)-1:] == comand

	}
	return false
}


func wonOrLost(s status.RoomStatus) bool {
	return s == status.Lost || s == status.Won
}

func formingPattern(previousCommands command.Command, comand command.Command) bool {
	oppositeCommand := comand.OppositeCommand()
	length := len(previousCommands)
	if length > 1 {
		return previousCommands[length - 2:] == comand + oppositeCommand
	}
	return false
}

func necessaryMove(previousCommands command.Command, comand command.Command, status status.RoomStatus) bool {
	return (!previousMove(previousCommands, comand)) && (!wonOrLost(status)) && (!formingPattern(previousCommands, comand))
}

func commonPosition(positions []Position, brynjolfPosition Position) bool {
	for _, position := range positions {
		if position.row == brynjolfPosition.row && position.col == brynjolfPosition.col{
			return true
		}
	}
	return false
}

func getRoomStatus(movableEntitiesPositions []Position, exitPosition []Position) status.RoomStatus {
	brynjolfPosition := movableEntitiesPositions[len(movableEntitiesPositions) - 1]
	guardsPositions := movableEntitiesPositions[0:len(movableEntitiesPositions) - 1]
	if commonPosition(guardsPositions, brynjolfPosition){
		return status.Lost
	}else if commonPosition(exitPosition, brynjolfPosition) {
		return status.Won
	}
	return status.Undecided
}


//assuming length of smallest possible way is less than 2 * height of room
func findPossibleWays(rs RoomSimulator, movableEntitiesPositions []Position, exitPosition []Position) []command.Command {
	smallestPossibleWayLenth = 2 * len(rs.room.state)
	findWays(rs.room, movableEntitiesPositions, exitPosition, "", 0)
	return filterWays(possibleWays)
}

func findWays(room Room, movableEntitiesPositions []Position, exitPosition []Position, previousCommands command.Command, levelCount int) {
	newStatus := getRoomStatus(movableEntitiesPositions, exitPosition)
	if newStatus == status.Won && len(previousCommands) <= smallestPossibleWayLenth {
		possibleWays = append(possibleWays, previousCommands)
		smallestPossibleWayLenth = len(previousCommands)
		return
	}

	if levelCount > smallestPossibleWayLenth {return}

	for _, cmd := range possibleCommands {
		if necessaryMove(previousCommands, cmd, newStatus) && movableEntitiesNotBlocked(room, movableEntitiesPositions, cmd){
			movableEntitiesPositionsCopy := makeCopy(movableEntitiesPositions)
			newRoom := room.moveEntities(movableEntitiesPositionsCopy, cmd)
			findWays(newRoom, movableEntitiesPositionsCopy, exitPosition, previousCommands +cmd, levelCount + 1)
		}
	}
}

func makeCopy(positions []Position) []Position {
	duplicate := make([]Position, len(positions))
	copy(duplicate, positions)
	return duplicate
}

func filterWays(possibleWays []command.Command) []command.Command {
	var filteredWays []command.Command
	for _, way := range possibleWays {
		if len(way) == smallestPossibleWayLenth {
			filteredWays = append(filteredWays, way)
		}
	}
	return filteredWays
}