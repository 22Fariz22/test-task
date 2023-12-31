package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/test-task/internal/entity"
	"github.com/test-task/internal/user"
	"github.com/test-task/pkg/logger"
)

type Handler struct {
	l       logger.Interface
	useCase user.UseCase
}

func NewHandler(l logger.Interface, uc user.UseCase) *Handler {
	return &Handler{
		l:       l,
		useCase: uc,
	}
}

func (h *Handler) Get(c *gin.Context) {
	// ctx := context.Background()
}

func (h *Handler) Create(c *gin.Context) {
	ctx := context.Background()

	in := new(entity.User)
	in.UserID = uuid.New()

	if err := c.BindJSON(in); err != nil {
		h.l.Error("failed unmarshal person entity:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.useCase.Create(ctx, h.l, in)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Delete(c *gin.Context) {
	ctx := context.Background()

	in := new(entity.User)

	if err := c.BindJSON(in); err != nil {
		h.l.Error("failed unmarshal userID entity:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.useCase.Delete(ctx, h.l, in)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Update(c *gin.Context) {
	// ctx := context.Background()

}
