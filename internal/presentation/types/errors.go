package types

import "errors"

var (
	ErrInvalidUUID             = errors.New("invalid UUID")
	ErrInvalidIDQueryParameter = errors.New("invalid ID query parameter")
	ErrInvalidInputBody        = errors.New("invalid input body")
)

var (
	ErrOrganizationsNotFound               = errors.New("organizations not found")
	ErrInvalidOrganizationIDQueryParameter = errors.New("invalid organization ID query parameter")
	ErrCreatingOrganization                = errors.New("error creating organization")
)

var (
	ErrInvalidCompanyIDQueryParameter = errors.New("invalid company ID query parameter")
)

var (
	ErrInvalidFieldIDQueryParameter = errors.New("invalid field ID query parameter")
)

var (
	ErrInvalidSiteIDQueryParameter = errors.New("invalid site ID query parameter")
)

var (
	ErrInvalidWellIDQueryParameter = errors.New("invalid well ID query parameter")
)

var (
	ErrInvalidWellboreIDQueryParameter = errors.New("invalid wellbore ID query parameter")
)