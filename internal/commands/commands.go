package commands

import (
	"fmt"

	"github.com/threadactor777/gator/internal/config"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CliCommands map[string]func(*State, Command) error
}

// This method runs a given command with the provided state if it exists.
func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.CliCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("Unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}

// This method registers a new handler function for a command name
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CliCommands[name] = f
}
