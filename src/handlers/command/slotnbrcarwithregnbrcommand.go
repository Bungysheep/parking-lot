package command

import (
	"errors"
	"models/parkingcenter"
	"strings"
	"utils/constant"
	"utils/message"
)

// SlotNbrCarWithRegNbrCommand type
type SlotNbrCarWithRegNbrCommand struct {
	Command
	RegNbr string
}

// NewSlotNbrCarWithRegNbrCommand - Create a new slot_number_for_registration_number command
func NewSlotNbrCarWithRegNbrCommand() *SlotNbrCarWithRegNbrCommand {
	sn := new(SlotNbrCarWithRegNbrCommand)
	sn.ActionName = constant.SlotNumberByRegistrationNumberAction

	return sn
}

// Parse - Parse an argument string of slot_number_for_registration_number command
func (sn *SlotNbrCarWithRegNbrCommand) Parse(argString string) error {
	sn.Command.Parse(argString)
	if len(sn.Args) != 1 {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	sn.RegNbr = sn.Args[0]

	return nil
}

// Validate - Validate whether the parameters of slot_number_for_registration_number command is valid
func (sn *SlotNbrCarWithRegNbrCommand) Validate() error {
	return nil
}

// Execute - Run a slot_number_for_registration_number command
func (sn *SlotNbrCarWithRegNbrCommand) Execute() string {
	result := parkingcenter.Get().GetParking().GetSlotNbrByRegNbr(sn.RegNbr)

	if len(result) == 0 {
		return message.NotFound()
	}

	return strings.Join(result, constant.Comma)
}
