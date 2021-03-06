package test

import (
	_room "legacy-of-brynjolf/room"
	_blocks "legacy-of-brynjolf/room/blocks"
	"legacy-of-brynjolf/room/entities"
	"reflect"
	"testing"
)

var wall = entities.Entity('x')
var brynjolf = entities.Entity('b')
var guard = entities.Entity('g')
var exit = entities.Entity('e')

func TestFindEntitiesPosition(t *testing.T) {
	checkPositions := func(t *testing.T, got, want []_blocks.Block) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("get positions for brynjolf in room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,b,0,0\nx,0,0,0"
		room, _ := _room.NewRoom(data)
		got := room.GetBlocks([]entities.Entity{brynjolf})
		want := []_blocks.Block{_blocks.NewBlock(brynjolf, 2, 1)}
		checkPositions(t, got, want)
	})

	t.Run("get positions for brynjolf and guard in room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,b,g,0\nx,0,0,0"
		room, _ := _room.NewRoom(data)
		got := room.GetBlocks([]entities.Entity{brynjolf, guard})
		want := []_blocks.Block{_blocks.NewBlock(guard, 2, 2), _blocks.NewBlock(brynjolf, 2, 1)}
		checkPositions(t, got, want)
	})

	t.Run("get positions of exit in room", func(t *testing.T) {
		data := "0,x,0,x\n0,0,0,e\n0,b,g,0\nx,0,0,0"
		room, _ := _room.NewRoom(data)
		got := room.GetBlocks([]entities.Entity{exit})
		want := []_blocks.Block{_blocks.NewBlock(exit, 1, 3)}
		checkPositions(t, got, want)
	})
}

func TestNewRoom(t *testing.T){
	t.Run("expect to build a room object", func(t *testing.T) {
		data := "0,x,0,0\ng,0,0,e\n0,0,0,g\n0,0,0,b"
		_, err := _room.NewRoom(data)

		if err != nil {
			t.Error("error should not have occur in building a room")
		}
	})

	t.Run("expect to get an error in building a room object", func(t *testing.T) {
		data := "0,x,0,0\ng,s,0,e\n0,0,0,g\n0,0,0,b"
		_, err := _room.NewRoom(data)

		if err != entities.InvalidEntityErr {
			t.Errorf("error should have occur in building a room")
		}
	})
}
