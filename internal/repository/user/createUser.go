package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateUser(ctx context.Context, user model.User) (*orm.User, error) {
	ormUser := user.ToOrmUser()
	err := ormUser.Insert(ctx, i.dbConn, boil.Blacklist())
	if err != nil {
		return nil, err
	}

	return ormUser, nil
}
