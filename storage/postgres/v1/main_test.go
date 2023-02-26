package v1

import (
	"context"
	"fmt"
	"os"
	"shortener-url/config"
	"shortener-url/pkg/logger"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	db  *pgxpool.Pool
	log logger.LoggerI
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	conf.MaxConns = cfg.PostgresMaxConnections

	db, err = pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		panic(err)
	}
	log := logger.NewLogger(cfg.ServiceName, logger.LevelDebug)
	defer logger.Cleanup(log)

	os.Exit(m.Run())
}
