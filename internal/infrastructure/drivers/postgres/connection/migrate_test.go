package postgres

import (
	"strings"
	"testing"
)

func TestEmbeddedMigrationsAreOrderedAndBaselineUnchanged(t *testing.T) {
	known, err := loadMigrations()
	if err != nil {
		t.Fatal(err)
	}
	if len(known) < 5 || known[0].Version != "0001_initial" || known[1].Version != "0002_user_roles" || known[2].Version != "0003_wellplan_import" || known[3].Version != "0004_foreign_key_indexes" || known[4].Version != "0005_neutral_import_labels" {
		t.Fatalf("unexpected migration order: %+v", versions(known))
	}
	// Applied databases record this checksum; editing 0001 would strand them.
	if !strings.HasPrefix(known[0].Checksum, "0cf7bc7641ba") {
		t.Fatalf("0001_initial.sql changed (checksum %s); add a new migration instead", known[0].Checksum[:12])
	}
}

func TestCheckApplied(t *testing.T) {
	known := []migration{{Version: "0001", Checksum: "a"}, {Version: "0002", Checksum: "b"}}
	cases := []struct {
		name    string
		applied []appliedMigration
		done    int
		wantErr string
	}{
		{"fresh", nil, 0, ""},
		{"partial", []appliedMigration{{"0001", "a"}}, 1, ""},
		{"complete", []appliedMigration{{"0001", "a"}, {"0002", "b"}}, 2, ""},
		{"edited history", []appliedMigration{{"0001", "x"}}, 0, "differs"},
		{"newer database", []appliedMigration{{"0001", "a"}, {"0002", "b"}, {"0003", "c"}}, 0, "newer binary"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			done, err := checkApplied(known, tc.applied)
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("want error containing %q, got %v", tc.wantErr, err)
				}
				return
			}
			if err != nil || done != tc.done {
				t.Fatalf("got (%d, %v), want (%d, nil)", done, err, tc.done)
			}
		})
	}
}

func versions(ms []migration) []string {
	out := make([]string, len(ms))
	for i, m := range ms {
		out[i] = m.Version
	}
	return out
}
