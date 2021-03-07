package entity

import "errors"

// ErrNoTicker - no ticker symbol given
var ErrNoTicker = errors.New("No ticker symbol given")

// ErrInvalidEntity - invalid entity
var ErrInvalidEntity = errors.New("Invalid entity parameters")
