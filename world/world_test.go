package world

import (
	"testing"
)

func TestCanCreateWorld(t *testing.T) {
	got := NewWorld(10)
	if got == nil {
		t.Error("Unable to create the world with size of 10")
	}
	if got.Size != 10 {
		t.Errorf("Expecting size of 10 got %v", got.Size)
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

}

func TestWillBorn(t *testing.T) {
	got := NewWorld(3)
	got.matrix[0][0] = 1
	got.matrix[0][1] = 1
	got.matrix[0][2] = 1

	if !got.willBorn(1, 1) {
		t.Errorf("Expecting element in 1 1 will born, %v", got.matrix)
	}
	if !got.willBorn(0, 1)  {
		t.Errorf("Expecting element in 1 0 will born, %v", got.matrix)
	}

}

func TestWillBornInToroWorld(t *testing.T) {
    got := NewWorld(3)
    got.SetState(0,0, true)
    got.SetState(0,2, true)
    got.SetState(2,1, true)

    if !got.willBorn(0, 1) {
        t.Errorf("Expecting element in 1 0 will born, %v", got.matrix)
    }
    if !got.willBorn(1, 1)  {
        t.Errorf("Expecting element in 1 1 will born, %v", got.matrix)
    }

}

func TestNext(t *testing.T) {
	got := NewWorld(3)
	got.SetState(0, 0, true)
	got.SetState(1, 0, true)
	got.SetState(0, 1, true)

	got = got.Next()
	if !got.IsAlive(0, 0) {
		t.Errorf("Element in 0 0 must be alive, with configuration %v", got.matrix)
	}
	if !got.IsAlive(1, 1) {
		t.Errorf("Expecting element in 1 1 will born, %v", got.matrix)
	}


}

func TestAliveShouldBeZeroAtStart(t *testing.T) {
	w := NewWorld(10)
	if w.CountAlive() > 0 {
		t.Fatal("Newly initialized world should be empty")
	}
}

func TestAliveShouldReturnTheExpectedValue(t *testing.T) {
	w := NewWorld(10)
	for i := 0; i < 10; i++ {
		w.matrix[i][i] = ALIVE
	}
	if w.CountAlive() != 10 {
		t.Fatalf("Expecting 10 alive, got %v", w.CountAlive())
	}
}

func TestRandomInitializer(t *testing.T) {
	w := Random(10, func() bool { return true })
	alive := w.CountAlive()
	if alive != 100 {
		t.Errorf("Expected 100 alive, got %v", alive)
	}

}
