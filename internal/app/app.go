package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/test-task/internal/config"
	"github.com/test-task/internal/user"
	userRouter "github.com/test-task/internal/user/delivery/http"
	"github.com/test-task/internal/user/repository"
	"github.com/test-task/internal/user/usecase"
	"github.com/test-task/pkg/logger"
	"github.com/test-task/pkg/postgres"
)

type App struct {
	cfg        *config.Config
	httpServer *http.Server
	userUC     user.UseCase
}

func NewApp(cfg *config.Config) *App {
	databaseDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresqlUser,
		cfg.PostgresqlPassword,
		cfg.PostgresqlHost,
		cfg.PostgresqlPort,
		cfg.PostgresqlDbname,
	)

	// Repository
	db, err := postgres.New(databaseDSN, postgres.MaxPoolSize(2))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	userRepo := repository.NewPgRepository(db)

	return &App{
		cfg:        cfg,
		httpServer: nil,
		userUC:     usecase.NewUseCase(userRepo),
	}
}

func (a *App) Run() {
	l := logger.New("debug")

	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/")

	userRouter.RegisterHTTPEndpointsUser(l, api, a.userUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           a.cfg.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			l.Fatal("Failed to listen and serve: %+v", err)
		}
	}()

	//gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	a.httpServer.Shutdown(ctx)
}
