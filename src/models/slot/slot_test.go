package slot

import (
	"models/car"
	"testing"
)

func TestCreateSlot(t *testing.T) {
	s := New(1)

	if s == nil {
		t.Errorf("Expected to have a new Slot.")
	}

	if s.SlotNbr != 1 {
		t.Errorf("Expected to have Slot Number %v, but got %v", 1, s.SlotNbr)
	}

	if s.Car != nil {
		t.Errorf("Expected to have no Car in this Slot, but got a Car")
	}
}

func TestEnterCar(t *testing.T) {
	s := New(1)

	s.Enter(car.New("KA-01-HH-1234", "White"))

	if s.GetSlotNbr() != 1 {
		t.Errorf("Expected to have Slot Number %v, but got %v", 1, s.GetSlotNbr())
	}

	if s.IsAvailable() {
		t.Errorf("Expected to have a Car in this Slot, but got no Car")
	}

	if s.GetCar().GetRegNbr() != "KA-01-HH-1234" {
		t.Errorf("Expected to have Car Registration Number %s, but got %s", "KA-01-HH-1234", s.GetCar().GetRegNbr())
	}

	if s.GetCar().GetColour() != "White" {
		t.Errorf("Expected to have Car Colour %s, but got %s", "White", s.GetCar().GetColour())
	}
}

func TestExitCar(t *testing.T) {
	s := New(1)

	s.Enter(car.New("KA-01-HH-1234", "White"))

	s.Exit()

	if s.GetSlotNbr() != 1 {
		t.Errorf("Expected to have Slot Number %v, but got %v", 1, s.GetSlotNbr())
	}

	if !s.IsAvailable() {
		t.Errorf("Expected to have no Car in this Slot, but got a Car")
	}
}
