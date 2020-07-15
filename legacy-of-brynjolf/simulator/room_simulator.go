package simulator

import (
	"fmt"
	"legacy-of-brynjolf/command"
	"legacy-of-brynjolf/entities"
	"legacy-of-brynjolf/path"
	"legacy-of-brynjolf/position"
	room2 "legacy-of-brynjolf/room"
	status2 "legacy-of-brynjolf/status"
	"strconv"
)

type RoomSimulator struct {
	room   room2.Room
	status status2.RoomStatus
}

func NewRoomSimulator(room room2.Room) RoomSimulator {
	return RoomSimulator{room: room, status: status2.Undecided}
}

func Simulate(room room2.Room, commands []command.Command) {
	simulator := NewRoomSimulator(room)
	commandsExecuted, ways := simulator.Start(commands)
	mssg := string(simulator.Status()) + ": executed " + strconv.Itoa(commandsExecuted) + " commands out of " + strconv.Itoa(len(commands))
	simulator.DisplayRoom(mssg, ways)
}

func (rs *RoomSimulator) Room() room2.Room {
	return rs.room
}

func (rs *RoomSimulator) Status() status2.RoomStatus {
	return rs.status
}

func (rs *RoomSimulator) DisplayRoom(mssg string, ways []command.Command) {
	rs.room.Display(mssg)
	fmt.Println("\npossible ways to win")
	fmt.Println(ways)
}

func (rs *RoomSimulator) wonOrLost() bool {
	return rs.status == status2.Won || rs.status == status2.Lost
}

func (rs *RoomSimulator) executeCommand(commands []command.Command, commandExecuted int, movableEntitiesPositions []position.Position, exitPosition []position.Position) int {
	for index, cmd := range commands {
		commandExecuted = index + 1
		rs.room = rs.room.MoveEntities(movableEntitiesPositions, cmd)
		rs.status = path.GetRoomStatus(movableEntitiesPositions, exitPosition)
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
		ways = path.FindPossibleWays(rs.room, movableEntitiesPositions, exitPosition)
	}
	return commandsExecuted, ways
}