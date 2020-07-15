package test

import (
	"legacy-of-brynjolf/entities"
	"reflect"
	"testing"
)

func TestEntities(t *testing.T) {

	checkEntity := func(t *testing.T, got, want entities.RoomEntity) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("an error should have occur")
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("an error should not have occur")
		}
	}

	t.Run("get a valid entitty", func(t *testing.T) {
		got, err := entities.BuildEntity("b")
		want := entities.RoomEntity("b")
		checkEntity(t, got, want)
		assertNoError(t, err)
	})

	t.Run("got an error for invalid entity", func(t *testing.T) {
		_, err := entities.BuildEntity("s")
		assertError(t, err)
	})
}

func TestGetBlockingEntities(t *testing.T) {
	checkBlockingEntities := func(t *testing.T, got, want []entities.RoomEntity) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("get wall as blocking entity for brynjolf", func(t *testing.T) {
		got := entities.RoomEntity("b").GetBlockingEntities()
		want := []entities.RoomEntity{wall}
		checkBlockingEntities(t, got, want)
	})

	t.Run("get wall and exit as blocking entity for guard", func(t *testing.T) {
		got := entities.RoomEntity("g").GetBlockingEntities()
		want := []entities.RoomEntity{wall, exit}
		checkBlockingEntities(t, got, want)
	})
}