package main

import (
	"context"
	"fmt"

	"github.com/MoXcz/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	f := func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("Error getting user: %w", err)
		}

		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
	return f
}
