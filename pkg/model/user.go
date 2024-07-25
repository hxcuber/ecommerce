package model

import (
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"github.com/hxcuber/ecommerce/pkg/util"
	"time"
)

type User struct {
	UserID       util.UUIDString `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Username     string          `boil:"username" json:"username" toml:"username" yaml:"username"`
	Email        string          `boil:"email" json:"email" toml:"email" yaml:"email"`
	PasswordHash string          `boil:"password_hash" json:"password_hash" toml:"password_hash" yaml:"password_hash"`
	CreatedAt    time.Time       `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
}

func FromOrmUser(user *orm.User) User {
	return User{
		UserID:       util.UUIDString(user.UserID),
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}
}

func (u User) ToOrmUser() *orm.User {
	return &orm.User{
		UserID:       string(u.UserID),
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
	}
}
