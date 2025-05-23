package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Burncoat/Blog_Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: 	   uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: 	   name,
		Url: 	   url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %v", err)
	}

	fmt.Println("Feed created successfully")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("==================================================")

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:			%v\n", feed.ID)
	fmt.Printf("* Created:		%v\n", feed.CreatedAt)
	fmt.Printf("* Updated:		%v\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%v\n", feed.Name)
	fmt.Printf("* URL:			%v\n", feed.Url)
	fmt.Printf("* User:			%v\n", user.Name)
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get feeds: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %v", err)
		}
		printFeed(feed, user)
		fmt.Println("==================================================")
	}

	return nil
}