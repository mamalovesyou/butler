package context

import (
	"errors"
)

var (
	ErrUserIDNotFound   = errors.New("UserID not found")
	ErrIdentityNotFound = errors.New("Identity not found")
)
