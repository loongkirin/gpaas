package implement

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	dto "github.com/loongkirin/gpaas/api/dto"
	app "github.com/loongkirin/gpaas/app"
	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
	oauth "github.com/loongkirin/gpaas/domain/oauth"
	repo "github.com/loongkirin/gpaas/domain/repository"
	repoImpl "github.com/loongkirin/gpaas/domain/repository/implement"
	svc "github.com/loongkirin/gpaas/service"
	util "github.com/loongkirin/gpaas/util"
)

type UserServiceImpl struct {
	userRepo    repo.UserRepository
	sessionRepo repo.OAuthSessionRepository
	oauthMaker  oauth.OAuthMaker
}

func NewUserService() svc.UserService {
	userReop := repoImpl.NewUserRepository(app.AppContext.APP_DbContext)
	sessionRepo := repoImpl.NewOAuthSessionRepository(app.AppContext.APP_DbContext)
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)

	if err != nil {
		panic("NewUserService error")
	}

	userSvc := &UserServiceImpl{
		userRepo:    userReop,
		sessionRepo: sessionRepo,
		oauthMaker:  oauthMaker,
	}

	return userSvc
}

func (s *UserServiceImpl) Login(ctx *gin.Context, u *dto.LoginRequest) (r *dto.LoginResponse, err *core.AppError) {
	user, err := s.userRepo.FindByMobile(ctx, u.Mobile)
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

	accessToken, _, oauthErr := s.oauthMaker.GenerateAccessToken(user.Mobile, user.Name)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取access token失败")
	}
	refreshToken, claims, oauthErr := s.oauthMaker.GenerateRefreshToken(user.Mobile, user.Name)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取refresh token失败")
	}

	session := model.OAuthSession{
		UserId:       user.DbBaseModel.Id,
		Mobile:       user.Mobile,
		UserName:     user.Name,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    claims.ExpiredAt.UnixMilli(),

		DbBaseModel: core.DbBaseModel{
			Id:          claims.Id,
			TenantId:    "1",
			DataVersion: 1,
			DataStatus:  1,
		},
	}

	_, err = s.sessionRepo.Insert(ctx, &session)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert Session Data Failure")
	}

	r = &dto.LoginResponse{
		Mobile:       user.Mobile,
		UserId:       user.DbBaseModel.Id,
		UserName:     user.Name,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return r, nil
}

func (s *UserServiceImpl) Register(ctx *gin.Context, u *dto.RegisterRequest) *core.AppError {
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
		DbBaseModel: core.DbBaseModel{
			Id:          util.GenerateId(),
			TenantId:    "1",
			DataVersion: 1,
			DataStatus:  1,
		},
	}

	_, err := s.userRepo.InsertUser(ctx, &user)
	return err
}

func (s *UserServiceImpl) RefreshToken(ctx *gin.Context, u *dto.RefreshTokenRequest) (r *dto.RefreshTokenResponse, err *core.AppError) {
	claims, oauthErr := s.oauthMaker.VerifyToken(u.RefreshToken)
	if oauthErr != nil {
		return nil, core.NewValidationError("Refresh Token is Invalid")
	}

	session, err := s.sessionRepo.FindById(ctx, claims.Id)
	if err != nil {
		return nil, err
	}

	if session.IsBlocked {
		return nil, core.NewAuthenticationError("Session Blocked")
	}

	if session.UserName != claims.UserName {
		return nil, core.NewAuthenticationError("Incorrect Session User")
	}

	if session.RefreshToken != u.RefreshToken {
		return nil, core.NewAuthenticationError("Incorrect Session Refresh Token")
	}

	expiredAt := time.UnixMilli(session.ExpiredAt)
	if time.Now().After(expiredAt) {
		return nil, core.NewAuthenticationError("Session Expired")
	}

	accessToken, _, oauthErr := s.oauthMaker.GenerateAccessToken(session.Mobile, session.UserName)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取access token失败")
	}
	refreshToken, claims, oauthErr := s.oauthMaker.GenerateRefreshToken(session.Mobile, session.UserName)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取refresh token失败")
	}

	_, err = s.sessionRepo.DeleteById(ctx, session.Id)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete Session Data Failure")
	}

	session = &model.OAuthSession{
		UserId:       session.UserId,
		Mobile:       session.Mobile,
		UserName:     session.UserName,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    claims.ExpiredAt.UnixMilli(),

		DbBaseModel: core.DbBaseModel{
			Id:          claims.Id,
			TenantId:    "1",
			DataVersion: 1,
			DataStatus:  1,
		},
	}

	_, err = s.sessionRepo.Insert(ctx, session)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert Session Data Failure")
	}

	r = &dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return r, nil
}
