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

func (p *pgRepo)Create(ctx context.Context, l logger.Interface,data *entity.User)error{
	fmt.Println("Repo - Create().")
	fmt.Println("Repo - Create() - data: ",data)

	_, err := p.Pool.Exec(ctx, `insert into person (user_id, name, surname, patronymic, age, gender, nationality) values($1, $2, $3, $4, $5, $6, $7)`,
		data.UserID, data.Name, data.Surname,data.Patronymic,data.Age,data.Gender,data.Nationality)
	if err != nil {
		l.Error("error Save in PG:", err)
		return err
	}

	return nil
}
	
