package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/pkg/logerr"
	"github.com/hxcuber/ecommerce/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (i impl) RegisterUser(ctx context.Context, user model.User, password string) (model.User, error) {
	// Check unique email
	if _, err := i.reg.User().GetUserByEmail(context.Background(), user.Email); err == nil {
		log.Printf(logerr.LogErrMessage("CreateUser", "validating email uniqueness", ErrEmailInUse))
		return model.User{}, ErrEmailInUse
	}

	// Check unique username
	if _, err := i.reg.User().GetUserByUsername(context.Background(), user.Username); err == nil {
		log.Printf(logerr.LogErrMessage("CreateUser", "validating username uniqueness", ErrUsernameInUse))
		return model.User{}, ErrUsernameInUse
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf(logerr.LogErrMessage("RegisterUser", "hashing password", err))
	}
	user.PasswordHash = string(passwordHash)

	// Create new UUID for user
	user.UserID = uuid.New()

	// Create user in db with transaction
	var returnUser model.User
	err = i.reg.DoInTx(context.Background(), func(ctx context.Context, registry repository.Registry) error {
		var err error
		returnUser, err = registry.User().CreateUser(ctx, user)
		if err != nil {
			logerr.LogErrMessage("CreateUser", "creating user", err)
			return err
		}
		return nil
	}, nil)

	if err != nil {
		logerr.LogErrMessage("CreateUser", "doing in transaction", err)
		return model.User{}, err
	}

	return returnUser, nil
}
