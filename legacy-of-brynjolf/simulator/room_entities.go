package simulator

import (
	"errors"
	"strings"
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
	switch RoomEntity(strings.ToLower(e)) {
	case Brynjolf, Guard, Wall, EmptySpace, Exit:
		return RoomEntity(e), nil
	default:
		return "", InvalidEntityErr
	}
}

func (e RoomEntity) GetBlockingEntities() []RoomEntity {
	if e == Guard {
		return []RoomEntity{Wall, Exit}
	}else {
		return []RoomEntity{Wall}
	}
}
