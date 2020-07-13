package simulator

import "errors"

type Command string

const (
	Up   		Command = "u"
	Down      	Command = "d"
	Right       Command = "r"
	Left       	Command = "l"
)

var InvalidCommandErr = errors.New("invalid command to execute")

func GetValidCommand(c string) (Command, error) {
	switch Command(c) {
	case Up, Down, Left, Right:
		return Command(c), nil
	default:
		return "", InvalidCommandErr
	}
}
