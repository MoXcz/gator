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

	feed, err := s.db.AddFeed(context.Background(),
		database.AddFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      feedName,
			Url:       feedURL,
			UserID:    user.ID,
		})

	if err != nil {
		return fmt.Errorf("Error: Could not add feed, %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error: Could not follow feed: %w", err)
	}

	fmt.Println("Feed added successfuly:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("Feed followed successfuly")
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

	if len(feeds) <= 0 {
		fmt.Println("User has not added any feeds")
		return nil
	}

	fmt.Printf("Found %d feeds: \n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error reading user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println()
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Println("ID: ", feed.ID)
	fmt.Println("Created: ", feed.CreatedAt)
	fmt.Println("Updated: ", feed.UpdatedAt)
	fmt.Println("Name: ", feed.Name)
	fmt.Println("URL: ", feed.Url)
	fmt.Println("Created: ", feed.CreatedAt)
	fmt.Println("User: ", user.Name)
}
