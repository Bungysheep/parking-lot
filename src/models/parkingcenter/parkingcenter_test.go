package parkingcenter

import (
	"models/parking"
	"testing"
)

func TestCreateParkingCenter(t *testing.T) {
	pc := New()

	if pc == nil {
		t.Errorf("Expected to have a new Parking Center.")
	}

	if Get() == nil {
		t.Errorf("Expected to have a Parking Center.")
	}

	if Get().Parking == nil {
		t.Errorf("Expected to have a Parking.")
	}

	if Get().Parking.GetTotalCapacity() != 1 {
		t.Errorf("Expected to have a default Parking Center with Capacity %d, but got %d.", 1, Get().Parking.GetTotalCapacity())
	}
}

func TestCreateSetGetParking(t *testing.T) {
	pc := New()

	if pc == nil {
		t.Errorf("Expected to have a new Parking Center.")
	}

	if Get() == nil {
		t.Errorf("Expected to have a Parking Center.")
	}

	Get().SetParking(parking.New(5))

	if Get().GetParking() == nil {
		t.Errorf("Expected to have a Parking.")
	}

	if Get().GetParking().GetTotalCapacity() != 5 {
		t.Errorf("Expected to have a default Parking Center with Capacity %d, but got %d.", 5, Get().GetParking().GetTotalCapacity())
	}
}
