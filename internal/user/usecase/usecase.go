package usecase

import (
	"context"

	"github.com/test-task/internal/entity"
	"github.com/test-task/internal/user"
	"github.com/test-task/pkg/logger"
	richdata "github.com/test-task/pkg/rich-data"
)

type useCase struct{
	repo user.Repo
}

func NewUseCase(repo user.Repo)*useCase{
	return &useCase{repo: repo}
}

func (u *useCase)Get(ctx context.Context, l logger.Interface)error{
	return nil
}

func (u *useCase)Delete(ctx context.Context, l logger.Interface,data *entity.User)error{
	err:=u.repo.Delete(ctx,l,data)
	if err!=nil{
				return err
		}
	return nil
}

func (u *useCase)Update(ctx context.Context, l logger.Interface,data *entity.User)error{
	return nil
}

func (u *useCase)Create(ctx context.Context, l logger.Interface,data *entity.User)error{
	richData:=richdata.RichAPI(ctx,l,data)
	data.Age = richData.Age
	data.Gender = richData.Gender
	data.Nationality = richData.Nationality

	err:=u.repo.Create(ctx,l,data)
	if err!=nil{
			return err
	}
	
	return nil
}