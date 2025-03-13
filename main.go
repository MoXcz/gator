package main

import (
	"fmt"

	"github.com/MoXcz/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	cfg.SetUser("mocos")
}
