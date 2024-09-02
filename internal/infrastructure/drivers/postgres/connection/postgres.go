package postgres

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
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

		var db *gorm.DB
		var err error

		for i := 0; i < 5; i++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
			if err == nil {
				break
			}
			fmt.Printf("failed to connect to database: %v\n", err)
			time.Sleep(10 * time.Second)
		}

		if err != nil {
			logrus.Fatalf("failed to connect database: %v", err)
		}

		if err = execSqlFromFile(db, "internal/infrastructure/drivers/postgres/setup/setup.sql"); err != nil {
			logrus.Fatalf("failed to execute setup sql file: %v", err)
		}

		// Auto-migrate the database schema
		err = db.AutoMigrate(
			&models.Organization{},
			&models.User{},
			&models.Company{},
			&models.Field{},
			&models.Site{},
			&models.Well{},
			&models.Wellbore{},
			&models.Design{},
			&models.Trajectory{},
			&models.TrajectoryHeader{},
			&models.TrajectoryUnit{},
			&models.Case{},
			&models.Hole{},
			&models.String{},
			&models.SectionType{},
			&models.SectionAttribute{},
			&models.Section{},
			&models.SectionValue{},
			&models.SectionValueType{},
			&models.Language{},
			&models.SectionAttributeU18n{},
			&models.Fluid{},
			&models.FluidType{},
		)
		if err != nil {
			logrus.Fatalf("failed to auto-migrate database: %v", err)
		}

		if err = execSqlFromFile(db, "internal/infrastructure/drivers/postgres/setup/indexes.sql"); err != nil {
			logrus.Fatalf("failed to execute indexes sql file: %v", err)
		}
		logrus.Print("Indexes created")

		if err = execSqlFromFile(db, "internal/infrastructure/drivers/postgres/setup/seed.sql"); err != nil {
				logrus.Fatalf("failed to execute seed sql file: %v", err)
		}

		dbInstance = &Database{Conn: db}
		logrus.Info("Database connection established and migrated")
	})

	return dbInstance
}

func execSqlFromFile(db *gorm.DB, filePath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("could not determine absolute path: %v", err)
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("could not read sql file: %v", err)
	}

	err = db.Exec(string(content)).Error
	if err != nil {
		return fmt.Errorf("could not execute sql file: %v", err)
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
