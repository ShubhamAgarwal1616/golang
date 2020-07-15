package entities

import (
	"errors"
	"strings"
)

type Entity string

const (
	Brynjolf   Entity = "b"
	Guard      Entity = "g"
	Wall       Entity = "x"
	Exit       Entity = "e"
	EmptySpace Entity = "0"
)

var InvalidEntityErr = errors.New("invalid Entity in input file")

func BuildEntity(e string) (Entity, error) {
	switch Entity(strings.ToLower(e)) {
	case Brynjolf, Guard, Wall, EmptySpace, Exit:
		return Entity(e), nil
	default:
		return "", InvalidEntityErr
	}
}

func (e Entity) GetBlockingEntities() []Entity {
	if e == Guard {
		return []Entity{Wall, Exit}
	}else {
		return []Entity{Wall}
	}
}
