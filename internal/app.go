package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/application/service"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/configs"
	postgres "github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/connection"
	infrastructure "github.com/munaiplan/munaiplan-backend/internal/infrastructure/http"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/middleware"
	"github.com/sirupsen/logrus"
	//"github.com/xuri/excelize/v2"
)

// @title MunaiPlan API
// @version 1.0
// @description REST API endpoints for Munai Plan App

// @host localhost:8000
// @BasePath /api/v1/

func Run(command, configPath string) error {
	switch command {
	case "serve", "migrate", "dev-seed", "create-admin":
	default:
		return fmt.Errorf("unknown command %q (expected serve, migrate, dev-seed or create-admin)", command)
	}
	var cfg *configs.Config
	var err error
	if command == "serve" {
		cfg, err = configs.Init(configPath)
		if err != nil {
			return err
		}
	}
	if command == "dev-seed" && os.Getenv("APP_ENV") != configs.EnvLocal {
		return fmt.Errorf("dev-seed requires APP_ENV=local")
	}
	db, err := postgres.Open()
	if err != nil {
		return err
	}
	defer db.Close()
	if command == "migrate" {
		return postgres.Migrate(db.Conn)
	}
	if err := postgres.RequireSchema(db.Conn); err != nil {
		return err
	}
	if command == "dev-seed" {
		return postgres.SeedLocalUser(db.Conn)
	}
	if command == "create-admin" {
		return postgres.BootstrapAdmin(db.Conn)
	}

	// fmt.Println(cfg.Catalog.ApiDrillCollar)
	// file := excelize.NewFile()
	// defer func() {
	//     // Save the Excel file once all catalogs have been processed
	//     if err := file.SaveAs("data/catalog.xlsx"); err != nil {
	//         log.Fatalf("Failed to save the Excel file: %v", err)
	//     }
	// }()
	// catalog := catalog.NewCatalogCache(cfg.Catalog, file)

	jwt, err := helpers.NewJwt()
	if err != nil {
		return err
	}

	// Initializing repositories
	repos := repository.NewRepositories(db.Conn)

	// Initializing services
	services := service.NewServices(repos, jwt, helpers.GetEnv("PREDICTION_SERVICE_URL", "http://localhost:8001/predict"))

	// Initializing middleware
	authMiddleware := middleware.NewAuthMiddleware(jwt, services.Users, repos.Ownership)

	// Initializing router and handlers
	router := infrastructure.NewRouter(services, authMiddleware)

	// HTTP Server
	srv := infrastructure.NewServer(cfg, router.Init(cfg))

	serverErrors := make(chan error, 1)
	go func() { serverErrors <- srv.Run() }()

	logrus.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("HTTP server stopped: %w", err)
	case <-quit:
	}

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		return fmt.Errorf("stop HTTP server: %w", err)
	}
	return nil
}
