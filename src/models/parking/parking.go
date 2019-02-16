package parking

import (
	"models/car"
	"models/slot"
)

// Parking type
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
func (p *Parking) EnterCar(c *car.Car) *slot.Slot {
	as := p.getNearestAvailableSlot()
	if as != nil {
		as.Enter(c)
	}

	return as
}

// ExitCar - Exit car from a parking
func (p *Parking) ExitCar(slotNbr int) {
	p.Slots[slotNbr-1].Exit()
}

// GetUnavailableSlots - Return all unavailable slots
func (p *Parking) GetUnavailableSlots() []*slot.Slot {
	s := make([]*slot.Slot, 0)

	for _, slot := range p.Slots {
		if !slot.IsAvailable() {
			s = append(s, slot)
		}
	}

	return s
}

// GetRegNbrByColour - Return Registration Numbers by Colour
func (p *Parking) GetRegNbrByColour(colour string) []string {
	rn := make([]string, 0)

	for _, slot := range p.Slots {
		if !slot.IsAvailable() {
			c := slot.GetCar()

			if c != nil && c.GetColour() == colour {
				rn = append(rn, c.GetRegNbr())
			}
		}
	}

	return rn
}

// GetSlotNbrByColour - Return Slot Numbers by Colour
func (p *Parking) GetSlotNbrByColour(colour string) []string {
	sn := make([]string, 0)

	for _, slot := range p.Slots {
		if !slot.IsAvailable() {
			c := slot.GetCar()

			if c != nil && c.GetColour() == colour {
				sn = append(sn, string(slot.GetSlotNbr()))
			}
		}
	}

	return sn
}

// GetSlotNbrByRegNbr - Return Slot Numbers by Registration Number
func (p *Parking) GetSlotNbrByRegNbr(regNbr string) []string {
	sn := make([]string, 0)

	for _, slot := range p.Slots {
		if !slot.IsAvailable() {
			c := slot.GetCar()

			if c != nil && c.GetRegNbr() == regNbr {
				sn = append(sn, string(slot.GetSlotNbr()))
			}
		}
	}

	return sn
}

func (p *Parking) getNearestAvailableSlot() *slot.Slot {
	for _, s := range p.Slots {
		if s.IsAvailable() {
			return s
		}
	}

	return nil
}
