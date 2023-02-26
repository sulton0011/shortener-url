package main

import (
	"context"
	"shortener-url/api"
	"shortener-url/api/handlers"
	"shortener-url/config"
	"shortener-url/pkg/logger"
	serviceV1 "shortener-url/services/v1"
	postgresV1 "shortener-url/storage/postgres/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}
	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	pgStore, err := postgresV1.NewPostgres(context.Background(), cfg, log)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	srvs := serviceV1.NewService(pgStore, cfg, log)
	h := handlers.NewHandler(cfg, log, srvs)
	r := api.SetUpRouter(h, cfg)

	r.Run(cfg.HTTPPort)

}
