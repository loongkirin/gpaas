package repository

import (
	"context"

	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
)

type OAuthSessionRepository interface {
	FindById(ctx context.Context, id string) (*model.OAuthSession, *core.AppError)
	Insert(ctx context.Context, session *model.OAuthSession) (*model.OAuthSession, *core.AppError)
	Update(ctx context.Context, session *model.OAuthSession) (*model.OAuthSession, *core.AppError)
	DeleteById(ctx context.Context, id string) (bool, *core.AppError)
}
