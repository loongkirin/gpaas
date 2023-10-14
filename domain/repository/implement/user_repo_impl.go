package implement

import (
	"errors"
	"fmt"

	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
	util "github.com/loongkirin/gpaas/util"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	BaseRepository
}

func NewUserRepository(dbContext core.DbContext) *UserRepositoryImpl {
	userRepo := UserRepositoryImpl{
		BaseRepository: BaseRepository{
			dbContext: dbContext,
		},
	}
	return &userRepo
}

func (r *UserRepositoryImpl) FindById(id string) (*model.User, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("用户名不存在或者密码错误")
		}
		return nil, core.AsAppError(err)
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByMobile(mobile string) (*model.User, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var user model.User
	err := db.Where("mobile = ?", mobile).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("用户名不存在或者密码错误")
		}
		return nil, core.AsAppError(err)
	}
	return &user, nil
}

func (r *UserRepositoryImpl) InsertUser(u *model.User) (*model.User, *core.AppError) {
	db, appErr := r.getDb()
	if appErr == nil {
		return nil, appErr
	}
	var user model.User
	err := db.Where("mobile = ?", u.Mobile).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		u.Password = util.BcryptHash(u.Password)
		uid := util.GenerateId()
		fmt.Println(uid)
		u.DbBaseModel.Id = uid
		err = db.Create(u).Error
		if err != nil {
			return nil, core.AsAppError(err)
		}
		return u, nil
	}
	return nil, core.NewUnexpectedError("用户名已注册")
}

func (r *UserRepositoryImpl) UpdateUser(u *model.User) (*model.User, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var user model.User
	err := db.Where("mobile = ?", u.Mobile).First(&user).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	user.Name = u.Name
	user.DenyLogin = u.DenyLogin
	err = db.Save(&user).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return &user, nil
}

func (r *UserRepositoryImpl) DeleteUserByMobile(mobile string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	user := model.User{}
	err := db.Where("mobile = ?", mobile).Delete(&user).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}

func (r *UserRepositoryImpl) DeleteUserById(id string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	user := model.User{}
	err := db.Delete(&user, id).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}
