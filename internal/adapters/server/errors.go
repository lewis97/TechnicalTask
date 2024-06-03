package server

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

var UnimplementedErr = huma.Error501NotImplemented("method is not implemented yet")

func DomainToRESTError(dErr error) huma.StatusError {
	errMsg := dErr.Error()

	switch err := dErr.(type) {
	case *entities.InvalidInputError:
		return huma.Error400BadRequest(err.Error())
	case *entities.AccountNotFound:
		return huma.Error404NotFound(errMsg)
	case *entities.AccountAlreadyExists:
		return huma.Error409Conflict(errMsg)
	default:
		// When we can't match to a domain error, return a generic 500 response.
		// Purposely omitting the error message to avoid exposing any internal implementation details in the response
		return huma.Error500InternalServerError("Unknown error occurred, please contact support.")
	}
}
