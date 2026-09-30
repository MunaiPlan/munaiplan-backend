package postgres

import (
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestBuildDSNPreservesAwkwardValues(t *testing.T) {
	for _, password := range []string{"plainhex0123", "env.py first", `q'uo"te`, "a b=c&d/e?f#g@h:i%j"} {
		cfg, err := pgx.ParseConfig(buildDSN("postgres", "5432", "munaiplan_dev", "munaiplan_dev", password))
		if err != nil {
			t.Fatalf("parse DSN: %v", err)
		}
		if cfg.Password != password || cfg.User != "munaiplan_dev" || cfg.Database != "munaiplan_dev" || cfg.Host != "postgres" || cfg.Port != 5432 {
			t.Fatalf("round trip changed connection fields for a %d-byte password", len(password))
		}
		if cfg.RuntimeParams["TimeZone"] != "UTC" {
			t.Fatalf("TimeZone runtime parameter was not preserved")
		}
	}
}
