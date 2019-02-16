package command

import (
	"errors"
	"models/parking"
	"models/parkingcenter"
	"strconv"
	"utils/constant"
	"utils/message"
)

// CreateParkingLotCommand type
type CreateParkingLotCommand struct {
	Command
	Capacity int
}

// NewCreateParkingLotCommand - Create a new create_parking_lot command
func NewCreateParkingLotCommand() *CreateParkingLotCommand {
	cp := new(CreateParkingLotCommand)
	cp.ActionName = constant.CreateParkingLotAction

	return cp
}

// Parse - Parse an argument string of create_parking_lot command
func (cp *CreateParkingLotCommand) Parse(argString string) error {
	cp.Command.Parse(argString)
	if len(cp.Args) != 1 || cp.Args[0] == constant.EmptyString {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	capacity, err := strconv.Atoi(cp.Args[0])
	if err != nil {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	cp.Capacity = capacity

	return nil
}

// Validate - Validate whether the parameters of create_parking_lot command is valid
func (cp *CreateParkingLotCommand) Validate() error {
	if cp.Capacity < constant.MinimumCapacity {
		return errors.New(message.ParkingCapacityLessThanMinimumCapacity())
	}

	return nil
}

// Execute - Run a create_parking_lot command
func (cp *CreateParkingLotCommand) Execute() string {
	parkingcenter.Get().SetParking(parking.New(cp.Capacity))

	return message.ParkingCreated(cp.Capacity)
}
