package implement

import (
	"context"
	"errors"

	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
	repo "github.com/loongkirin/gpaas/domain/repository"
	"gorm.io/gorm"
)

type TenantRepositoryImpl struct {
	BaseRepository
}

func NewTenantRepository(dbContext core.DbContext) repo.TenantRepository {
	tenantRepo := &TenantRepositoryImpl{
		BaseRepository: BaseRepository{
			dbContext: dbContext,
		},
	}
	return tenantRepo
}

func (r *TenantRepositoryImpl) FindById(ctx context.Context, id string) (*model.Tenant, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.Tenant
	err := db.WithContext(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("Data Recrod not found")
		}
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *TenantRepositoryImpl) Insert(ctx context.Context, d *model.Tenant) (*model.Tenant, *core.AppError) {
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

func (r *TenantRepositoryImpl) Update(ctx context.Context, d *model.Tenant) (*model.Tenant, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.Tenant
	err := db.WithContext(ctx).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	err = db.WithContext(ctx).Save(&d).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return d, nil
}

func (r *TenantRepositoryImpl) DeleteById(ctx context.Context, id string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	data := model.Tenant{}
	err := db.WithContext(ctx).Delete(&data, id).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}
