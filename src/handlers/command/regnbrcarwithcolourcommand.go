package command

import (
	"errors"
	"models/parkingcenter"
	"strings"
	"utils/constant"
	"utils/message"
)

// RegNbrCarWithColourCommand type
type RegNbrCarWithColourCommand struct {
	Command
	Colour string
}

// NewRegNbrCarWithColourCommand - Create a new registration_numbers_for_cars_with_colour command
func NewRegNbrCarWithColourCommand() *RegNbrCarWithColourCommand {
	rn := new(RegNbrCarWithColourCommand)
	rn.ActionName = constant.RegistrationNumberOfCarsByColourAction

	return rn
}

// Parse - Parse an argument string of registration_numbers_for_cars_with_colour command
func (rn *RegNbrCarWithColourCommand) Parse(argString string) error {
	rn.Command.Parse(argString)
	if len(rn.Args) != 1 || rn.Args[0] == constant.EmptyString {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	rn.Colour = rn.Args[0]

	return nil
}

// Validate - Validate whether the parameters of registration_numbers_for_cars_with_colour command is valid
func (rn *RegNbrCarWithColourCommand) Validate() error {
	return nil
}

// Execute - Run a registration_numbers_for_cars_with_colour command
func (rn *RegNbrCarWithColourCommand) Execute() string {
	result := parkingcenter.Get().GetParking().GetRegNbrByColour(rn.Colour)

	if len(result) == 0 {
		return message.NotFound()
	}

	return strings.Join(result, constant.Comma)
}
