package command

import (
	"errors"
	"models/parkingcenter"
	"strconv"
	"utils/constant"
	"utils/message"
)

// LeaveCommand type
type LeaveCommand struct {
	Command
	SlotNbr int
}

// NewLeaveCommand - Create a new leave command
func NewLeaveCommand() *LeaveCommand {
	l := new(LeaveCommand)
	l.ActionName = constant.LeaveAction

	return l
}

// Parse - Parse an argument string of leave command
func (l *LeaveCommand) Parse(argString string) error {
	l.Command.Parse(argString)
	if len(l.Args) != 1 || l.Args[0] == constant.EmptyString {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	slotNbr, err := strconv.Atoi(l.Args[0])
	if err != nil {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	l.SlotNbr = slotNbr

	return nil
}

// Validate - Validate whether the parameters of leave command is valid
func (l *LeaveCommand) Validate() error {
	return nil
}

// Execute - Run a leave command
func (l *LeaveCommand) Execute() string {
	parkingcenter.Get().GetParking().ExitCar(l.SlotNbr)

	return message.CarExit(l.SlotNbr)
}
