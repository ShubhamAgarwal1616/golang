package simulator

import (
	"errors"
)

type RoomEntity string

const (
	Brynjolf   RoomEntity = "b"
	Guard      RoomEntity = "g"
	Wall       RoomEntity = "x"
	Exit       RoomEntity = "e"
	EmptySpace RoomEntity = "0"
)

var InvalidEntityErr = errors.New("invalid Entity in input file")

func ConvertToRoomEntity(e string) (RoomEntity, error) {
	switch RoomEntity(e) {
	case Brynjolf, Guard, Wall, EmptySpace, Exit:
		return RoomEntity(e), nil
	default:
		return "", InvalidEntityErr
	}
}