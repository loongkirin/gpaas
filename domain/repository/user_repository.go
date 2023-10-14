package repository

import (
	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
)

type UserRepository interface {
	FindById(id string) (*model.User, *core.AppError)
	FindByMobile(mobile string) (*model.User, *core.AppError)
	InsertUser(user *model.User) (*model.User, *core.AppError)
	UpdateUser(user *model.User) (*model.User, *core.AppError)
	DeleteUserById(id string) (bool, *core.AppError)
	DeleteUserByMobile(mobile string) (bool, *core.AppError)
}
