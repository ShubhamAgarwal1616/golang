package simulator

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

func (rs *RoomSimulator) Simulate(commands []Command){
	movableEntitiesPositions := rs.room.FindEntitiesPosition([]RoomEntity{Guard, Brynjolf})
	for _, command := range commands {
		rs.room = rs.room.moveEntities(movableEntitiesPositions, command)
	}
}