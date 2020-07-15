package simulator

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