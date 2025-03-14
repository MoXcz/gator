package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	commandHandlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandHandlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	_, ok := c.commandHandlers[cmd.name]
	if !ok {
		return fmt.Errorf("Error: Not a valid command")
	}

	err := c.commandHandlers[cmd.name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}
