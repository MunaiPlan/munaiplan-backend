package postgres

import (
	"strings"
	"testing"
)

func TestOwnershipQueryWalksToCompanies(t *testing.T) {
	q := ownershipQuery("strings")
	for _, want := range []string{"FROM strings t0", "JOIN cases t1 ON t1.id = t0.case_id", "JOIN trajectories t2", "JOIN companies t8 ON t8.id = t7.company_id", "SELECT t8.organization_id"} {
		if !strings.Contains(q, want) {
			t.Fatalf("query missing %q:\n%s", want, q)
		}
	}
	if q := ownershipQuery("companies"); !strings.HasPrefix(q, "SELECT t0.organization_id::text FROM companies t0 WHERE") {
		t.Fatalf("companies query: %s", q)
	}
	// Every level must exclude soft-deleted rows.
	if n := strings.Count(ownershipQuery("strings"), "deleted_at IS NULL"); n != 9 {
		t.Fatalf("expected 9 live-row filters, got %d", n)
	}
}
