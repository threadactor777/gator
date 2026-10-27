package main

// goose postgres 'postgres://postgres:@localhost:5432/gator' up
// goose postgres 'postgres://postgres:@localhost:5432/gator' down

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/threadactor777/gator/internal/commands"
	"github.com/threadactor777/gator/internal/config"
	"github.com/threadactor777/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("Error:", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Errorf("Error:", err)
	}

	dbQueries := database.New(db)

	s := commands.State{
		Cfg: &cfg,
		Db:  dbQueries,
	}

	c := commands.Commands{
		CliCommands: make(map[string]func(*commands.State, commands.Command) error),
	}

	c.Register("login", commands.HandlerLogin)
	c.Register("register", commands.HandlerRegister)
	c.Register("reset", commands.HandlerReset)
	c.Register("users", commands.HandlerUsers)
	c.Register("agg", commands.HandlerFetchFeed)
	c.Register("addfeed", commands.HandlerAddFeed)
	c.Register("feeds", commands.HandlerFeeds)

	if len(os.Args) < 2 {
		fmt.Println("error: insufficient arguments")
		fmt.Println("usage: gator <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := commands.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = c.Run(&s, cmd)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

}
