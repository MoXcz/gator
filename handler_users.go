package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MoXcz/gator/internal/database"
	"github.com/google/uuid"
)

// Login user (modify config file)
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("Error: user does not exist, %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("user was not set: %w", err)
	}

	fmt.Println("The user has been set")

	return nil
}

// Lists all users
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

// Create new user
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

// Reset created users
func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error: users were not deleted, %w", err)
	}

	fmt.Println("Users table reset successful")
	return nil
}
