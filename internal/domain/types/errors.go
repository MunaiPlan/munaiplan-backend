package errors

import "errors"

var (
	ErrUserNotFound          = errors.New("user doesn't exists")
	ErrUserAlreadyExists     = errors.New("user with such email already exists")
	ErrUserPasswordIncorrect = errors.New("password incorrect")
	// ErrInvalidCredentials deliberately does not say whether the email exists.
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailTaken         = errors.New("an account or organization with this email already exists")
	ErrForbidden          = errors.New("administrator access required")
)

var (
	ErrOrganizationNotFound = errors.New("organization not found")
)


var (
	ErrCompanyNotFound = errors.New("company doesn't exists")
	ErrCompanyAlreadyExists = errors.New("company with such name already exists")
	ErrCompanyWasNotUpdated = errors.New("company was not updated")
	ErrCompanyWasNotDeleted = errors.New("company was not deleted")
)
