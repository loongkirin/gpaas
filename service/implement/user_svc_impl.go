package implement

import (
	"strings"

	dto "github.com/loongkirin/gpaas/api/dto"
	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
	repo "github.com/loongkirin/gpaas/domain/repository"
	util "github.com/loongkirin/gpaas/util"
)

type UserServiceImpl struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) Login(u *dto.LoginRequest) (r *dto.LoginResponse, err *core.AppError) {
	user, err := s.userRepo.FindByMobile(u.Mobile)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, core.NewValidationError("手机号不存在")
	}

	if user.DenyLogin {
		return nil, core.NewValidationError("用户被禁止登录")
	}

	if isPass := util.BcryptCheck(u.Password, user.Password); !isPass {
		return nil, core.NewValidationError("密码错误")
	}

	r = &dto.LoginResponse{
		Mobile:      user.Mobile,
		UserId:      user.DbBaseModel.Id,
		UserName:    user.Name,
		AccessToken: "",
	}

	return r, nil
}

func (s *UserServiceImpl) Register(u *dto.RegisterRequest) *core.AppError {
	if u == nil {
		return core.NewValidationError("参数错误")
	}

	if strings.Trim(u.Mobile, " ") == "" {
		return core.NewValidationError("手机号不能为空")
	}

	if strings.Trim(u.Password, " ") == "" {
		return core.NewValidationError("密码不能为空")
	}

	user := model.User{
		Mobile:   u.Mobile,
		Password: u.Password,
		Name:     u.Name,
	}

	_, err := s.userRepo.InsertUser(&user)
	return err
}
