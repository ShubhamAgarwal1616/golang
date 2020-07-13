package main

import (
	"io/ioutil"
	simulator "legacy-of-brynjolf/simulator"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "room.txt"

func readInput() string {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error while reading file room.txt")
	}
	return string(data)
}

func checkEntity(cols []string) []simulator.RoomEntity {
	var entities []simulator.RoomEntity
	for _, entity := range cols {
		entity, err := simulator.ConvertToRoomEntity(entity)
		if err != nil {
			log.Fatal(err)
		}
		entities = append(entities, entity)
	}
	return entities
}

func buildRoomState(data string) [][]simulator.RoomEntity {
	var state [][]simulator.RoomEntity
	rows := strings.Split(strings.TrimSpace(data), "\n")
	for _, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), ",")
		entities := checkEntity(cols)
		state = append(state, entities)
	}
	return state
}

func buildCommands(data string) []simulator.Command {
	var commands []simulator.Command
	for _, char := range data {
		command, err := simulator.GetValidCommand(string(char))
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, command)
	}
	return commands
}

func main() {
	data := readInput()
	room := simulator.NewRoom(buildRoomState(data))
	commands := []simulator.Command{}
	if len(os.Args) > 1 {
		commands = buildCommands(os.Args[1])
	}
	simulator := simulator.NewRoomSimulator(room)
	commandsExecuted := simulator.Simulate(commands)
	simulator.DisplayRoom(string(simulator.Status()) + ": executed " + strconv.Itoa(commandsExecuted) + " commands out of " + strconv.Itoa(len(commands)))
}


