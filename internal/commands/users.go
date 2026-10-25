package commands

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not get all users")
	}

	currentUser := s.Cfg.CurrentUserName

	for _, user := range users {
		if user == currentUser {
			fmt.Printf("%v (current)\n", user)
		} else {
			fmt.Printf("%v\n", user)
		}
	}
	return nil
}
