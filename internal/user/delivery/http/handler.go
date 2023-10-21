package http

import (
	"github.com/gin-gonic/gin"
	"github.com/test-task/internal/user"
	"github.com/test-task/pkg/logger"
)

type Handler struct{
	l logger.Interface
	useCase user.UseCase
}

func NewHandler(l logger.Interface,uc user.UseCase)*Handler{
	return &Handler{
		l:l,
		useCase:uc,
	}
}

func (h *Handler)Get(c *gin.Context){

}

func (h *Handler)Create(c *gin.Context){

}

func (h *Handler)Delete(c *gin.Context){

}

func (h *Handler)Update(c *gin.Context){

}