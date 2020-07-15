package command

import (
	"errors"
	"strings"
)

type Command string

const (
	Up    Command = "u"
	Down  Command = "d"
	Right Command = "r"
	Left  Command = "l"
)

var InvalidCommandErr = errors.New("invalid command to execute")

func getValidCommand(c string) (Command, error) {
	switch Command(strings.ToLower(c)) {
	case Up, Down, Left, Right:
		return Command(c), nil
	default:
		return "", InvalidCommandErr
	}
}

func Build(data string) ([]Command, error) {
	var commands []Command
	for _, char := range data {
		command, err := getValidCommand(string(char))
		if err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func (c Command) OppositeCommand() Command {
	switch c {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	default:
		return Left
	}
}
