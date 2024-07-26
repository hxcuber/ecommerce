package main

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/hxcuber/ecommerce/cmd/router"
	"github.com/hxcuber/ecommerce/internal/controller/product"
	"github.com/hxcuber/ecommerce/internal/controller/user"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/pkg/db/pg"
	"github.com/hxcuber/ecommerce/pkg/env"
	"github.com/hxcuber/ecommerce/pkg/httpserv"
	"log"
	"strconv"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Exiting...")
}

func run(ctx context.Context) error {
	log.Println("Starting app initialization")

	dbOpenConns, err := strconv.Atoi(env.GetAndValidateF("DB_POOL_MAX_OPEN_CONNS"))
	if err != nil {
		return errors.WithStack(fmt.Errorf("invalid db pool max open conns: %w", err))
	}
	dbIdleConns, err := strconv.Atoi(env.GetAndValidateF("DB_POOL_MAX_IDLE_CONNS"))
	if err != nil {
		return errors.WithStack(fmt.Errorf("invalid db pool max idle conns: %w", err))
	}

	conn, err := pg.NewPool(env.GetAndValidateF("DATABASE_URL"), dbOpenConns, dbIdleConns)

	if err != nil {
		return err
	}

	defer conn.Close()

	rtr, _ := initRouter(ctx, conn)

	log.Println("App initialization completed")

	httpserv.NewServer(rtr.Handler()).Start(ctx)

	return nil
}

func initRouter(
	ctx context.Context,
	dbConn pg.BeginnerExecutor) (router.Router, error) {
	registry := repository.New(dbConn)
	return router.New(
		ctx,
		user.NewController(registry),
		product.NewController(registry),
	), nil
}
