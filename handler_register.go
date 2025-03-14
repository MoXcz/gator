package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MoXcz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("Error: User could not be created, %w", err)
	}

	fmt.Printf("User '%v: %v' was added successfully\n", user.ID, user.Name)
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	return nil
}
