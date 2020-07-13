package simulator

import "fmt"

type RoomSimulator struct {
	room Room
	status RoomStatus
}

func NewRoomSimulator(room Room) *RoomSimulator {
	return &RoomSimulator{room: room, status: Undecided}
}

func (rs *RoomSimulator) Simulate(commands []Command){
	fmt.Println("testing")
}