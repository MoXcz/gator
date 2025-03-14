package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MoXcz/gator/internal/database"
	"github.com/google/uuid"
)

// Fetch RSS feed
func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Error while fetching feed: %w", err)
	}

	fmt.Printf("Feed: %+v", feed)
	return nil
}

// Add RSS feed
func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feedParams := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}
	feed, err := s.db.AddFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("Error: Could not add feed, %w", err)
	}

	fmt.Println("Added record: ")
	fmt.Println("ID: ", feed.ID)
	fmt.Println("Name: ", feed.Name)
	fmt.Println("URL: ", feed.Url)
	fmt.Println("Created: ", feed.CreatedAt)
	fmt.Println("On User: ", user.Name)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error: Could not get feeds: %w", err)
	}

	fmt.Printf("Found %d feeds: \n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error reading user: %w", err)
		}
		fmt.Printf("%s added: (%s, %s)\n", user.Name, feed.Name, feed.Url)
	}

	return nil
}
