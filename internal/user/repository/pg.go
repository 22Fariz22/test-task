package repository

import (
	"context"

	"github.com/test-task/internal/entity"
	"github.com/test-task/pkg/logger"
	"github.com/test-task/pkg/postgres"
)

type pgRepo struct{
	*postgres.Postgres
}

func NewPgRepository(db *postgres.Postgres)*pgRepo{
return &pgRepo{db}
}

func (r *pgRepo)Get(ctx context.Context, l logger.Interface){}
func (r *pgRepo)Delete(ctx context.Context, l logger.Interface,id int){}
func (r *pgRepo)Update(ctx context.Context, l logger.Interface,id int){}
func (r *pgRepo)Create(ctx context.Context, l logger.Interface,data *entity.User){}