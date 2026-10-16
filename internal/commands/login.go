package commands

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Arguments is empty")
	}

	username := cmd.Args[0]

	err := s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Could not set username")
	}

	fmt.Printf("The user has been set!")

	return nil
}
