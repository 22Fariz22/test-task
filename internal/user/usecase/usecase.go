package usecase

import (
	"context"
	"fmt"

	"github.com/test-task/internal/entity"
	"github.com/test-task/internal/user"
	"github.com/test-task/pkg/logger"
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

	func (u *useCase)Delete(ctx context.Context, l logger.Interface,id int)error{
		return nil

	}

	func (u *useCase)Update(ctx context.Context, l logger.Interface,id int)error{
		return nil

	}

	func (u *useCase)Create(ctx context.Context, l logger.Interface,data *entity.User)error{
		fmt.Println("UseCase - Create()")
		u.repo.Create(ctx,l,data)
		return nil
	}