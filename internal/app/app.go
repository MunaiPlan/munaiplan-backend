package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/app/config"
	delivery "github.com/munaiplan/munaiplan-backend/internal/app/delivery/http"
	"github.com/munaiplan/munaiplan-backend/internal/app/repository"
	"github.com/munaiplan/munaiplan-backend/internal/app/service"
	"github.com/munaiplan/munaiplan-backend/pkg/auth"
	mongo "github.com/munaiplan/munaiplan-backend/pkg/database/mongodb"
	"github.com/munaiplan/munaiplan-backend/pkg/hash"
	"github.com/munaiplan/munaiplan-backend/pkg/logger"
	"github.com/munaiplan/munaiplan-backend/internal/app/server"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)

		return
	}
	fmt.Println(cfg.Mongo.URI)
	fmt.Println(cfg.Mongo.User)
	fmt.Println(cfg.Mongo.Password)

	// Dependencies
	mongoClient, err := mongo.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		logger.Error(err)

		return
	}
	fmt.Println("client est")

	db := mongoClient.Database(cfg.Mongo.Name)

	fmt.Println(db.Name())

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	fmt.Println(cfg.Auth.JWT.SigningKey)
	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)

		return
	}

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Deps{
		Repos:                  repos,
		Hasher:                 hasher,
		TokenManager:           tokenManager,
		AccessTokenTTL:         cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL:        cfg.Auth.JWT.RefreshTokenTTL,
		Environment:            cfg.Environment,
	})

	handlers := delivery.NewHandler(services, tokenManager)

		// HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}
}