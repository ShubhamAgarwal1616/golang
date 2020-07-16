package path

import (
	"legacy-of-brynjolf/command"
	_blocks "legacy-of-brynjolf/room/blocks"
	"legacy-of-brynjolf/status"
)

var possibleCommands = []command.Command{command.Up, command.Down, command.Left, command.Right}
var smallestPossibleWayLenth = 0
var possibleWays []command.Command

//func movableEntitiesNotBlocked(room _room.Room, positions []position.Position, command command.Command) bool {
//	for _, pos := range positions {
//		blockingEntities := pos.Entity().GetBlockingEntities()
//		if room.NotAtEdgeOrBlocked(pos, blockingEntities, command){
//			return true
//		}
//	}
//	return false
//}
//
//func previousMove(previousCommands command.Command, comand command.Command) bool {
//	if len(previousCommands) > 0 {
//		return previousCommands[len(previousCommands)-1:] == comand
//
//	}
//	return false
//}
//
//
//func wonOrLost(s status.RoomStatus) bool {
//	return s == status.Lost || s == status.Won
//}
//
//func formingPattern(previousCommands command.Command, comand command.Command) bool {
//	oppositeCommand := comand.OppositeCommand()
//	length := len(previousCommands)
//	if length > 1 {
//		return previousCommands[length - 2:] == comand + oppositeCommand
//	}
//	return false
//}
//
//func necessaryMove(previousCommands command.Command, comand command.Command, status status.RoomStatus) bool {
//	return (!previousMove(previousCommands, comand)) && (!wonOrLost(status)) && (!formingPattern(previousCommands, comand))
//}

func commonPosition(blocks []_blocks.Block, brynjolfBlock _blocks.Block) bool {
	for _, block := range blocks {
		if block.SamePosition(brynjolfBlock){
			return true
		}
	}
	return false
}

func GetRoomStatus(movableEntitiesBlocks []_blocks.Block, exitBlock []_blocks.Block) status.RoomStatus {
	brynjolfBlock := movableEntitiesBlocks[len(movableEntitiesBlocks) - 1]
	guardBlocks := movableEntitiesBlocks[0:len(movableEntitiesBlocks) - 1]
	if commonPosition(guardBlocks, brynjolfBlock){
		return status.Lost
	}else if commonPosition(exitBlock, brynjolfBlock) {
		return status.Won
	}
	return status.Undecided
}


//assuming length of smallest possible way is less than 2 * height of room
//func FindPossibleWays(room _room.Room, movableEntitiesPositions []position.Position, exitPosition []position.Position) []command.Command {
//	smallestPossibleWayLenth = 2 * room.Size()
//	findWays(room, movableEntitiesPositions, exitPosition, "", 0)
//	return filterWays(possibleWays)
//}
//
//func findWays(room _room.Room, movableEntitiesPositions []position.Position, exitPosition []position.Position, previousCommands command.Command, levelCount int) {
//	newStatus := GetRoomStatus(movableEntitiesPositions, exitPosition)
//	if newStatus == status.Won && len(previousCommands) <= smallestPossibleWayLenth {
//		possibleWays = append(possibleWays, previousCommands)
//		smallestPossibleWayLenth = len(previousCommands)
//		return
//	}
//
//	if levelCount > smallestPossibleWayLenth {return}
//
//	for _, cmd := range possibleCommands {
//		if necessaryMove(previousCommands, cmd, newStatus) && movableEntitiesNotBlocked(room, movableEntitiesPositions, cmd){
//			movableEntitiesPositionsCopy := makeCopy(movableEntitiesPositions)
//			newRoom := room.MoveEntities(movableEntitiesPositionsCopy, cmd)
//			findWays(newRoom, movableEntitiesPositionsCopy, exitPosition, previousCommands +cmd, levelCount + 1)
//		}
//	}
//}
//
//func makeCopy(positions []position.Position) []position.Position {
//	duplicate := make([]position.Position, len(positions))
//	copy(duplicate, positions)
//	return duplicate
//}

func filterWays(possibleWays []command.Command) []command.Command {
	var filteredWays []command.Command
	for _, way := range possibleWays {
		if len(way) == smallestPossibleWayLenth {
			filteredWays = append(filteredWays, way)
		}
	}
	return filteredWays
}