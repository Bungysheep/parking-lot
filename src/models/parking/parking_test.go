package parking

import (
	"models/car"
	"testing"
)

func TestCreateParking(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if !p.HasAvailableSlot() {
		t.Errorf("Expected to have Parking available slot, but got no available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if !s.IsAvailable() {
			t.Errorf("Expected to have no Car in this Slot, but got a Car")
		}
	}
}

func TestEnterCar(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	p.EnterCar(car.New("KA-01-HH-1234", "White"))
	p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if !p.HasAvailableSlot() {
		t.Errorf("Expected to have Parking available slot, but got no available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if i < 3 {
			if s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
			}
		} else {
			if !s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have no Car in this Slot, but got a Car", s.GetSlotNbr())
			}
		}
	}
}

func TestExitFirstCar(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	p.EnterCar(car.New("KA-01-HH-1234", "White"))
	p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	p.ExitCar(1)

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if !p.HasAvailableSlot() {
		t.Errorf("Expected to have Parking available slot, but got no available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if i == 1 || i == 2 {
			if s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
			}
		} else {
			if !s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have no Car in this Slot, but got a Car", s.GetSlotNbr())
			}
		}
	}
}

func TestExitLastCar(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	p.EnterCar(car.New("KA-01-HH-1234", "White"))
	p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	p.ExitCar(3)

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if !p.HasAvailableSlot() {
		t.Errorf("Expected to have Parking available slot, but got no available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if i < 2 {
			if s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
			}
		} else {
			if !s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have no Car in this Slot, but got a Car", s.GetSlotNbr())
			}
		}
	}
}

func TestExitMiddleCar(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	p.EnterCar(car.New("KA-01-HH-1234", "White"))
	p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	p.ExitCar(2)

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if !p.HasAvailableSlot() {
		t.Errorf("Expected to have Parking available slot, but got no available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if i == 0 || i == 2 {
			if s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
			}
		} else {
			if !s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have no Car in this Slot, but got a Car", s.GetSlotNbr())
			}
		}
	}
}

func TestEnterCarOnNearestSlot(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	p.EnterCar(car.New("KA-01-HH-1234", "White"))
	p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	p.ExitCar(2)

	p.EnterCar(car.New("KA-01-HH-2701", "Blue"))

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if !p.HasAvailableSlot() {
		t.Errorf("Expected to have Parking available slot, but got no available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if i < 3 {
			if s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
			}
		} else {
			if !s.IsAvailable() {
				t.Errorf("Slot %d - Expected to have no Car in this Slot, but got a Car", s.GetSlotNbr())
			}
		}
	}
}

func TestEnterCarNoAvailableSlot(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	p.EnterCar(car.New("KA-01-HH-1234", "White"))
	p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	p.EnterCar(car.New("KA-01-HH-7777", "Red"))
	p.EnterCar(car.New("KA-01-HH-2701", "Blue"))
	p.EnterCar(car.New("KA-01-HH-3141", "Black"))

	p.EnterCar(car.New("KA-01-HH-9999", "White"))

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if p.HasAvailableSlot() {
		t.Errorf("Expected to have no Parking available slot, but got a available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if s.IsAvailable() {
			t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
		}
	}
}

func TestParkingNoAvailableSlot(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	p.EnterCar(car.New("KA-01-HH-1234", "White"))
	p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	p.EnterCar(car.New("KA-01-HH-7777", "Red"))
	p.EnterCar(car.New("KA-01-HH-2701", "Blue"))
	p.EnterCar(car.New("KA-01-HH-3141", "Black"))

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if p.HasAvailableSlot() {
		t.Errorf("Expected to have no Parking available slot, but got a available slot")
	}

	for i, s := range p.Slots {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if s.IsAvailable() {
			t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
		}
	}
}
