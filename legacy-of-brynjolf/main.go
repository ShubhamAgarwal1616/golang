package main

import (
	"io/ioutil"
	simulator "legacy-of-brynjolf/simulator"
	"log"
	"os"
)

const inputFile = "room.txt"

func readInput() string {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error while reading file room.txt")
	}
	return string(data)
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
	room, err := simulator.NewRoom(data)
	if err != nil {
		log.Fatal(err)
	}
	commands := []simulator.Command{}
	if len(os.Args) > 1 {
		commands = buildCommands(os.Args[1])
	}
	simulator.Simulate(room, commands)
}


