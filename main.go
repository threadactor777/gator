package main

import (
	"fmt"
	"os"

	"github.com/threadactor777/gator/internal/commands"
	"github.com/threadactor777/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("Error:", err)
	}

	s := commands.State{
		Cfg: &cfg,
	}

	c := commands.Commands{
		CliCommands: make(map[string]func(*commands.State, commands.Command) error),
	}

	c.Register("login", commands.HandlerLogin)

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
