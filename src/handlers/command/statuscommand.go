package command

import (
	"errors"
	"fmt"
	"models/parkingcenter"
	"strings"
	"utils/constant"
	"utils/message"
)

// StatusCommand type
type StatusCommand struct {
	Command
}

// NewStatusCommand - Create a new status command
func NewStatusCommand() *StatusCommand {
	s := new(StatusCommand)
	s.ActionName = constant.StatusAction

	return s
}

// Parse - Parse an argument string of status command
func (s *StatusCommand) Parse(argString string) error {
	s.Command.Parse(argString)
	if len(s.Args) != 1 || s.Args[0] != constant.EmptyString {
		return errors.New(message.ParameterIsInvalid(argString))
	}

	return nil
}

// Validate - Validate whether the parameters of status command is valid
func (s *StatusCommand) Validate() error {
	return nil
}

// Execute - Run a status command
func (s *StatusCommand) Execute() string {
	var result = []string{
		fmt.Sprintf("%-10s%-18s%-10s",
			"Slot No.",
			"Registration No",
			"Colour",
		),
	}

	us := parkingcenter.Get().GetParking().GetUnavailableSlots()

	if len(us) == 0 {
		result = []string{"No Data Found"}
	} else {
		for _, slot := range us {
			c := slot.GetCar()
			result = append(
				result,
				fmt.Sprintf(
					"%-10d%-18s%-10s",
					slot.GetSlotNbr(),
					c.GetRegNbr(),
					c.GetColour(),
				),
			)
		}
	}

	return strings.Join(result, constant.NewLine)
}
