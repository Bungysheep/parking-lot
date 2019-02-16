package command

import (
	"errors"
	"models/parkingcenter"
	"strings"
	"utils/constant"
	"utils/message"
)

// SlotNbrCarWithColourCommand type
type SlotNbrCarWithColourCommand struct {
	Command
	Colour string
}

// NewSlotNbrCarWithColourCommand - Create a new slot_numbers_for_cars_with_colour command
func NewSlotNbrCarWithColourCommand() *SlotNbrCarWithColourCommand {
	sn := new(SlotNbrCarWithColourCommand)
	sn.ActionName = constant.SlotNumberOfCarsByColourAction

	return sn
}

// Parse - Parse an argument string of slot_numbers_for_cars_with_colour command
func (sn *SlotNbrCarWithColourCommand) Parse(argString string) error {
	sn.Command.Parse(argString)
	if len(sn.Args) != 1 || sn.Args[0] == constant.EmptyString {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	sn.Colour = sn.Args[0]

	return nil
}

// Validate - Validate whether the parameters of slot_numbers_for_cars_with_colour command is valid
func (sn *SlotNbrCarWithColourCommand) Validate() error {
	return nil
}

// Execute - Run a slot_numbers_for_cars_with_colour command
func (sn *SlotNbrCarWithColourCommand) Execute() string {
	result := parkingcenter.Get().GetParking().GetSlotNbrByColour(sn.Colour)

	if len(result) == 0 {
		return message.NotFound()
	}

	return strings.Join(result, constant.Comma)
}
