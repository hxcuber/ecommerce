package user

import (
	"context"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/util"
)

func (i impl) GetUserById(ctx context.Context, id util.UUIDString) (*orm.User, error) {
	ormUser, err := orm.FindUser(ctx, i.dbConn, string(id))
	if err != nil {
		return nil, err
	}

	return ormUser, nil
}
