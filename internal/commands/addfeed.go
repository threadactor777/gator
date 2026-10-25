package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/threadactor777/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {

	if len(cmd.Args) > 2 {
		return fmt.Errorf("Not enough arguments to add feed\n")
		os.Exit(1)
	}

	currentUser := s.Cfg.CurrentUserName
	currentUserDetails, err := s.Db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("Could not get user details\n")
	}

	currentUserID := currentUserDetails.ID
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUserID,
	})
	if err != nil {
		return fmt.Errorf("Could not create feed\n")
	}

	fmt.Printf(feed.Url)

	return nil

}
