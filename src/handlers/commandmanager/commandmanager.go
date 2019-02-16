package commandmanager

import (
	"errors"
	"handlers/command"
	"strings"
	"utils/constant"
	"utils/message"
)

// CommandManager type
type CommandManager struct {
	Action     string
	ActionArgs string
	Commands   map[string]command.ICommand
}

// NewCommandManager - Create a new Command Manager
func NewCommandManager() *CommandManager {
	cm := new(CommandManager)
	cm.init()

	cm.AddCommand(command.NewCreateParkingLotCommand())
	cm.AddCommand(command.NewParkCommand())
	cm.AddCommand(command.NewLeaveCommand())
	cm.AddCommand(command.NewStatusCommand())
	cm.AddCommand(command.NewRegNbrCarWithColourCommand())
	cm.AddCommand(command.NewSlotNbrCarWithColourCommand())
	cm.AddCommand(command.NewSlotNbrCarWithRegNbrCommand())

	return cm
}

func (cm *CommandManager) init() {
	cm.Commands = make(map[string]command.ICommand)
}

// AddCommand - Add a command
func (cm *CommandManager) AddCommand(c command.ICommand) {
	cm.Commands[c.GetActionName()] = c
}

// Parse - Parse an action of a command
func (cm *CommandManager) Parse(action string) error {
	result := strings.SplitN(strings.TrimSpace(action), constant.WhiteSpace, 2)

	if result[0] == constant.EmptyString {
		return errors.New(message.ActionIsInvalid(action))
	}
	cm.Action = strings.ToLower(result[0])

	cm.ActionArgs = constant.EmptyString
	if len(result) > 1 {
		cm.ActionArgs = result[1]
	}

	return nil
}

// Validate - Validate a command
func (cm *CommandManager) Validate() error {
	if cm.Commands[cm.Action] == nil {
		return errors.New(message.CommandIsInvalid(cm.Action))
	}

	return nil
}

// Execute - Run a command
func (cm *CommandManager) Execute(action string) (string, error) {
	if err := cm.Parse(action); err != nil {
		return constant.EmptyString, err
	}

	if err := cm.Validate(); err != nil {
		return constant.EmptyString, err
	}

	c := cm.Commands[cm.Action]
	c.Clear()

	if err := c.Parse(cm.ActionArgs); err != nil {
		return constant.EmptyString, err
	}

	if err := c.Validate(); err != nil {
		return constant.EmptyString, err
	}

	return c.Execute(), nil
}
