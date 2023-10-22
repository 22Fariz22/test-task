package repository

import (
	"context"
	"fmt"

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

func (r *pgRepo)Get(ctx context.Context, l logger.Interface)error{
	return nil
}

func (r *pgRepo)Delete(ctx context.Context, l logger.Interface,id int)error{
	return nil
}

func (r *pgRepo)Update(ctx context.Context, l logger.Interface,id int)error{
	return nil
}

func (r *pgRepo)Create(ctx context.Context, l logger.Interface,data *entity.User)error{
	fmt.Println("Repo - Create().")
	return nil
}