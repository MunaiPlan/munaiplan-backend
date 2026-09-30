package postgres

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct{ Conn *gorm.DB }

// Open connects only. Schema changes belong to the explicit migrate command.
func Open() (*Database, error) {
	keys := []string{"DB_HOST", "DB_PORT", "DB_NAME", "DB_USERNAME", "DB_PASSWORD"}
	for _, key := range keys {
		if strings.TrimSpace(os.Getenv(key)) == "" {
			return nil, fmt.Errorf("required environment variable %s is empty", key)
		}
	}
	dsn := buildDSN(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}
	return &Database{Conn: db}, nil
}

// buildDSN URL-encodes every component. An unquoted key=value DSN silently
// truncates values containing spaces or quotes and fails authentication.
func buildDSN(host, port, name, user, password string) string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, password),
		Host:     net.JoinHostPort(host, port),
		Path:     "/" + name,
		RawQuery: url.Values{"sslmode": {"disable"}, "TimeZone": {"UTC"}}.Encode(),
	}
	return u.String()
}

func (db *Database) Close() error {
	sqlDB, err := db.Conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
