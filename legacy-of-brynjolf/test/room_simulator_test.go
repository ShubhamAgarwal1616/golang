package test

import (
	"legacy-of-brynjolf/command"
	_room "legacy-of-brynjolf/room"
	"legacy-of-brynjolf/simulator"
	"legacy-of-brynjolf/status"
	"reflect"
	"testing"
)

func TestSimulate(t *testing.T) {
	assertRoomState := func(t *testing.T, got, want _room.Room) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertStatus := func(t *testing.T, got, want status.RoomStatus) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("expect brynjolf to move in upward direction till a wall comes up", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,b,0,0\nx,0,0,0"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,x\n0,b,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Undecided)
	})

	t.Run("expect brynjolf to move in upward direction and caught by a guard", func(t *testing.T) {
		data := "0,x,0,x\n0,g,0,e\n0,0,0,0\nx,b,0,0"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,x\n0,g,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Lost)
	})

	t.Run("expect brynjolf to move in upward direction and exit the room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,0,0,0\nx,g,0,b"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,x\n0,g,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Won)
	})

	t.Run("expect brynjolf and guards to move in left direction", func(t *testing.T) {
		data := "0,x,0,g\n0,0,0,e\n0,0,0,0\nx,0,0,b"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,g,0\n0,0,0,e\n0,0,0,0\nx,b,0,0"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Left})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Undecided)
	})

	t.Run("expect brynjolf and guards to move in right direction", func(t *testing.T) {
		data := "0,x,0,g\ng,0,0,e\n0,0,0,0\nx,b,0,0"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,g\n0,0,g,e\n0,0,0,0\nx,0,0,b"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Right})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Undecided)
	})

	t.Run("expect brynjolf and guards to move in downward direction", func(t *testing.T) {
		data := "0,x,0,g\ng,b,0,e\n0,0,0,0\nx,0,0,0"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,g\n0,0,0,e\ng,0,0,0\nx,b,0,0"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Down})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Undecided)
	})

	t.Run("expect brynjolf to win on execution of ru", func(t *testing.T) {
		data := "0,x,0,g\ng,0,0,e\n0,0,0,0\nx,b,0,0"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,g,g\n0,0,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Right, command.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Won)
	})

	t.Run("expect brynjolf to lose on execution of lu", func(t *testing.T) {
		data := "0,x,0,0\ng,0,0,e\n0,0,0,g\n0,0,0,b"
		room, _ := _room.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "g,x,0,0\n0,0,0,e\n0,0,0,0\n0,0,0,0"
		want, _ := _room.NewRoom(data)

		roomSimulator.Start([]command.Command{command.Left, command.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), status.Lost)
	})
}
