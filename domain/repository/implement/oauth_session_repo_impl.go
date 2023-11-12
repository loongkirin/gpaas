package implement

import (
	"context"
	"errors"

	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
	repo "github.com/loongkirin/gpaas/domain/repository"
	"gorm.io/gorm"
)

type OAuthSessionRepositoryImpl struct {
	BaseRepository
}

func NewOAuthSessionRepository(dbContext core.DbContext) repo.OAuthSessionRepository {
	seesionRepo := &OAuthSessionRepositoryImpl{
		BaseRepository: BaseRepository{
			dbContext: dbContext,
		},
	}
	return seesionRepo
}

func (r *OAuthSessionRepositoryImpl) FindById(ctx context.Context, id string) (*model.OAuthSession, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.OAuthSession
	err := db.WithContext(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("Data Recrod not found")
		}
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *OAuthSessionRepositoryImpl) Insert(ctx context.Context, d *model.OAuthSession) (*model.OAuthSession, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	err := db.WithContext(ctx).Create(d).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return d, nil
}

func (r *OAuthSessionRepositoryImpl) Update(ctx context.Context, d *model.OAuthSession) (*model.OAuthSession, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.OAuthSession
	err := db.WithContext(ctx).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	data.IsBlocked = d.IsBlocked
	err = db.WithContext(ctx).Save(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *OAuthSessionRepositoryImpl) DeleteById(ctx context.Context, id string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	data := model.OAuthSession{}
	err := db.WithContext(ctx).Delete(&data, id).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}
