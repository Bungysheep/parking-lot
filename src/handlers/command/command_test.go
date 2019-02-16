package command

import (
	"testing"
	"utils/message"
)

func TestCreateParkingLotWithoutArgs(t *testing.T) {
	cp := NewCreateParkingLotCommand()
	if cp == nil {
		t.Errorf("Expected to have a new Create Parking Log Command.")
	}

	argString := " "
	err := cp.Parse(argString)
	if err == nil || err.Error() != message.ParameterIsInvalid(argString) {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParameterIsInvalid(argString))
	}
}

func TestCreateParkingLotWithInvalidArgs(t *testing.T) {
	cp := NewCreateParkingLotCommand()
	if cp == nil {
		t.Errorf("Expected to have a new Create Parking Log Command.")
	}

	argString := "ABC"
	err := cp.Parse(argString)
	if err == nil || err.Error() != message.ParameterIsInvalid(argString) {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParameterIsInvalid(argString))
	}
}

func TestCreateParkingLotLessThanMinimumCapacity(t *testing.T) {
	cp := NewCreateParkingLotCommand()
	if cp == nil {
		t.Errorf("Expected to have a new Create Parking Log Command.")
	}

	argString := "0"
	err := cp.Parse(argString)
	if err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	err = cp.Validate()
	if err == nil || err.Error() != message.ParkingCapacityLessThanMinimumCapacity() {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParkingCapacityLessThanMinimumCapacity())
	}
}

func TestCreateParkingLotWithValidArgs(t *testing.T) {
	cp := NewCreateParkingLotCommand()
	if cp == nil {
		t.Errorf("Expected to have a new Create Parking Log Command.")
	}

	argString := "5"
	err := cp.Parse(argString)
	if err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	err = cp.Validate()
	if err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	result := cp.Execute()
	if result != message.ParkingCreated(cp.Capacity) {
		t.Errorf("Expected to have result: %s, but got %s", message.ParkingCreated(cp.Capacity), result)
	}
}
