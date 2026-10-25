package commands

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not delete all users\n")
	}
	return nil
}
