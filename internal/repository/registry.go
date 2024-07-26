package repository

import (
	"context"
	"github.com/cenkalti/backoff/v4"
	"github.com/friendsofgo/errors"
	"github.com/hxcuber/ecommerce/internal/repository/order"
	"github.com/hxcuber/ecommerce/internal/repository/orderitem"
	"github.com/hxcuber/ecommerce/internal/repository/product"
	"github.com/hxcuber/ecommerce/internal/repository/user"
	"github.com/hxcuber/ecommerce/pkg/db/pg"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"log"
	"time"
)

type Registry interface {
	User() user.Repository
	Product() product.Repository
	Order() order.Repository
	OrderItem() orderitem.Repository
	DoInTx(ctx context.Context, txFunc func(context.Context, Registry) error, overrideBackoffPolicy backoff.BackOff) error
}

type impl struct {
	dbConn        pg.BeginnerExecutor
	tx            pg.ContextExecutor
	userRepo      user.Repository
	productRepo   product.Repository
	orderRepo     order.Repository
	orderItemRepo orderitem.Repository
}

func New(dbConn pg.BeginnerExecutor) Registry {
	return impl{
		dbConn:        dbConn,
		userRepo:      user.NewRepo(dbConn),
		productRepo:   product.NewRepo(dbConn),
		orderRepo:     order.NewRepo(dbConn),
		orderItemRepo: orderitem.NewRepo(dbConn),
	}
}

func (i impl) User() user.Repository {
	return i.userRepo
}

func (i impl) Product() product.Repository {
	return i.productRepo
}

func (i impl) Order() order.Repository {
	return i.orderRepo
}

func (i impl) OrderItem() orderitem.Repository {
	return i.orderItemRepo
}

func (i impl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error {
	if i.tx != nil {
		log.Printf(logerr.Message("DoInTx", "doing in transaction", errNestedTx))
		return errors.WithStack(errNestedTx)
	}

	if overrideBackoffPolicy == nil {
		overrideBackoffPolicy = pg.ExponentialBackOff(3, time.Minute)
	}

	return pg.TxWithBackOff(ctx, overrideBackoffPolicy, i.dbConn, func(tx pg.ContextExecutor) error {
		newI := impl{
			tx:            tx,
			userRepo:      user.NewRepo(tx),
			productRepo:   product.NewRepo(tx),
			orderRepo:     order.NewRepo(tx),
			orderItemRepo: orderitem.NewRepo(tx),
		}
		return txFunc(ctx, newI)
	})
}
