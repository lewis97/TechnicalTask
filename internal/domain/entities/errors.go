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
type AccountAlreadyExists struct {
	BaseError
}

type InvalidInputError struct {
	BaseError
}

func NewAccountNotFoundByIDError(accountID string) *AccountNotFound {
	msg := fmt.Sprintf("No account with ID %s found", accountID)
	return &AccountNotFound{
		BaseError{Msg: msg},
	}
}

func NewAccountNotFoundByDocNumError(docNum uint) *AccountNotFound {
	msg := fmt.Sprintf("No account with document number %d found", docNum)
	return &AccountNotFound{
		BaseError{Msg: msg},
	}
}

func NewAccountAlreadyExistsError(docNum uint) *AccountAlreadyExists {
	msg := fmt.Sprintf("Account already exists with document number %d", docNum)
	return &AccountAlreadyExists{
		BaseError{Msg: msg},
	}
}

func NewInvalidInputError(msg string) *InvalidInputError {
	return &InvalidInputError{
		BaseError{Msg: msg},
	}
}
