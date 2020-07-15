package simulator

var possibleCommands = []Command{Up, Down, Left, Right}
var smallestPossibleWayLenth = 0
var possibleWays []Command

func movableEntitiesNotBlocked(room Room, positions []Position, command Command) bool {
	for _, position := range positions {
		blockingEntities := position.entity.GetBlockingEntities()
		if room.notAtEdgeOrBlocked(position, blockingEntities, command){
			return true
		}
	}
	return false
}

func previousMove(previousCommands Command, comand Command) bool {
	if len(previousCommands) > 0 {
		return previousCommands[len(previousCommands)-1:] == comand

	}
	return false
}


func wonOrLost(status RoomStatus) bool {
	return status == Lost || status == Won
}

func formingPattern(previousCommands Command, comand Command) bool {
	oppositeCommand := comand.getOppositeCommand()
	length := len(previousCommands)
	if length > 1 {
		return previousCommands[length - 2:] == comand + oppositeCommand
	}
	return false
}

func necessaryMove(previousCommands Command, comand Command, status RoomStatus) bool {
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

func getRoomStatus(movableEntitiesPositions []Position, exitPosition []Position) RoomStatus {
	brynjolfPosition := movableEntitiesPositions[len(movableEntitiesPositions) - 1]
	guardsPositions := movableEntitiesPositions[0:len(movableEntitiesPositions) - 1]
	if commonPosition(guardsPositions, brynjolfPosition){
		return Lost
	}else if commonPosition(exitPosition, brynjolfPosition) {
		return Won
	}
	return Undecided
}


//assuming length of smallest possible way is less than 2 * height of room
func findPossibleWays(rs RoomSimulator, movableEntitiesPositions []Position, exitPosition []Position) []Command {
	smallestPossibleWayLenth = 2 * len(rs.room.state)
	findWays(rs.room, movableEntitiesPositions, exitPosition, Command(""), 0)
	return filterWays(possibleWays)
}

func findWays(room Room, movableEntitiesPositions []Position, exitPosition []Position, previousCommands Command, levelCount int) {
	status := getRoomStatus(movableEntitiesPositions, exitPosition)
	if status == Won && len(previousCommands) <= smallestPossibleWayLenth {
		possibleWays = append(possibleWays, previousCommands)
		smallestPossibleWayLenth = len(previousCommands)
		return
	}

	if levelCount > smallestPossibleWayLenth {return}

	for _, command := range possibleCommands {
		if necessaryMove(previousCommands, command, status) && movableEntitiesNotBlocked(room, movableEntitiesPositions, command){
			movableEntitiesPositionsCopy := makeCopy(movableEntitiesPositions)
			newRoom := room.moveEntities(movableEntitiesPositionsCopy, command)
			findWays(newRoom, movableEntitiesPositionsCopy, exitPosition, previousCommands + command, levelCount + 1)
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