package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

const uniqueViolation = "23505"

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) ListOrganizations(ctx context.Context) ([]*entities.OrganizationSummary, error) {
	var rows []*entities.OrganizationSummary
	err := r.db.WithContext(ctx).Raw(`
		SELECT o.id, o.name, o.email, o.phone, o.address, o.created_at,
		       COUNT(u.id) FILTER (WHERE u.deleted_at IS NULL) AS user_count
		FROM organizations o
		LEFT JOIN users u ON u.organization_id = o.id
		WHERE o.deleted_at IS NULL
		GROUP BY o.id
		ORDER BY o.created_at`).Scan(&rows).Error
	return rows, err
}

func (r *adminRepository) CreateOrganizationWithUser(ctx context.Context, org *entities.Organization, user *entities.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := ensureEmailsFree(tx, org.Email, user.Email); err != nil {
			return err
		}
		gormOrg := models.Organization{Name: org.Name, Email: org.Email, Phone: org.Phone, Address: org.Address}
		if err := tx.Create(&gormOrg).Error; err != nil {
			return mapWriteError(err)
		}
		org.ID = gormOrg.ID.String()
		return createUser(tx, gormOrg.ID, user)
	})
}

func (r *adminRepository) ListUsers(ctx context.Context, organizationID string) ([]*entities.User, error) {
	orgID, err := r.liveOrganization(r.db.WithContext(ctx), organizationID)
	if err != nil {
		return nil, err
	}
	var users []*entities.User
	err = r.db.WithContext(ctx).Table("users").
		Select("id, organization_id, name, surname, email, phone, role, created_at").
		Where("organization_id = ? AND deleted_at IS NULL", orgID).
		Order("created_at").Scan(&users).Error
	return users, err
}

func (r *adminRepository) CreateUser(ctx context.Context, organizationID string, user *entities.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orgID, err := r.liveOrganization(tx, organizationID)
		if err != nil {
			return err
		}
		if err := ensureEmailsFree(tx, "", user.Email); err != nil {
			return err
		}
		return createUser(tx, orgID, user)
	})
}

func (r *adminRepository) liveOrganization(db *gorm.DB, id string) (uuid.UUID, error) {
	orgID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, domainErrors.ErrOrganizationNotFound
	}
	var count int64
	if err := db.Model(&models.Organization{}).Where("id = ?", orgID).Count(&count).Error; err != nil {
		return uuid.Nil, err
	}
	if count == 0 {
		return uuid.Nil, domainErrors.ErrOrganizationNotFound
	}
	return orgID, nil
}

func createUser(tx *gorm.DB, orgID uuid.UUID, user *entities.User) error {
	gormUser := models.User{
		OrganizationID: orgID, Name: user.Name, Surname: user.Surname, Email: user.Email,
		Phone: user.Phone, Password: user.Password, Role: user.Role,
	}
	if err := tx.Create(&gormUser).Error; err != nil {
		return mapWriteError(err)
	}
	user.ID = gormUser.ID.String()
	user.OrganizationID = orgID.String()
	user.CreatedAt = gormUser.CreatedAt
	return nil
}

// ensureEmailsFree gives a clear error before insert; the unique indexes remain the final guard.
func ensureEmailsFree(tx *gorm.DB, orgEmail, userEmail string) error {
	var count int64
	if orgEmail != "" {
		if err := tx.Model(&models.Organization{}).Where("email = ?", orgEmail).Count(&count).Error; err != nil {
			return err
		}
	}
	if count == 0 && userEmail != "" {
		if err := tx.Model(&models.User{}).Where("email = ?", userEmail).Count(&count).Error; err != nil {
			return err
		}
	}
	if count > 0 {
		return domainErrors.ErrEmailTaken
	}
	return nil
}

func mapWriteError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
		return domainErrors.ErrEmailTaken
	}
	return err
}
