package configs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitRequiresAuthConfiguration(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "main.yml"), []byte("http:\n  port: 8000\n"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("APP_ENV", "local")
	t.Setenv("USER_ACCESS_TOKEN_SECRET", "")
	if _, err := Init(dir); err == nil || !strings.Contains(err.Error(), "USER_ACCESS_TOKEN_SECRET") {
		t.Fatalf("expected named missing secret error, got %v", err)
	}
	t.Setenv("USER_ACCESS_TOKEN_SECRET", "test-only-access")
	t.Setenv("USER_REFRESH_TOKEN_SECRET", "test-only-refresh")
	t.Setenv("PASSWORD_SALT", "test-only-salt")
	t.Setenv("ACCESS_TOKEN_LIFETIME_MINUTES", "15")
	t.Setenv("REFRESH_TOKEN_LIFETIME_MINUTES", "1440")
	t.Setenv("HTTP_PORT", "8001")
	cfg, err := Init(dir)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.HTTP.Host != "127.0.0.1" || cfg.HTTP.Port != "8001" {
		t.Fatalf("unexpected listener %s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	}
	t.Setenv("HTTP_PORT", "0")
	if _, err := Init(dir); err == nil || !strings.Contains(err.Error(), "HTTP_PORT") {
		t.Fatalf("expected invalid port error, got %v", err)
	}
}
