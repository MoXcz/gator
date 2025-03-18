package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MoXcz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	url := cmd.Args[0]
	currentFeed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error: Could not get feed: %w", err)
	}

	currentUserFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    currentFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error: Could not follow feed: %w", err)
	}

	fmt.Println("Followed feed:")
	fmt.Println("User:", currentUserFollow.UserName)
	fmt.Println("Feed:", currentUserFollow.FeedName)
	return nil
}

func handleListFollowsOfUser(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	userFeedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error: Could not get followed feeds of current user, %w", err)
	}

	if len(userFeedFollows) <= 0 {
		fmt.Println("This user does not have any followed feeds")
	}

	fmt.Println("User feed following:", userFeedFollows[0].Username)
	for _, follow := range userFeedFollows {
		fmt.Println(follow.Feedname)
	}
	return nil
}
