package errors

import "errors"

var (
	ErrUserNotFound          = errors.New("user doesn't exists")
	ErrUserAlreadyExists     = errors.New("user with such email already exists")
	ErrUserPasswordIncorrect = errors.New("password incorrect")
)


var (
	ErrCompanyNotFound = errors.New("company doesn't exists")
	ErrCompanyAlreadyExists = errors.New("company with such name already exists")
	ErrCompanyWasNotUpdated = errors.New("company was not updated")
	ErrCompanyWasNotDeleted = errors.New("company was not deleted")
)
