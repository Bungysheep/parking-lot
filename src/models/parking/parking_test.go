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

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
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

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
	}

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

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
	}

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

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
	}

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

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
	}

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

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))
	s4 := p.EnterCar(car.New("KA-01-HH-2701", "Blue"))
	s5 := p.EnterCar(car.New("KA-01-HH-3141", "Black"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
	}

	if s4.GetSlotNbr() != 4 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 4, s4.GetSlotNbr())
	}

	if s5.GetSlotNbr() != 5 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 5, s5.GetSlotNbr())
	}

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

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))
	s4 := p.EnterCar(car.New("KA-01-HH-2701", "Blue"))
	s5 := p.EnterCar(car.New("KA-01-HH-3141", "Black"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
	}

	if s4.GetSlotNbr() != 4 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 4, s4.GetSlotNbr())
	}

	if s5.GetSlotNbr() != 5 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 5, s5.GetSlotNbr())
	}

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

func TestGetUnavailableSlots(t *testing.T) {
	p := New(5)

	if p == nil {
		t.Errorf("Expected to have a new Parking.")
	}

	s1 := p.EnterCar(car.New("KA-01-HH-1234", "White"))
	s2 := p.EnterCar(car.New("KA-01-BB-0001", "Black"))
	s3 := p.EnterCar(car.New("KA-01-HH-7777", "Red"))

	if s1.GetSlotNbr() != 1 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 1, s1.GetSlotNbr())
	}

	if s2.GetSlotNbr() != 2 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 2, s2.GetSlotNbr())
	}

	if s3.GetSlotNbr() != 3 {
		t.Errorf("Expected the Car is parked at Slot %d, but got it is parked at Slot %d", 3, s3.GetSlotNbr())
	}

	if p.GetTotalCapacity() != 5 {
		t.Errorf("Expected to have Parking with Total Capacity %d, but got %d", 10, p.GetTotalCapacity())
	}

	if !p.HasAvailableSlot() {
		t.Errorf("Expected to have Parking available slot, but got no available slot")
	}

	us := p.GetUnavailableSlots()

	if len(us) != 3 {
		t.Errorf("Expected to have %d unavailable slots, but got %d", 3, len(us))
	}

	for i, s := range us {
		if s.GetSlotNbr() != i+1 {
			t.Errorf("Slot %d - Expected to have Slot Number %d, but got %d", i+1, i+1, s.GetSlotNbr())
		}

		if s.IsAvailable() {
			t.Errorf("Slot %d - Expected to have a Car in this Slot, but got no Car", s.GetSlotNbr())
		}
	}
}
