package simulator

import (
	"errors"
	"strings"
)

type Command string

const (
	Up   		Command = "u"
	Down      	Command = "d"
	Right       Command = "r"
	Left       	Command = "l"
)

var InvalidCommandErr = errors.New("invalid command to execute")

func GetValidCommand(c string) (Command, error) {
	switch Command(strings.ToLower(c)) {
	case Up, Down, Left, Right:
		return Command(c), nil
	default:
		return "", InvalidCommandErr
	}
}
