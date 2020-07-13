package test

import (
	"legacy-of-brynjolf/simulator"
	"reflect"
	"testing"
)

var wall = simulator.RoomEntity('x')
var space = simulator.RoomEntity('0')
var brynjolf = simulator.RoomEntity('b')
var guard = simulator.RoomEntity('g')
var exit = simulator.RoomEntity('e')

func TestFindEntitiesPosition(t *testing.T) {
	checkPositions := func(t *testing.T, got, want []simulator.Position) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("get positions for brynjolf in room", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
			                                               {space,space,space,exit},
			                                               {space,brynjolf,space,space},
			                                               {wall,space,space,space}})
		got := room.FindEntitiesPosition([]simulator.RoomEntity{brynjolf})
		want := []simulator.Position{simulator.NewPostion(brynjolf, 2, 1)}
		checkPositions(t, got, want)
	})

	t.Run("get positions for brynjolf and guard in room", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
			                                              {space,space,space,exit},
														  {space,brynjolf,guard,space},
														  {wall,space,space,space}})
		got := room.FindEntitiesPosition([]simulator.RoomEntity{brynjolf, guard})
		want := []simulator.Position{simulator.NewPostion(guard, 2, 2),simulator.NewPostion(brynjolf, 2, 1)}
		checkPositions(t, got, want)
	})

	t.Run("get positions of exit in room", func(t *testing.T) {
		room := simulator.NewRoom([][]simulator.RoomEntity{{space, wall,space,wall},
			                                               {space,space,space,exit},
			                                               {space,brynjolf,guard,space},
			                                               {wall,space,space,space}})
		got := room.FindEntitiesPosition([]simulator.RoomEntity{exit})
		want := []simulator.Position{simulator.NewPostion(exit, 1, 3)}
		checkPositions(t, got, want)
	})
}
