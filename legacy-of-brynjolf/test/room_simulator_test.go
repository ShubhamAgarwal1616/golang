package test

import (
	"legacy-of-brynjolf/simulator"
	"reflect"
	"testing"
)

func TestSimulate(t *testing.T) {
	assertSimulatorStatus := func(t *testing.T, got, want simulator.RoomSimulator) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("expect brynjolf to move in upward direction till a wall comes up", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
														   {space,space,space,exit},
														   {space,brynjolf,space,space},
														   {wall,space,space,space}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Up})
		newRoom := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
														      {space,brynjolf,space,exit},
														      {space,space,space,space},
														      {wall,space,space,space}})
		want := simulator.NewRoomSimulator(newRoom)
		assertSimulatorStatus(t, roomSimulator, want)
	})

	t.Run("expect brynjolf to move in upward direction and caught by a guard", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
														   {space,guard,space,exit},
														   {space,space,space,space},
														   {wall,brynjolf,space,space}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Up})
		newRoom := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
															  {space,guard,space,exit},
															  {space,space,space,space},
															  {wall,space,space,space}})
		want := simulator.NewRoomSimulator(newRoom)
		assertSimulatorStatus(t, roomSimulator, want)
	})

	t.Run("expect brynjolf to move in upward direction and exit the room", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
														   {space,space,space,exit},
														   {space,space,space,space},
														   {wall,guard,space,brynjolf}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Up})
		newRoom := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
															  {space,guard,space,exit},
															  {space,space,space,space},
															  {wall,space,space,space}})
		want := simulator.NewRoomSimulator(newRoom)
		assertSimulatorStatus(t, roomSimulator, want)
	})
}
