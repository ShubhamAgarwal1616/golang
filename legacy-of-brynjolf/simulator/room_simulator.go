package simulator

import (
	"fmt"
	"legacy-of-brynjolf/command"
	"legacy-of-brynjolf/entities"
	status2 "legacy-of-brynjolf/status"
	"strconv"
)

type RoomSimulator struct {
	room   Room
	status status2.RoomStatus
}

func NewRoomSimulator(room Room) RoomSimulator {
	return RoomSimulator{room: room, status: status2.Undecided}
}

func Simulate(room Room, commands []command.Command) {
	simulator := NewRoomSimulator(room)
	commandsExecuted, ways := simulator.Start(commands)
	mssg := string(simulator.Status()) + ": executed " + strconv.Itoa(commandsExecuted) + " commands out of " + strconv.Itoa(len(commands))
	simulator.DisplayRoom(mssg, ways)
}

func (rs *RoomSimulator) Room() Room {
	return rs.room
}

func (rs *RoomSimulator) Status() status2.RoomStatus {
	return rs.status
}

func (rs *RoomSimulator) DisplayRoom(mssg string, ways []command.Command) {
	rs.room.display(mssg)
	fmt.Println("\npossible ways to win")
	fmt.Println(ways)
}

func (rs *RoomSimulator) wonOrLost() bool {
	return rs.status == status2.Won || rs.status == status2.Lost
}

func (rs *RoomSimulator) executeCommand(commands []command.Command, commandExecuted int, movableEntitiesPositions []Position, exitPosition []Position) int {
	for index, cmd := range commands {
		commandExecuted = index + 1
		rs.room = rs.room.moveEntities(movableEntitiesPositions, cmd)
		rs.status = getRoomStatus(movableEntitiesPositions, exitPosition)
		if rs.wonOrLost() {
			break
		}
	}
	return commandExecuted
}

func (rs *RoomSimulator) Start(commands []command.Command) (int, []command.Command){
	movableEntitiesPositions := rs.room.FindEntitiesPosition([]entities.RoomEntity{entities.Guard, entities.Brynjolf})
	exitPosition := rs.room.FindEntitiesPosition([]entities.RoomEntity{entities.Exit})
	var commandsExecuted int
	commandsExecuted = rs.executeCommand(commands, commandsExecuted, movableEntitiesPositions, exitPosition)
	var ways []command.Command
	if !rs.wonOrLost() {
		ways = findPossibleWays(*rs, movableEntitiesPositions, exitPosition)
	}
	return commandsExecuted, ways
}