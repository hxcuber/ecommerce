package user

import (
	"github.com/friendsofgo/errors"
)

var (
	ErrEmailInUse    = errors.New("email already in use")
	ErrUsernameInUse = errors.New("username already in use")
)
