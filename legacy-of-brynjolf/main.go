package main

import (
	"io/ioutil"
	"legacy-of-brynjolf/command"
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

func main() {
	data := readInput()
	room, err := simulator.NewRoom(data)
	if err != nil {
		log.Fatal(err)
	}
	var commands []command.Command
	if len(os.Args) > 1{
		commands, err = command.Build(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	simulator.Simulate(room, commands)
}


