package postgres

import (
	"crypto/sha256"
	"embed"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"gorm.io/gorm"
)

// Migrations are immutable once applied: add a new numbered file instead of editing one.
//
//go:embed migrations/*.sql
var migrationFiles embed.FS

// migrationLockKey serialises concurrent migrate runs across processes.
const migrationLockKey = 7_311_024_001

type migration struct {
	Version  string
	SQL      string
	Checksum string
}

type appliedMigration struct {
	Version  string
	Checksum string
}

func loadMigrations() ([]migration, error) {
	names, err := fs.Glob(migrationFiles, "migrations/*.sql")
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	result := make([]migration, 0, len(names))
	for _, name := range names {
		body, err := migrationFiles.ReadFile(name)
		if err != nil {
			return nil, err
		}
		result = append(result, migration{
			Version:  strings.TrimSuffix(path.Base(name), ".sql"),
			SQL:      string(body),
			Checksum: fmt.Sprintf("%x", sha256.Sum256(body)),
		})
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no embedded migrations")
	}
	return result, nil
}

func readApplied(db *gorm.DB) ([]appliedMigration, error) {
	var rows []appliedMigration
	err := db.Raw("SELECT version, checksum FROM schema_migrations ORDER BY version").Scan(&rows).Error
	return rows, err
}

// checkApplied verifies that applied rows are a checksum-identical prefix of the embedded list.
// It returns how many embedded migrations are already applied.
func checkApplied(known []migration, applied []appliedMigration) (int, error) {
	if len(applied) > len(known) {
		return 0, fmt.Errorf("database has %d migrations but this binary knows %d; deploy a newer binary", len(applied), len(known))
	}
	for i, row := range applied {
		if row.Version != known[i].Version || row.Checksum != known[i].Checksum {
			return 0, fmt.Errorf("migration %s differs from the applied history (found %s); refusing to continue", known[i].Version, row.Version)
		}
	}
	return len(applied), nil
}

// RequireSchema performs only reads. Serving never repairs or seeds a database.
func RequireSchema(db *gorm.DB) error {
	known, err := loadMigrations()
	if err != nil {
		return err
	}
	if !db.Migrator().HasTable("schema_migrations") {
		return fmt.Errorf("schema is missing; run app migrate")
	}
	applied, err := readApplied(db)
	if err != nil {
		return err
	}
	done, err := checkApplied(known, applied)
	if err != nil {
		return err
	}
	if done != len(known) {
		return fmt.Errorf("schema is at %s but this binary needs %s; run app migrate", lastVersion(applied), known[len(known)-1].Version)
	}
	for _, table := range []string{"organizations", "users", "companies", "fields", "sites", "wells", "wellbores", "designs", "trajectories", "trajectory_headers", "trajectory_units", "cases", "holes", "caisings", "strings", "sections", "library_sections", "fluids", "fluid_types", "pore_pressures", "fracture_gradients", "rigs"} {
		if !db.Migrator().HasTable(table) {
			return fmt.Errorf("schema is incomplete: missing table %s", table)
		}
	}
	for _, index := range []string{"idx_organizations_email", "idx_users_email", "idx_companies_name"} {
		var found bool
		if err := db.Raw("SELECT to_regclass(?) IS NOT NULL", "public."+index).Scan(&found).Error; err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("schema is incomplete: missing index %s", index)
		}
	}
	return nil
}

// Migrate applies pending embedded migrations in order, each in its own transaction.
// Repeating it is a no-op. It refuses to adopt an existing unversioned schema.
func Migrate(db *gorm.DB) error {
	known, err := loadMigrations()
	if err != nil {
		return err
	}
	// Pin one connection so the session-level advisory lock covers every step.
	return db.Connection(func(conn *gorm.DB) error {
		if err := conn.Exec("SELECT pg_advisory_lock(?)", migrationLockKey).Error; err != nil {
			return fmt.Errorf("acquire migration lock: %w", err)
		}
		defer conn.Exec("SELECT pg_advisory_unlock(?)", migrationLockKey)

		if err := conn.Exec("CREATE TABLE IF NOT EXISTS schema_migrations (version text PRIMARY KEY, checksum text NOT NULL, applied_at timestamptz NOT NULL DEFAULT now())").Error; err != nil {
			return err
		}
		applied, err := readApplied(conn)
		if err != nil {
			return err
		}
		if len(applied) == 0 {
			var preexisting *string
			if err := conn.Raw("SELECT to_regclass('public.organizations')::text").Scan(&preexisting).Error; err != nil {
				return err
			}
			if preexisting != nil {
				return fmt.Errorf("refusing to baseline an existing unversioned schema")
			}
		}
		done, err := checkApplied(known, applied)
		if err != nil {
			return err
		}
		for _, m := range known[done:] {
			err := conn.Transaction(func(tx *gorm.DB) error {
				if err := tx.Exec(m.SQL).Error; err != nil {
					return fmt.Errorf("apply %s: %w", m.Version, err)
				}
				return tx.Exec("INSERT INTO schema_migrations (version, checksum) VALUES (?, ?)", m.Version, m.Checksum).Error
			})
			if err != nil {
				return err
			}
			fmt.Printf("applied migration %s\n", m.Version)
		}
		return nil
	})
}

func lastVersion(applied []appliedMigration) string {
	if len(applied) == 0 {
		return "no version"
	}
	return applied[len(applied)-1].Version
}
