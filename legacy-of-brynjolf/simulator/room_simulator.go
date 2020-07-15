package simulator

import (
	"fmt"
	"strconv"
)

type simulator interface {
	simulate(commands []Command)
}

type RoomSimulator struct {
	room Room
	status RoomStatus
}

func NewRoomSimulator(room Room) RoomSimulator {
	return RoomSimulator{room: room, status: Undecided}
}

func Simulate(room Room, commands []Command) {
	simulator := NewRoomSimulator(room)
	commandsExecuted, ways := simulator.Start(commands)
	mssg := string(simulator.Status()) + ": executed " + strconv.Itoa(commandsExecuted) + " commands out of " + strconv.Itoa(len(commands))
	simulator.DisplayRoom(mssg, ways)
}

func (rs *RoomSimulator) Room() Room {
	return rs.room
}

func (rs *RoomSimulator) Status() RoomStatus {
	return rs.status
}

func (rs *RoomSimulator) DisplayRoom(mssg string, ways []Command) {
	rs.room.display(mssg)
	fmt.Println("\npossible ways to win")
	fmt.Println(ways)
}

func (rs *RoomSimulator) wonOrLost() bool {
	return rs.status == Won || rs.status == Lost
}

func (rs *RoomSimulator) executeCommand(commands []Command, commandExecuted int, movableEntitiesPositions []Position, exitPosition []Position) int {
	for index, command := range commands {
		commandExecuted = index + 1
		rs.room = rs.room.moveEntities(movableEntitiesPositions, command)
		rs.status = getRoomStatus(movableEntitiesPositions, exitPosition)
		if rs.wonOrLost() {
			break
		}
	}
	return commandExecuted
}

func (rs *RoomSimulator) Start(commands []Command) (int, []Command){
	movableEntitiesPositions := rs.room.FindEntitiesPosition([]RoomEntity{Guard, Brynjolf})
	exitPosition := rs.room.FindEntitiesPosition([]RoomEntity{Exit})
	var commandsExecuted int
	commandsExecuted = rs.executeCommand(commands, commandsExecuted, movableEntitiesPositions, exitPosition)
	var ways []Command
	if !rs.wonOrLost() {
		ways = findPossibleWays(*rs, movableEntitiesPositions, exitPosition)
	}
	return commandsExecuted, ways
}