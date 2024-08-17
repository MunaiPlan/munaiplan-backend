package types

import "errors"

var (
	ErrInvalidUUID                         = errors.New("invalid UUID")
	ErrInvalidIDQueryParameter             = errors.New("invalid ID query parameter")
	ErrInvalidInputBody                    = errors.New("invalid input body")
)

var (
	ErrOrganizationsNotFound = errors.New("organizations not found")
	ErrInvalidOrganizationIDQueryParameter = errors.New("invalid organization ID query parameter")
	ErrCreatingOrganization = errors.New("error creating organization")
)

var (
	ErrInvalidCompanyIDQueryParameter      = errors.New("invalid company ID query parameter")
)