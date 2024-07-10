package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/munaiplan/munaiplan-backend/infrastructure/database/postgres/models"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
        &models.User{}, 
        &models.Company{}, 
        &models.Field{}, 
        &models.Site{}, 
        &models.Well{}, 
        &models.Wellbore{}, 
        &models.Design{}, 
        &models.Case{}, 
        &models.Trajectory{}, 
        &models.TrajectoryHeader{}, 
        &models.TrajectoryUnit{},
    )
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
        os.Exit(1)
    }
    io.WriteString(os.Stdout, stmts)
}