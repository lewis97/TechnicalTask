package entities

import "fmt"

type BaseError struct {
	error
	Msg string
}

func (e *BaseError) Error() string {
	return e.Msg
}

// Custom domain specific errors are defined below

type AccountNotFound struct {
	BaseError
}

type InvalidInputError struct {
	BaseError
}

func NewAccountNotFoundError(accountID string) *AccountNotFound {
	msg := fmt.Sprintf("No account with ID %s found", accountID)
	return &AccountNotFound{
		BaseError{Msg: msg},
	}
}

func NewInvalidInputError(msg string) *InvalidInputError {
	return &InvalidInputError{
		BaseError{Msg: msg},
	}
}
