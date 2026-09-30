package service

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/importers/wellplan"
)

type fakeAdminRepo struct {
	createdOrg  *entities.Organization
	createdUser *entities.User
}

func (f *fakeAdminRepo) ListOrganizations(context.Context) ([]*entities.OrganizationSummary, error) {
	return nil, nil
}
func (f *fakeAdminRepo) CreateOrganizationWithUser(_ context.Context, org *entities.Organization, user *entities.User) error {
	f.createdOrg, f.createdUser = org, user
	return nil
}
func (f *fakeAdminRepo) ListUsers(context.Context, string) ([]*entities.User, error) { return nil, nil }
func (f *fakeAdminRepo) CreateUser(_ context.Context, _ string, user *entities.User) error {
	f.createdUser = user
	return nil
}

func validRequest() *requests.AdminCreateOrganizationRequest {
	return &requests.AdminCreateOrganizationRequest{
		Organization: requests.AdminOrganizationInput{Name: "  Acme Drilling ", Email: " Ops@Acme.KZ "},
		User:         requests.AdminUserInput{Name: "Aida", Surname: "Nur", Email: " Aida@Acme.KZ", Password: "long-enough-password"},
	}
}

func TestAdminCreateOrganizationNormalisesAndHashes(t *testing.T) {
	repo := &fakeAdminRepo{}
	if _, _, err := NewAdminService(repo).CreateOrganization(context.Background(), validRequest()); err != nil {
		t.Fatal(err)
	}
	if repo.createdOrg.Name != "Acme Drilling" || repo.createdOrg.Email != "ops@acme.kz" {
		t.Fatalf("organization not normalised: %+v", repo.createdOrg)
	}
	u := repo.createdUser
	if u.Email != "aida@acme.kz" || u.Role != entities.RoleUser {
		t.Fatalf("user not normalised or wrong default role: %+v", u)
	}
	if u.Password == "long-enough-password" || !helpers.CheckPasswordHash("long-enough-password", u.Password) {
		t.Fatal("password must be stored as a bcrypt hash")
	}
}

func TestAdminCreateOrganizationRejectsInvalidInput(t *testing.T) {
	cases := map[string]func(*requests.AdminCreateOrganizationRequest){
		"short password": func(r *requests.AdminCreateOrganizationRequest) { r.User.Password = "short" },
		"blank org name": func(r *requests.AdminCreateOrganizationRequest) { r.Organization.Name = "   " },
		"blank surname":  func(r *requests.AdminCreateOrganizationRequest) { r.User.Surname = " " },
		"unknown role":   func(r *requests.AdminCreateOrganizationRequest) { r.User.Role = "root" },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			req := validRequest()
			mutate(req)
			repo := &fakeAdminRepo{}
			_, _, err := NewAdminService(repo).CreateOrganization(context.Background(), req)
			var v *ValidationError
			if !errors.As(err, &v) {
				t.Fatalf("want ValidationError, got %v", err)
			}
			if repo.createdOrg != nil {
				t.Fatal("nothing may be written on invalid input")
			}
		})
	}
}

func TestAdminCreateUserAllowsExplicitAdminRole(t *testing.T) {
	repo := &fakeAdminRepo{}
	in := &requests.AdminUserInput{Name: "A", Surname: "B", Email: "x@y.kz", Password: strings.Repeat("p", 12), Role: entities.RoleAdmin}
	u, err := NewAdminService(repo).CreateUser(context.Background(), "org", in)
	if err != nil || u.Role != entities.RoleAdmin {
		t.Fatalf("got %+v, %v", u, err)
	}
}

func TestFingerprintIgnoresFileOrderAndNames(t *testing.T) {
	a := fingerprintOf([]wellplan.SourceFile{{Name: "r.docx", Kind: "report", SHA256: "aa"}, {Name: "s.txt", Kind: "survey", SHA256: "bb"}})
	b := fingerprintOf([]wellplan.SourceFile{{Name: "other.txt", Kind: "survey", SHA256: "bb"}, {Name: "x.docx", Kind: "report", SHA256: "aa"}})
	c := fingerprintOf([]wellplan.SourceFile{{Name: "r.docx", Kind: "report", SHA256: "aa"}})
	if a != b || a == c {
		t.Fatalf("fingerprint must depend only on file kinds and contents")
	}
}
