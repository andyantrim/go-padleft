package main

import (
	"errors"
)

var (
	ErrInvalidPadLength = errors.New("Invalid pad length, must be 0 or greater")
)
