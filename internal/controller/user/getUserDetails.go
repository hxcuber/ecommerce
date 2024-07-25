package user

import (
	"context"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
	"log"
)

func (i impl) GetUserDetails(ctx context.Context, id util.UUIDString) (model.User, error) {
	user, err := i.reg.User().GetUserById(context.Background(), id)
	if err != nil {
		log.Printf(logerr.LogErrMessage("GetUserDetails", "getting user by id", err))
		return model.User{}, err
	}

	return model.FromOrmUser(user), nil
}
