package postgres

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/munaiplan/munaiplan-backend/infrastructure/database/postgres/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
    dbUsernameKey = "DB_USERNAME"
    dbPasswordKey = "DB_PASSWORD"
    dbHostKey     = "DB_HOST"
    dbPortKey     = "DB_PORT"
    dbNameKey     = "DB_NAME"
)

var (
    dbInstance *Database
    once       sync.Once
)

type Database struct {
    Conn *gorm.DB
}

type DbCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Host     string `json:"host"`
    Dbname   string `json:"dbname"`
    Port     string `json:"port"`
}

func NewDatabase() *Database {
    once.Do(func() {
        dbCredentials := DbCredentials{
            Username: getEnv(dbUsernameKey, "default_user"),
            Password: getEnv(dbPasswordKey, "default_password"),
            Host:     getEnv(dbHostKey, "localhost"),
            Port:     getEnv(dbPortKey, "5432"),
            Dbname:   getEnv(dbNameKey, "default_db"),
        }

        dsn := fmt.Sprintf(
            "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
            dbCredentials.Host, dbCredentials.Username, dbCredentials.Password, dbCredentials.Dbname, dbCredentials.Port,
        )

        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
        })
        if err != nil {
            logrus.Fatalf("failed to connect database: %v", err)
        }

        // Create UUID extension
        if err := createUUIDExtension(db); err != nil {
            logrus.Fatalf("failed to create UUID extension: %v", err)
        }

        // Auto-migrate the database schema
        err = db.AutoMigrate(
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
            logrus.Fatalf("failed to auto-migrate database: %v", err)
        }

        err = seedData(db, "infrastructure/database/postgres/seeds/users.sql")
        if err != nil {
            logrus.Fatalf("failed to connect database: %v", err)
        }

        dbInstance = &Database{Conn: db}
        logrus.Info("Database connection established and migrated")
    })

    return dbInstance
}

func seedData(db *gorm.DB, seedFile string) error {
    absPath, err := filepath.Abs(seedFile)
    if err != nil {
        return fmt.Errorf("could not determine absolute path: %v", err)
    }

    content, err := os.ReadFile(absPath)
    if err != nil {
        return fmt.Errorf("could not read seed file: %v", err)
    }

    err = db.Exec(string(content)).Error
    if err != nil {
        return fmt.Errorf("could not execute seed file: %v", err)
    }

    return nil
}

func createUUIDExtension(db *gorm.DB) error {
    return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
