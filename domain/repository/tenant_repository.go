package repository

import (
	"context"

	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
)

type TenantRepository interface {
	FindById(ctx context.Context, id string) (*model.Tenant, *core.AppError)
	Insert(ctx context.Context, tenant *model.Tenant) (*model.Tenant, *core.AppError)
	Update(ctx context.Context, tenant *model.Tenant) (*model.Tenant, *core.AppError)
	DeleteById(ctx context.Context, id string) (bool, *core.AppError)
}
