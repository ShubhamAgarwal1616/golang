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
		data := "0,x,0,x\n0,0,0,e\n0,b,0,0\nx,0,0,0"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,x\n0,b,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf to move in upward direction and caught by a guard", func(t *testing.T) {
		data := "0,x,0,x\n0,g,0,e\n0,0,0,0\nx,b,0,0"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,x\n0,g,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Lost)
	})

	t.Run("expect brynjolf to move in upward direction and exit the room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,0,0,0\nx,g,0,b"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,x\n0,g,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Won)
	})

	t.Run("expect brynjolf and guards to move in left direction", func(t *testing.T) {
		data := "0,x,0,g\n0,0,0,e\n0,0,0,0\nx,0,0,b"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,g,0\n0,0,0,e\n0,0,0,0\nx,b,0,0"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Left})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf and guards to move in right direction", func(t *testing.T) {
		data := "0,x,0,g\ng,0,0,e\n0,0,0,0\nx,b,0,0"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,g\n0,0,g,e\n0,0,0,0\nx,0,0,b"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Right})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf and guards to move in downward direction", func(t *testing.T) {
		data := "0,x,0,g\ng,b,0,e\n0,0,0,0\nx,0,0,0"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,0,g\n0,0,0,e\ng,0,0,0\nx,b,0,0"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Down})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Undecided)
	})

	t.Run("expect brynjolf to win on execution of ru", func(t *testing.T) {
		data := "0,x,0,g\ng,0,0,e\n0,0,0,0\nx,b,0,0"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "0,x,g,g\n0,0,0,e\n0,0,0,0\nx,0,0,0"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Right, simulator.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Won)
	})

	t.Run("expect brynjolf to lose on execution of lu", func(t *testing.T) {
		data := "0,x,0,0\ng,0,0,e\n0,0,0,g\n0,0,0,b"
		room, _ := simulator.NewRoom(data)
		roomSimulator := simulator.NewRoomSimulator(room)
		data = "g,x,0,0\n0,0,0,e\n0,0,0,0\n0,0,0,0"
		want, _ := simulator.NewRoom(data)

		roomSimulator.Start([]simulator.Command{simulator.Left, simulator.Up})

		assertRoomState(t, roomSimulator.Room(), want)
		assertStatus(t, roomSimulator.Status(), simulator.Lost)
	})
}
