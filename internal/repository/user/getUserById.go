package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/model"
)

func (i impl) GetUserById(ctx context.Context, id uuid.UUID) (model.User, error) {
	ormUser, err := orm.FindUser(ctx, i.dbConn, id.String())
	if err != nil {
		return model.User{}, err
	}

	return model.FromOrmUser(ormUser)
}
