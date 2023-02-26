package v1

import (
	"context"
	"os"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/logger"
	"shortener-url/storage"
	v1 "shortener-url/storage/postgres/v1"

	"testing"
)

var (
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
	err  errors.ErrorService
)

func TestMain(m *testing.M) {
	cfg = config.Load()
	log = logger.NewLogger(cfg.ServiceName, logger.LevelDebug)
	defer logger.Cleanup(log)
	err = *errors.NewErrorService(log, config.SrvsUrl, cfg.HTTPPort)

	strg, _ = v1.NewPostgres(context.Background(), cfg, log)

	os.Exit(m.Run())
}
