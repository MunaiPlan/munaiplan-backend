package main


import "github.com/munaiplan/munaiplan-backend/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
