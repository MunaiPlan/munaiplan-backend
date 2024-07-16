package types

import "errors"

var (
	ErrInvalidUUID                         = errors.New("invalid UUID")
	ErrInvalidCompanyIDQueryParameter      = errors.New("invalid company ID query parameter")
	ErrInvalidOrganizationIDQueryParameter = errors.New("invalid organization ID query parameter")
	ErrInvalidIDQueryParameter             = errors.New("invalid ID query parameter")
	ErrInvalidInputBody                    = errors.New("invalid input body")
)
