package main

import "github.com/munaiplan/munaiplan-backend/internal/app/server"

const configsDir = "configs"

func main() {
	server.Run(configsDir)
}