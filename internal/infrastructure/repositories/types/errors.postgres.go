package types

import "errors"

var (
	ErrUserNotFound          = errors.New("user doesn't exists")
	ErrUserAlreadyExists     = errors.New("user with such email already exists")
	ErrTransactionInvalid    = errors.New("transaction is invalid")
	ErrUserPasswordIncorrect = errors.New("password incorrect")
)

var (
	ErrCompanyNotFound      = errors.New("company doesn't exists")
	ErrCompanyAlreadyExists = errors.New("company with such name already exists")
	ErrCompanyWasNotUpdated = errors.New("company was not updated")
	ErrCompanyWasNotDeleted = errors.New("company was not deleted")
	ErrComanyNotChanged     = errors.New("company was not changed")
)

var (
	ErrOrganizationExistsWithEmail = errors.New("organization with such email already exists")
	ErrGettingOrganizationByEmail  = errors.New("error getting organization by email")
	ErrCreatingOrganization        = errors.New("error creating organization")
)
