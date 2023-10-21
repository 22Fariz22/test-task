package user

import (
	"context"
	"github.com/test-task/internal/entity"
	"github.com/test-task/pkg/logger"
)

type Repo interface{
	Get(ctx context.Context, l logger.Interface)
	Delete(ctx context.Context, l logger.Interface,id int)
	Update(ctx context.Context, l logger.Interface,id int)
	Create(ctx context.Context, l logger.Interface, data *entity.User)
}