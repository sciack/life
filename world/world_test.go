package world

import (
	"testing"
)

func TestCanCreateWorld(t *testing.T) {
	got := NewWorld(10)
	if got == nil {
		t.Error("Unable to create the world with size of 10")
	}
	if got.size != 10 {
		t.Errorf("Expecting size of 10 got %v", got.size)
	}
}

func TestCanSurvive(t *testing.T) {
	got := NewWorld(3)
	got.matrix[0][0] = 1
	got.matrix[0][1] = 1
	got.matrix[1][0] = 1

	if got.canSurvive(1, 1) == false {
		t.Errorf("Expecting element in 1 1 can survive, %v", got.matrix)
	}
	if got.canSurvive(0, 0) == false {
		t.Errorf("Expecting element in 0 0 can survive, %v", got.matrix)
	}

	if got.canSurvive(2, 2) == true {
		t.Errorf("Expecting element in 2 2 cannot survive, %v", got.matrix)
	}
}

func TestWillBorn(t *testing.T) {
	got := NewWorld(3)
	got.matrix[0][0] = 1
	got.matrix[0][1] = 1

	if got.willBorn(1, 1) == false {
		t.Errorf("Expecting element in 1 1 will born, %v", got.matrix)
	}
	if got.willBorn(0, 1) == false {
		t.Errorf("Expecting element in 1 0 will born, %v", got.matrix)
	}

	if got.canSurvive(2, 2) == true {
		t.Errorf("Expecting element in 2 2 is still free, %v", got.matrix)
	}
}

func TestNext(t *testing.T) {
	got := NewWorld(3)
	got.SetState(0, 0, true)
	got.SetState(1, 0, true)
	got.SetState(2, 0, true)
	got.SetState(0, 1, true)

	got = got.Next()
	if !got.IsAlive(0, 0) {
		t.Errorf("Element in 0 0 must be alive, with configuration %v", got.matrix)
	}
	if got.IsAlive(1, 1) {
		t.Errorf("Expecting element in 1 1 will not born, %v", got.matrix)
	}
	if !got.IsAlive(2, 1) {
		t.Errorf("Expecting element in 1 2 will born, %v", got.matrix)
	}

	if got.matrix[1][1] != FREE {
		t.Errorf("Expecting element in 2 2 is still free, %v", got.matrix)
	}
}
