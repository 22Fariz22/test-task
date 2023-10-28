package http

import (
	"github.com/gin-gonic/gin"
	"github.com/test-task/internal/user"
	"github.com/test-task/pkg/logger"
)

func RegisterHTTPEndpointsUser(
	l logger.Interface,
	router *gin.RouterGroup,
	uc user.UseCase) {
	h := NewHandler(l, uc)

	user := router.Group("/")
	{
		user.POST("/", h.Create)
		user.GET("/", h.Get)
		user.DELETE("/", h.Delete)
		user.PUT("/", h.Update)
	}
}
