package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
)

func (i impl) GetUserByEmail(ctx context.Context, email string) (*orm.User, error) {
	var ormUser orm.User
	err := orm.Users(orm.UserWhere.Email.EQ(email)).Bind(ctx, i.dbConn, &ormUser)
	if err != nil {
		return nil, err
	}

	return &ormUser, nil
}
