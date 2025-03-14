package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error: Could not get users, %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUsername {
			fmt.Printf(" * %s (current)\n", user.Name)
			continue
		}
		fmt.Printf(" * %s\n", user.Name)
	}

	return nil
}
