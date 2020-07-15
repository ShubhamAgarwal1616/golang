package status

type RoomStatus string

const (
	Undecided RoomStatus = "undecided"
	Won       RoomStatus = "won"
	Lost      RoomStatus = "lost"
)
