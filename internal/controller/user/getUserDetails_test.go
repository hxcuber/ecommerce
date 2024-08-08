package user

import (
	"context"
	"errors"
	"github.com/hxcuber/ecommerce/internal/repository"
	"github.com/hxcuber/ecommerce/internal/repository/user"
	"github.com/hxcuber/ecommerce/pkg/model"
	"github.com/hxcuber/ecommerce/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_impl_GetUserDetails(t *testing.T) {
	tests := map[string]struct {
		id      util.UUIDString
		expUser model.User
		expErr  error
	}{
		"success": {
			id: "success-id",
			expUser: model.User{
				UserID:       "success-id",
				Username:     "success-name",
				Email:        "success@test.com",
				PasswordHash: "passwordhash",
				CreatedAt:    time.Now(),
			},
			expErr: nil,
		},
		"failure": {
			id:      "failure-id",
			expUser: model.User{},
			expErr:  errors.New("fail"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockRepo := user.NewMockRepository(t)
			mockReg := repository.NewMockRegistry(t)
			ctx := context.TODO()
			i := impl{
				reg: mockReg,
			}
			mockReg.EXPECT().User().Return(mockRepo)
			mockRepo.EXPECT().GetUserById(context.Background(), tt.id).Return(tt.expUser.ToOrmUser(), tt.expErr)
			got, err := i.GetUserDetails(ctx, tt.id)
			assert.ErrorIs(t, tt.expErr, err)
			assert.Equal(t, tt.expUser, got)
		})
	}
}
