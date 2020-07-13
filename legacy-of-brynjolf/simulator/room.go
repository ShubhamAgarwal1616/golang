package simulator

type Room struct {
	state [][]RoomEntity
}

func NewRoom(state [][]RoomEntity) Room {
	return Room{state: state}
}