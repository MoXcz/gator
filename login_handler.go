package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) <= 0 {
		return fmt.Errorf("Error: expected a username")
	}

	user := cmd.args[0]

	err := s.cfg.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Println("The user has been set")

	return nil
}
