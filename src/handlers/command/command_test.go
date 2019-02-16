package command

import (
	"models/car"
	"models/parking"
	"models/parkingcenter"
	"testing"
	"utils/message"
)

func TestCreateParkingLotWithoutArgs(t *testing.T) {
	cp := NewCreateParkingLotCommand()
	if cp == nil {
		t.Errorf("Expected to have a new Create Parking Lot Command.")
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
		t.Errorf("Expected to have a new Create Parking Lot Command.")
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
		t.Errorf("Expected to have a new Create Parking Lot Command.")
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
		t.Errorf("Expected to have a new Create Parking Lot Command.")
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

func TestParkCarWithoutArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	p := NewParkCommand()
	if p == nil {
		t.Errorf("Expected to have a new Park Command.")
	}

	argString := " "
	err := p.Parse(argString)
	if err == nil || err.Error() != message.ParameterIsInvalid(argString) {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParameterIsInvalid(argString))
	}
}

func TestParkCarWithInvalidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	p := NewParkCommand()
	if p == nil {
		t.Errorf("Expected to have a new Park Command.")
	}

	argString := "ABC"
	err := p.Parse(argString)
	if err == nil || err.Error() != message.ParameterIsInvalid(argString) {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParameterIsInvalid(argString))
	}
}

func TestParkCarWithValidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	p := NewParkCommand()
	if p == nil {
		t.Errorf("Expected to have a new Park Command.")
	}

	if err := p.Parse("KA-01-HH-1234 White"); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if err := p.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if result := p.Execute(); result != message.CarEntered(1) {
		t.Errorf("Expected to have result: %s, but got %s", message.CarEntered(1), result)
	}
}

func TestParkCarWhenParkingIsFull(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	p := NewParkCommand()
	if p == nil {
		t.Errorf("Expected to have a new Park Command.")
	}

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	if err := p.Parse("KA-01-HH-1234 White"); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if err := p.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if result := p.Execute(); result != message.ParkingIsFull() {
		t.Errorf("Expected to have result: %s, but got %s", message.ParkingIsFull(), result)
	}
}

func TestLeaveParkWithoutArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	l := NewLeaveCommand()
	if l == nil {
		t.Errorf("Expected to have a new Leave Command.")
	}

	argString := " "
	err := l.Parse(argString)
	if err == nil || err.Error() != message.ParameterIsInvalid(argString) {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParameterIsInvalid(argString))
	}
}

func TestLeaveParkWithInvalidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	l := NewLeaveCommand()
	if l == nil {
		t.Errorf("Expected to have a new Leave Command.")
	}

	argString := "ABC"
	err := l.Parse(argString)
	if err == nil || err.Error() != message.ParameterIsInvalid(argString) {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParameterIsInvalid(argString))
	}
}

func TestLeaveParkWithInvalidSlotNbr(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	l := NewLeaveCommand()
	if l == nil {
		t.Errorf("Expected to have a new Leave Command.")
	}

	if err := l.Parse("4"); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if err := l.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if result := l.Execute(); result != message.SlotIsInvalid(4) {
		t.Errorf("Expected to have result: %s, but got %s", message.SlotIsInvalid(4), result)
	}
}

func TestLeaveParkFromAvailableSlot(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(4))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	l := NewLeaveCommand()
	if l == nil {
		t.Errorf("Expected to have a new Leave Command.")
	}

	if err := l.Parse("4"); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if err := l.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if result := l.Execute(); result != message.NoCarParkedAtSlot(4) {
		t.Errorf("Expected to have result: %s, but got %s", message.NoCarParkedAtSlot(4), result)
	}
}

func TestLeaveParkWithValidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(3))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	l := NewLeaveCommand()
	if l == nil {
		t.Errorf("Expected to have a new Leave Command.")
	}

	if err := l.Parse("1"); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if err := l.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if result := l.Execute(); result != message.CarExit(1) {
		t.Errorf("Expected to have result: %s, but got %s", message.CarExit(1), result)
	}
}
