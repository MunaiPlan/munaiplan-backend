package main

import "github.com/munaiplan/munaiplan-backend/internal"
import _ "ariga.io/atlas-provider-gorm/gormschema"


const configsDir = "infrastructure/configs"

func main() {
	internal.Run(configsDir)
}
