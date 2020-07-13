package simulator

import "strconv"

type simulator interface {
	simulate(commands []Command)
}

type RoomSimulator struct {
	room Room
	status RoomStatus
}

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

func (rs *RoomSimulator) updateStatus(movableEntitiesPositions []Position, exitPosition []Position) {
	brynjolfPosition := movableEntitiesPositions[len(movableEntitiesPositions) - 1]
	guardsPositions := movableEntitiesPositions[0:len(movableEntitiesPositions) - 1]
	if commonPosition(guardsPositions, brynjolfPosition){
		rs.status = Lost
	}else if commonPosition(exitPosition, brynjolfPosition) {
		rs.status = Won
	}

}

func (rs *RoomSimulator) Simulate(commands []Command) {
	movableEntitiesPositions := rs.room.FindEntitiesPosition([]RoomEntity{Guard, Brynjolf})
	exitPosition := rs.room.FindEntitiesPosition([]RoomEntity{Exit})
	var commandExecuted int
	for index, command := range commands {
		commandExecuted = index + 1
		rs.room = rs.room.moveEntities(movableEntitiesPositions, command)
		rs.updateStatus(movableEntitiesPositions, exitPosition)
		if rs.wonOrLost(){
			break
		}
	}
	rs.room.display(string(rs.status) + ": executed " + strconv.Itoa(commandExecuted) + " commands out of " + strconv.Itoa(len(commands)))
}