package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
)

func (i impl) GetUserByUsername(ctx context.Context, username string) (*orm.User, error) {
	var ormUser orm.User
	err := orm.Users(orm.UserWhere.Username.EQ(username)).Bind(ctx, i.dbConn, &ormUser)
	if err != nil {
		return nil, err
	}

	return &ormUser, nil
}
