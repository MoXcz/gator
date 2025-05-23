package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/MoXcz/gator/internal/database"
	"github.com/google/uuid"
)

// Fetch RSS feed
func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time>", cmd.Name)
	}
	reqTime, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error: Could not parse time, %w", err)
	}

	fmt.Println("Collecting feeds every", cmd.Args[0])

	ticker := time.NewTicker(reqTime)

	for ; ; <-ticker.C {
		err := scrapeFeed(s)
		if err != nil {
			return err
		}
	}
}

// Add RSS feed
func handlerAddFeed(s *state, cmd command, user database.User) error {
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

func scrapeFeed(s *state) error {
	lastFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error: Could not get last feed, %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), lastFeed.ID)
	if err != nil {
		return fmt.Errorf("Error: Could not mark last feed, %w", err)
	}

	feeds, err := fetchFeed(context.Background(), lastFeed.Url)
	if err != nil {
		return err
	}

	// fmt.Printf("Feed %s with %d posts found:\n", feeds.Channel.Title, len(feeds.Channel.Item))
	// for i, feed := range feeds.Channel.Item {
	// 	fmt.Printf("Feed %d: %s\n", i, feed.Title)
	// }

	fmt.Printf("Feed %s with %d posts found:\n", feeds.Channel.Title, len(feeds.Channel.Item))
	for _, feed := range feeds.Channel.Item {
		var desc sql.NullString
		var pubDate sql.NullTime
		desc.String = feed.Description
		pubDate.Time, err = time.Parse(time.RFC1123, feed.PubDate)
		if err != nil {
			return err
		}
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       feed.Title,
			Url:         feed.Link,
			Description: desc,
			PublishedAt: pubDate,
			FeedID:      lastFeed.ID,
		})
		if err != nil {
			return fmt.Errorf("Error: Could not add post, %w", err)
		}
		fmt.Printf("Post %s added\n", post.Title)
	}

	return nil
}

func handleBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args) == 1 {
		l, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Error: Could not parse string, %w", err)
		}
		limit = int32(l)
	} else if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <limit> (optional)", cmd.Name)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("Error: Could not get posts of user, %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon 1 Jan"), post.FeedID)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
		fmt.Println()
	}

	return nil
}
