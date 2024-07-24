package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"github.com/hxcuber/ecommerce/pkg/model"
)

func (i impl) GetUserDetails(ctx context.Context, id uuid.UUID) (model.User, error) {
	user, err := i.reg.User().GetUserById(context.Background(), id)
	if err != nil {
		logerr.LogErrMessage("GetUserDetails", "getting user by id", err)
		return model.User{}, err
	}

	return user, nil
}
