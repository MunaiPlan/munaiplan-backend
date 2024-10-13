package internal

import (
	"context"
	"errors"
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

func Run(configPath string) {
	cfg, err := configs.Init(configPath)
	if err != nil {
		logrus.Error(err)
		return
	}

	// Dependencies
	db := postgres.NewDatabase()
	if db == nil {
		logrus.Error("failed to initialize database connection")
		return
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
		logrus.Error(err)
		return
	}

	// Initializing repositories
	repos := repository.NewRepositories(db.Conn)

	// Initializing services
	services := service.NewServices(repos, jwt, helpers.GetEnv("PREDICTION_SERVICE_URL", "http://localhost:8001"))

	// Initializing middleware
	authMiddleware := middleware.NewAuthMiddleware(jwt)

	// Initializing router and handlers
	router := infrastructure.NewRouter(services, authMiddleware)

	// HTTP Server
	srv := infrastructure.NewServer(cfg, router.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logrus.Errorf("failed to stop server: %v", err)
	}

	sqlDB, err := db.Conn.DB()
	if err != nil {
		logrus.Error(err.Error())
	}

	if err := sqlDB.Close(); err != nil {
		logrus.Error(err.Error())
	}
}
