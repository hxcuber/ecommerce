package model

import (
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"time"
)

type User struct {
	UserID       uuid.UUID `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Username     string    `boil:"username" json:"username" toml:"username" yaml:"username"`
	Email        string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	PasswordHash string    `boil:"password_hash" json:"password_hash" toml:"password_hash" yaml:"password_hash"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
}

func FromOrmUser(user *orm.User) (User, error) {
	id, err := uuid.Parse(user.UserID)
	if err != nil {
		return User{}, err
	}
	return User{
		UserID:       id,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}, nil
}

func (u User) ToOrmUser() *orm.User {
	return &orm.User{
		UserID:       u.UserID.String(),
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
	}
}
