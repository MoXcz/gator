package main

import (
	"fmt"
	"os"

	"github.com/MoXcz/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	s := state{
		&cfg,
	}
	cmds := commands{
		make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: Insufficient arguments")
		os.Exit(1)
	}

	err = cmds.run(&s, command{args[1], args[2:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
