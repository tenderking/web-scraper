package cmd

import (
	"fmt"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	registeredCommands map[string]func(Command) error
}

func (c *Commands) Register(name string, f func(Command) error) {
	if c.registeredCommands == nil {
		c.registeredCommands = make(map[string]func(Command) error)
	}
	c.registeredCommands[name] = f
}

func (c *Commands) Run(cmd Command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name) // Return error instead of fatal log
	}
	return handler(cmd)
}
