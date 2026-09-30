package main

import (
	"fmt"
	"github.com/munaiplan/munaiplan-backend/internal"
	"os"
)

func main() {
	command := "serve"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "usage: app [serve|migrate|dev-seed|create-admin]")
		os.Exit(2)
	}
	if err := internal.Run(command, "internal/infrastructure/configs"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
