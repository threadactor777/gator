package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/threadactor777/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Arguments is empty")
	}
	username := cmd.Args[0]

	_, err := s.Db.GetUser(context.Background(), username)
	if err == nil {
		fmt.Printf("Error: User '%s' already exists.\n", username)
		os.Exit(1)
	}

	dbUser, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("Could not create user")
	}

	err = s.Cfg.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("Could not set user")
	}

	fmt.Printf("User '%s' was created successfully!\n", dbUser.Name)
	log.Printf("Debug - User Data: %+v\n", dbUser)

	return nil

}
