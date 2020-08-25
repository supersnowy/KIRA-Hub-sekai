package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// errors
var (
	UnknownMessageCode   = errors.Register(ModuleName, 001, "UnknownMessageCode")
	IncorrectMessageCode = errors.Register(ModuleName, 002, "IncorrectMessageCode")
	UnknownQueryCode     = errors.Register(ModuleName, 101, "UnknownQueryCode")
	IncorrectQueryCode   = errors.Register(ModuleName, 102, "IncorrectQueryCode")
	EntityNotFoundCode   = errors.Register(ModuleName, 103, "EntityNotFoundCode")
)
