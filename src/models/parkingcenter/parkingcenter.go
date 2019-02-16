package parkingcenter

import (
	"models/parking"
	"utils/constant"
)

// ParkingCenter type
type ParkingCenter struct {
	Parking *parking.Parking
}

var pc *ParkingCenter

// New - Create a Parking Center
func New() *ParkingCenter {
	if pc == nil {
		pc = new(ParkingCenter)
		pc.init()
	}

	return pc
}

func (pc *ParkingCenter) init() {
	pc.Parking = parking.New(constant.MinimumCapacity)
}

// Get - Return a parking center
func Get() *ParkingCenter {
	return New()
}

// SetParking - Set a parking
func (pc *ParkingCenter) SetParking(p *parking.Parking) {
	pc.Parking = p
}

// GetParking - Return a parking
func (pc *ParkingCenter) GetParking() *parking.Parking {
	return pc.Parking
}
