package slot

import (
	"models/car"
)

type Slot struct {
	SlotNbr int
	Car     *car.Car
}

// New - Create a new Slot
func New(slotNbr int) *Slot {
	s := new(Slot)
	s.init(slotNbr)

	return s
}

func (s *Slot) init(slotNbr int) {
	s.SlotNbr = slotNbr
	s.Car = nil
}

// GetSlotNbr - Return the slot number of a slot
func (s *Slot) GetSlotNbr() int {
	return s.SlotNbr
}

// GetCar - Return the car parked at a slot
func (s *Slot) GetCar() *car.Car {
	return s.Car
}

// Enter - Assign a car to a slot
func (s *Slot) Enter(c *car.Car) {
	s.Car = c
}

// Exit - Remove a car from a slot
func (s *Slot) Exit() {
	s.Car = nil
}

// IsAvailable - Check whether a slot is free
func (s *Slot) IsAvailable() bool {
	return s.Car == nil
}
