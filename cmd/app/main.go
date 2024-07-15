package main

import "github.com/munaiplan/munaiplan-backend/internal"
import _ "ariga.io/atlas-provider-gorm/gormschema"


const configsDir = "internal/infrastructure/configs"

func main() {
	internal.Run(configsDir)
}
