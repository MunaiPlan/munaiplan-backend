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

var (
	ErrFieldNotChanged = errors.New("field was not changed")
)

var (
	ErrSiteNotChanged = errors.New("site was not changed")
)

var (
	ErrWellNotChanged = errors.New("well was not changed")
)

var (
	ErrWellboreNotChanged = errors.New("wellbore was not changed")
)

var (
	ErrDesignNotChanged = errors.New("design was not changed")
)

var (
	ErrTrajectoryNotChanged       = errors.New("trajectory was not changed")
	ErrTrajectoryHeaderIdNotFound = errors.New("trajectory header ID was not found")
	ErrTrajectoryUnitIdNotFound   = errors.New("trajectory unit ID was not found")
)
