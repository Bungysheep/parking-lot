package command

import (
	"strings"
	"utils/constant"
)

// ICommand interface
type ICommand interface {
	GetActionName() string
	Parse(argString string) error
	Validate() error
	Execute() string
	Clear()
}

// Command type
type Command struct {
	ActionName string
	Args       []string
}

// NewCommand - Create a new Command
func NewCommand() *Command {
	return new(Command)
}

// GetActionName - Return the action name of a command
func (c *Command) GetActionName() string {
	return c.ActionName
}

// Parse - Parse an argument string of a command
func (c *Command) Parse(argString string) error {
	c.Args = strings.Split(argString, constant.WhiteSpace)
	return nil
}

// Validate - Validate the arguments of a command
func (c *Command) Validate() error {
	return nil
}

// Execute - Run a command
func (c *Command) Execute() string {
	return constant.EmptyString
}

// Clear - Clear a command
func (c *Command) Clear() {
	c.Args = []string{}
}
