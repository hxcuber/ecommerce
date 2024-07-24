package repository

import (
	"context"
	"github.com/cenkalti/backoff/v4"
	"github.com/friendsofgo/errors"
	"github.com/hxcuber/ecommerce/internal/repository/user"
	"github.com/hxcuber/ecommerce/pkg/db/pg"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"time"
)

type Registry interface {
	User() user.Repository
	DoInTx(ctx context.Context, txFunc func(context.Context, Registry) error, overrideBackoffPolicy backoff.BackOff) error
}

type impl struct {
	dbConn   pg.BeginnerExecutor
	tx       pg.ContextExecutor
	userRepo user.Repository
}

func New(dbConn pg.BeginnerExecutor) Registry {
	return impl{
		dbConn: dbConn,
	}
}

func (i impl) User() user.Repository {
	return i.userRepo
}

func (i impl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error {
	if i.tx != nil {
		logerr.LogErrMessage("DoInTx", "doing in transaction", errNestedTx)
		return errors.WithStack(errNestedTx)
	}

	if overrideBackoffPolicy == nil {
		overrideBackoffPolicy = pg.ExponentialBackOff(3, time.Minute)
	}

	return pg.TxWithBackOff(ctx, overrideBackoffPolicy, i.dbConn, func(tx pg.ContextExecutor) error {
		newI := impl{
			tx: tx,
		}
		return txFunc(ctx, newI)
	})
}
