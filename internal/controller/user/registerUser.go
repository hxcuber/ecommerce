package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/handler/user/request"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func (i impl) RegisterUser(ctx context.Context, request request.PostUsersRegisterRequest) (model.User, error) {
	// Check unique email
	if _, err := i.reg.User().GetUserByEmail(context.Background(), request.Email); err == nil {
		log.Printf(logerr.LogErrMessage("CreateUser", "validating email uniqueness", ErrEmailInUse))
		return model.User{}, ErrEmailInUse
	}

	// Check unique username
	if _, err := i.reg.User().GetUserByUsername(context.Background(), request.Username); err == nil {
		log.Printf(logerr.LogErrMessage("CreateUser", "validating username uniqueness", ErrUsernameInUse))
		return model.User{}, ErrUsernameInUse
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf(logerr.LogErrMessage("RegisterUser", "hashing password", err))
	}

	// Create new UUID for user
	user := model.User{
		UserID:       util.UUIDString(uuid.New().String()),
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: string(passwordHash),
		CreatedAt:    time.Time{},
	}

	// Create user in db with transaction
	var returnUser *orm.User
	err = i.reg.DoInTx(context.Background(), func(ctx context.Context, registry repository.Registry) error {
		var err error
		returnUser, err = registry.User().CreateUser(ctx, user)
		if err != nil {
			log.Printf(logerr.LogErrMessage("CreateUser", "creating user", err))
			return err
		}
		return nil
	}, nil)

	if err != nil {
		log.Printf(logerr.LogErrMessage("CreateUser", "doing in transaction", err))
		return model.User{}, err
	}

	return model.FromOrmUser(returnUser), nil
}
