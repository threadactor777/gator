package commands

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Could not get feeds\n")
	}

	for _, value := range feeds {
		fmt.Println("Name: %v\nURL: %v\nUser: %v\n", value.FeedName, value.FeedUrl, value.Username)
	}
	return nil
}
