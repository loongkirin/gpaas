package repository

import (
	"context"

	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
)

type UserRepository interface {
	FindById(ctx context.Context, id string) (*model.User, *core.AppError)
	FindByMobile(ctx context.Context, mobile string) (*model.User, *core.AppError)
	InsertUser(ctx context.Context, user *model.User) (*model.User, *core.AppError)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, *core.AppError)
	DeleteUserById(ctx context.Context, id string) (bool, *core.AppError)
	DeleteUserByMobile(ctx context.Context, mobile string) (bool, *core.AppError)
}
