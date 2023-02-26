package main

import (
	"context"
	"fmt"
	"shortener-url/api"
	"shortener-url/api/handlers"
	"shortener-url/config"
	"shortener-url/pkg/logger"
	serviceV1 "shortener-url/services/v1"
	postgresV1 "shortener-url/storage/postgres/v1"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	cfg := config.Load()
	loggerLevel := logger.LevelDebug

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})

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

	pong, err := clientRedis.Ping().Result()
	if err != nil {
		log.Panic("clientRedis", logger.Error(err))
	}

	fmt.Println("RESULT_REDIS: ", pong)

	srvs := serviceV1.NewService(pgStore, cfg, log)
	h := handlers.NewHandler(cfg, log, srvs, clientRedis)
	r := api.SetUpRouter(h, cfg)

	r.Run(cfg.HTTPPort)

}
