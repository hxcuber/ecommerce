package repository

import (
	"github.com/friendsofgo/errors"
)

var (
	errNestedTx = errors.New("db txn nested in db txn")
)
