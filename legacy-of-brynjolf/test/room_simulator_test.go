package test

import (
	"legacy-of-brynjolf/simulator"
	"reflect"
	"testing"
)

func TestSimulate(t *testing.T) {
	assertRoomState := func(t *testing.T, got, want simulator.Room) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertStatus := func(t *testing.T, got, want simulator.RoomStatus) {
		t.Helper()
		if got != want {
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
		want := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
														      {space,brynjolf,space,exit},
														      {space,space,space,space},
														      {wall,space,space,space}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf to move in upward direction and caught by a guard", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
														   {space,guard,space,exit},
														   {space,space,space,space},
														   {wall,brynjolf,space,space}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Up})
		want := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
															  {space,guard,space,exit},
															  {space,space,space,space},
															  {wall,space,space,space}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Lost)
	})

	t.Run("expect brynjolf to move in upward direction and exit the room", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
														   {space,space,space,exit},
														   {space,space,space,space},
														   {wall,guard,space,brynjolf}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Up})
		want := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
															  {space,guard,space,exit},
															  {space,space,space,space},
															  {wall,space,space,space}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Won)
	})

	t.Run("expect brynjolf and guards to move in left direction", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,guard},
														   {space,space,space,exit},
														   {space,space,space,space},
														   {wall,space,space,brynjolf}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Left})
		want := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,guard,space},
														   {space,space,space,exit},
														   {space,space,space,space},
														   {wall,brynjolf,space,space}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf and guards to move in right direction", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,guard},
								 						   {guard,space,space,exit},
								 						   {space,space,space,space},
								 						   {wall,brynjolf,space,space}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Right})
		want := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,guard},
														   {space,space,guard,exit},
														   {space,space,space,space},
														   {wall,space,space,brynjolf}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf and guards to move in downward direction", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,guard},
														   {guard,brynjolf,space,exit},
														   {space,space,space,space},
														   {wall,space,space,space}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Down})
		want := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,guard},
														   {space,space,space,exit},
														   {guard,space,space,space},
														   {wall,brynjolf,space,space}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf to win on execution of ru", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,guard},
														   {guard,space,space,exit},
														   {space,space,space,space},
														   {wall,brynjolf,space,space}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Right, simulator.Up})
		want := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,guard,guard},
														   {space,space,space,exit},
														   {space,space,space,space},
														   {wall,space,space,space}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Won)
	})

	t.Run("expect brynjolf to lose on execution of lu", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,space},
														   {guard,space,space,exit},
														   {space,space,space,guard},
														   {space,space,space,brynjolf}})
		roomSimulator := simulator.NewRoomSimulator(room)
		roomSimulator.Simulate([]simulator.Command{simulator.Left, simulator.Up})
		want := simulator.NewRoom([][]simulator.RoomEntity{{guard, wall,space,space},
														   {space,space,space,exit},
														   {space,space,space,space},
														   {space,space,space,space}})
		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Lost)
	})
}
