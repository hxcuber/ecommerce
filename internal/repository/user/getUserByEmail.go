package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/model"
)

func (i impl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var ormUser orm.User
	err := orm.Users(orm.UserWhere.Email.EQ(email)).Bind(ctx, i.dbConn, &ormUser)
	if err != nil {
		return model.User{}, err
	}

	return model.FromOrmUser(&ormUser)
}
