package entity

import "errors"

// ErrNoTicker - no ticker symbol given
var ErrNoTicker = errors.New("no ticker symbol given")

// ErrInvalidEntity - invalid entity
var ErrInvalidEntity = errors.New("invalid entity parameters")

// ErrTwitterApiBadRequest - bad request
var ErrTwitterApiBadRequest = errors.New("bad request for Twitter API")
