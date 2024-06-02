package uuidgen

import "github.com/google/uuid"

// Create an interface for uuid generator so we can mock this and control UUIDs in unit tests

type UUIDGenerator interface {
	New() (uuid.UUID, error)
}

type GoogleUUIDGen struct{}

func NewGoogleUUIDGen() GoogleUUIDGen {
	return GoogleUUIDGen{}
}

func (g GoogleUUIDGen) New() (uuid.UUID, error) {
	return uuid.NewV7()
}
