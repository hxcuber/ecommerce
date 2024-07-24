package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/model"
)

func (i impl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var ormUser orm.User
	err := orm.Users(orm.UserWhere.Username.EQ(username)).Bind(ctx, i.dbConn, &ormUser)
	if err != nil {
		return model.User{}, err
	}

	return model.FromOrmUser(&ormUser)
}
