package postgres

import (
	"fmt"
	"os"
	"strings"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

// Account describes one provisioned login and the organization it belongs to.
type Account struct {
	Email        string
	Password     string
	Organization string
	Name         string
	Surname      string
	Role         string
}

// SeedLocalUser is an opt-in fixture for a disposable local database.
func SeedLocalUser(db *gorm.DB) error {
	account, err := accountFromEnv("DEV_SEED", entities.RoleUser, "Local", "Developer")
	if err != nil {
		return err
	}
	return EnsureAccount(db, account)
}

// BootstrapAdmin creates the first administrator from ADMIN_EMAIL, ADMIN_PASSWORD and
// optional ADMIN_ORGANIZATION. Further organizations and users are created in the admin panel.
func BootstrapAdmin(db *gorm.DB) error {
	if strings.TrimSpace(os.Getenv("ADMIN_ORGANIZATION")) == "" {
		os.Setenv("ADMIN_ORGANIZATION", "MunaiPlan Administration")
	}
	account, err := accountFromEnv("ADMIN", entities.RoleAdmin, "Platform", "Administrator")
	if err != nil {
		return err
	}
	return EnsureAccount(db, account)
}

func accountFromEnv(prefix, role, name, surname string) (Account, error) {
	for _, key := range []string{prefix + "_EMAIL", prefix + "_PASSWORD", prefix + "_ORGANIZATION", "PASSWORD_SALT"} {
		if strings.TrimSpace(os.Getenv(key)) == "" {
			return Account{}, fmt.Errorf("required environment variable %s is empty", key)
		}
	}
	return Account{
		Email:        os.Getenv(prefix + "_EMAIL"),
		Password:     os.Getenv(prefix + "_PASSWORD"),
		Organization: os.Getenv(prefix + "_ORGANIZATION"),
		Name:         name,
		Surname:      surname,
		Role:         role,
	}, nil
}

// EnsureAccount idempotently creates the organization (keyed by the account email) and the user.
// An existing account is left untouched; a role mismatch is reported rather than changed.
func EnsureAccount(db *gorm.DB, a Account) error {
	email := strings.ToLower(strings.TrimSpace(a.Email))
	if !strings.Contains(email, "@") {
		return fmt.Errorf("account email must be an email address")
	}
	if err := helpers.ValidatePassword(a.Password); err != nil {
		return err
	}
	name := strings.TrimSpace(a.Organization)
	return db.Transaction(func(tx *gorm.DB) error {
		var user models.User
		err := tx.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
		if err == nil {
			if user.Role != a.Role {
				return fmt.Errorf("account already exists with role %q; refusing to change it to %q", user.Role, a.Role)
			}
			return nil
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}
		var org models.Organization
		err = tx.Where("email = ? AND deleted_at IS NULL", email).First(&org).Error
		if err == gorm.ErrRecordNotFound {
			org = models.Organization{Name: name, Email: email}
			if err := tx.Create(&org).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else if org.Name != name {
			return fmt.Errorf("organization name differs from the existing organization for %s", email)
		}
		hash, err := helpers.HashPassword(a.Password)
		if err != nil {
			return err
		}
		return tx.Create(&models.User{OrganizationID: org.ID, Name: a.Name, Surname: a.Surname, Email: email, Password: hash, Role: a.Role}).Error
	})
}
