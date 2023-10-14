package service

import (
	dto "github.com/loongkirin/gpaas/api/dto"
	core "github.com/loongkirin/gpaas/core"
)

type UserService interface {
	Login(u *dto.LoginRequest) (r *dto.LoginResponse, err *core.AppError)
	Register(u *dto.RegisterRequest) *core.AppError
}
