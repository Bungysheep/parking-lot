package command

import (
	"fmt"
	"models/car"
	"models/parking"
	"models/parkingcenter"
	"strings"
	"testing"
	"utils/constant"
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

func TestCreateParkingLotLessThanEqualMinimumCapacity(t *testing.T) {
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
	if err == nil || err.Error() != message.ParkingCapacityMustBeGreaterMinimumCapacity() {
		t.Errorf("Expected to have an error saying: %s, but got no error", message.ParkingCapacityMustBeGreaterMinimumCapacity())
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

func TestParkCarWhenParkingUncreated(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(constant.MinimumCapacity))

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

	if result := p.Execute(); result != message.ParkingHasNotBeenCreated() {
		t.Errorf("Expected to have result: %s, but got %s", message.ParkingHasNotBeenCreated(), result)
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

func TestStatusWithoutArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	s := NewStatusCommand()
	if s == nil {
		t.Errorf("Expected to have a new Status Command.")
	}

	if err := s.Parse(""); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}
}

func TestStatusWithInvalidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	s := NewStatusCommand()
	if s == nil {
		t.Errorf("Expected to have a new Status Command.")
	}

	if err := s.Parse("ABC DEF"); err.Error() != message.ParameterIsInvalid("ABC DEF") {
		t.Errorf("Expected to have an error saying: %s, but got an error saying: %s", message.ParameterIsInvalid("ABC DEF"), err.Error())
	}
}

func TestStatusWithValidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	s := NewStatusCommand()
	if s == nil {
		t.Errorf("Expected to have a new Status Command.")
	}

	if err := s.Parse(""); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if err := s.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	var expResult = []string{
		fmt.Sprintf("%-10s%-18s%-10s",
			"Slot No.",
			"Registration No",
			"Colour",
		),
		fmt.Sprintf("%-10s%-18s%-10s",
			"1",
			"KA-01-HH-9999",
			"White",
		),
		fmt.Sprintf("%-10s%-18s%-10s",
			"2",
			"KA-01-BB-0001",
			"Black",
		),
		fmt.Sprintf("%-10s%-18s%-10s",
			"3",
			"KA-01-HH-7777",
			"Red",
		),
	}

	if result := s.Execute(); result != strings.Join(expResult, constant.NewLine) {
		t.Errorf("Expected to have result: %s, but got %s", strings.Join(expResult, constant.NewLine), result)
	}
}

func TestStatusWhenNoCarParked(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	s := NewStatusCommand()
	if s == nil {
		t.Errorf("Expected to have a new Status Command.")
	}

	if err := s.Parse(""); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	if err := s.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got an error saying: %s", err.Error())
	}

	var expResult = []string{
		message.NoDataFound(),
	}

	if result := s.Execute(); result != strings.Join(expResult, constant.NewLine) {
		t.Errorf("Expected to have result: %s, but got %s", strings.Join(expResult, constant.NewLine), result)
	}
}

func TestRegNbrByColorWithoutArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	rn := NewRegNbrCarWithColourCommand()
	if rn == nil {
		t.Errorf("Expected to have a new Registration Number with Colour Command.")
	}

	if err := rn.Parse(""); err.Error() != message.ParameterIsInvalid("") {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.ParameterIsInvalid(""), err.Error())
	}
}

func TestRegNbrByColorWithValidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	rn := NewRegNbrCarWithColourCommand()
	if rn == nil {
		t.Errorf("Expected to have a new Registration Number with Colour Command.")
	}

	if err := rn.Parse("White"); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if err := rn.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if result := rn.Execute(); result != "KA-01-HH-9999, KA-01-BB-0001" {
		t.Errorf("Expected to have result %s, but got %s", "KA-01-HH-9999, KA-01-BB-0001", result)
	}
}

func TestRegNbrByColorWhenNotFound(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	rn := NewRegNbrCarWithColourCommand()
	if rn == nil {
		t.Errorf("Expected to have a new Registration Number with Colour Command.")
	}

	if err := rn.Parse("Yellow"); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if err := rn.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if result := rn.Execute(); result != message.NotFound() {
		t.Errorf("Expected to have result %s, but got %s", message.NotFound(), result)
	}
}

func TestSlotNbrByColorWithoutArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	sn := NewSlotNbrCarWithColourCommand()
	if sn == nil {
		t.Errorf("Expected to have a new Slot Number with Colour Command.")
	}

	if err := sn.Parse(""); err.Error() != message.ParameterIsInvalid("") {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.ParameterIsInvalid(""), err.Error())
	}
}

func TestSlotNbrByColorWithValidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	sn := NewSlotNbrCarWithColourCommand()
	if sn == nil {
		t.Errorf("Expected to have a new Slot Number with Colour Command.")
	}

	if err := sn.Parse("White"); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if err := sn.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if result := sn.Execute(); result != "1, 2" {
		t.Errorf("Expected to have result %s, but got %s", "1, 2", result)
	}
}

func TestSlotNbrByColorWhenNotFound(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	sn := NewSlotNbrCarWithColourCommand()
	if sn == nil {
		t.Errorf("Expected to have a new Slot Number with Colour Command.")
	}

	if err := sn.Parse("Yellow"); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if err := sn.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if result := sn.Execute(); result != message.NotFound() {
		t.Errorf("Expected to have result %s, but got %s", message.NotFound(), result)
	}
}

func TestSlotNbrByRegNbrWithoutArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	sn := NewSlotNbrCarWithRegNbrCommand()
	if sn == nil {
		t.Errorf("Expected to have a new Slot Number with Registration Number Command.")
	}

	if err := sn.Parse(""); err.Error() != message.ParameterIsInvalid("") {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.ParameterIsInvalid(""), err.Error())
	}
}

func TestSlotNbrByRegNbrWithValidArgs(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	sn := NewSlotNbrCarWithRegNbrCommand()
	if sn == nil {
		t.Errorf("Expected to have a new Slot Number with Registration Number Command.")
	}

	if err := sn.Parse("KA-01-HH-9999"); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if err := sn.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if result := sn.Execute(); result != "1" {
		t.Errorf("Expected to have result %s, but got %s", "1", result)
	}
}

func TestSlotNbrByRegNbrWhenNotFound(t *testing.T) {
	parkingcenter.Get().SetParking(parking.New(5))

	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-9999", "White"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-BB-0001", "Black"))
	parkingcenter.Get().GetParking().EnterCar(car.New("KA-01-HH-7777", "Red"))

	sn := NewSlotNbrCarWithRegNbrCommand()
	if sn == nil {
		t.Errorf("Expected to have a new Slot Number with Registration Number Command.")
	}

	if err := sn.Parse("Yellow"); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if err := sn.Validate(); err != nil {
		t.Errorf("Expected to have no error, but got error saying: %s", err.Error())
	}

	if result := sn.Execute(); result != message.NotFound() {
		t.Errorf("Expected to have result %s, but got %s", message.NotFound(), result)
	}
}
