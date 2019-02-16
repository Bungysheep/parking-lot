package parking

import (
	"models/car"
	"models/slot"
)

type Parking struct {
	Slots []*slot.Slot
}

// New - Create a new Parking
func New(capacity int) *Parking {
	p := new(Parking)
	p.init(capacity)

	return p
}

func (p *Parking) init(capacity int) {
	p.Slots = make([]*slot.Slot, capacity)

	for i := range p.Slots {
		p.Slots[i] = slot.New(i + 1)
	}
}

// GetTotalCapacity - Return the total capacity of a parking
func (p *Parking) GetTotalCapacity() int {
	return len(p.Slots)
}

// HasAvailableSlot - Check whether a parking has available slot
func (p *Parking) HasAvailableSlot() bool {
	for _, s := range p.Slots {
		if s.IsAvailable() {
			return true
		}
	}

	return false
}

// EnterCar - Enter car to a parking
func (p *Parking) EnterCar(c *car.Car) {
	as := p.getNearestAvailableSlot()
	if as != nil {
		as.Enter(c)
	}
}

// ExitCar - Exit car from a parking
func (p *Parking) ExitCar(slotNbr int) {
	p.Slots[slotNbr-1].Exit()
}

func (p *Parking) getNearestAvailableSlot() *slot.Slot {
	for _, s := range p.Slots {
		if s.IsAvailable() {
			return s
		}
	}

	return nil
}
