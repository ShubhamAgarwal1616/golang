package test

import (
	simulator "legacy-of-brynjolf/simulator"
	"reflect"
	"testing"
)

func TestCommands(t *testing.T) {

	checkCommand := func(t *testing.T, got, want simulator.Command) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
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

	t.Run("get a valid command", func(t *testing.T) {
		got, err := simulator.GetValidCommand("d")
		want := simulator.Command("d")
		checkCommand(t, got, want)
		assertNoError(t, err)
	})

	t.Run("got an error for invalid command", func(t *testing.T) {
		_, err := simulator.GetValidCommand("s")
		assertError(t, err)
	})

}
