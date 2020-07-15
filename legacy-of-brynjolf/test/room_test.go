package test

import (
	"legacy-of-brynjolf/entities"
	"legacy-of-brynjolf/position"
	"legacy-of-brynjolf/simulator"
	"reflect"
	"testing"
)

var wall = entities.RoomEntity('x')
var brynjolf = entities.RoomEntity('b')
var guard = entities.RoomEntity('g')
var exit = entities.RoomEntity('e')

func TestFindEntitiesPosition(t *testing.T) {
	checkPositions := func(t *testing.T, got, want []position.Position) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("get positions for brynjolf in room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,b,0,0\nx,0,0,0"
		room, _ := simulator.NewRoom(data)
		got := room.FindEntitiesPosition([]entities.RoomEntity{brynjolf})
		want := []position.Position{position.NewPostion(brynjolf, 2, 1)}
		checkPositions(t, got, want)
	})

	t.Run("get positions for brynjolf and guard in room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,b,g,0\nx,0,0,0"
		room, _ := simulator.NewRoom(data)
		got := room.FindEntitiesPosition([]entities.RoomEntity{brynjolf, guard})
		want := []position.Position{position.NewPostion(guard, 2, 2), position.NewPostion(brynjolf, 2, 1)}
		checkPositions(t, got, want)
	})

	t.Run("get positions of exit in room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,b,g,0\nx,0,0,0"
		room, _ := simulator.NewRoom(data)
		got := room.FindEntitiesPosition([]entities.RoomEntity{exit})
		want := []position.Position{position.NewPostion(exit, 1, 3)}
		checkPositions(t, got, want)
	})
}

func TestNewRoom(t *testing.T){
	t.Run("expect to build a room object", func(t *testing.T) {
		data := "0,x,0,0\ng,0,0,e\n0,0,0,g\n0,0,0,b"
		_, err := simulator.NewRoom(data)

		if err != nil {
			t.Error("error should not have occur in building a room")
		}
	})

	t.Run("expect to get an error in building a room object", func(t *testing.T) {
		data := "0,x,0,0\ng,s,0,e\n0,0,0,g\n0,0,0,b"
		_, err := simulator.NewRoom(data)

		if err != entities.InvalidEntityErr {
			t.Errorf("error should have occur in building a room")
		}
	})
}
