package command

import (
	"errors"
	"models/car"
	"models/parkingcenter"
	"utils/constant"
	"utils/message"
)

// ParkCommand type
type ParkCommand struct {
	Command
	RegNbr string
	Colour string
}

// NewParkCommand - Create a new park command
func NewParkCommand() *ParkCommand {
	p := new(ParkCommand)
	p.ActionName = constant.ParkAction

	return p
}

// Parse - Parse an argument string of park command
func (p *ParkCommand) Parse(argString string) error {
	p.Command.Parse(argString)
	if len(p.Args) != 2 || p.Args[0] == constant.EmptyString || p.Args[1] == constant.EmptyString {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	p.RegNbr = p.Args[0]
	p.Colour = p.Args[1]

	return nil
}

// Validate - Validate whether the parameters of park command is valid
func (p *ParkCommand) Validate() error {
	return nil
}

// Execute - Run a park command
func (p *ParkCommand) Execute() string {
	parking := parkingcenter.Get().GetParking()
	if parking.GetTotalCapacity() <= constant.MinimumCapacity {
		return message.ParkingHasNotBeenCreated()
	}

	if !parking.HasAvailableSlot() {
		return message.ParkingIsFull()
	}

	car := car.New(p.RegNbr, p.Colour)
	slot := parking.EnterCar(car)

	return message.CarEntered(slot.GetSlotNbr())
}
