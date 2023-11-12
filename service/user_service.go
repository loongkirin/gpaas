package service

import (
	"github.com/gin-gonic/gin"

	dto "github.com/loongkirin/gpaas/api/dto"
	core "github.com/loongkirin/gpaas/core"
)

type UserService interface {
	Login(ctx *gin.Context, u *dto.LoginRequest) (r *dto.LoginResponse, err *core.AppError)
	Register(ctx *gin.Context, u *dto.RegisterRequest) *core.AppError
}
