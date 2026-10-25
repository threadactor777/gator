package commands

import (
	"context"
	"fmt"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Arguments is empty\n")
	}

	username := cmd.Args[0]

	_, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("Error: User '%s' does not exist.\n", username)
		os.Exit(1)
	}

	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Could not set username\n")
	}

	fmt.Printf("The user has been set!\n")

	return nil
}
