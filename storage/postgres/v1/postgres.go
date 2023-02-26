package v1

import (
	"context"
	"fmt"
	"shortener-url/config"
	"shortener-url/pkg/logger"
	"shortener-url/storage"

	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db  *pgxpool.Pool
	log logger.LoggerI

	url  storage.UrlRepoI
	user storage.UsersRepoI
}
type PGXStdLogger struct {
	log logger.LoggerI
}

func (p *PGXStdLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	var (
		args       = make([]interface{}, 0, len(data)+3) // making space for arguments + level + msg
		query bool = false
	)

	args = append(args, level, msg, "WARNING!!! SLOW QUERY")

	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
		if k == "time" {
			if v.(time.Duration) > config.PostgresLogTime {
				query = true
			}
		}
	}
	if query {
		p.log.Info("SLOW _QUERY!!!", logger.Any("query", args))
	}
}

func NewPostgres(ctx context.Context, cfg config.Config, log logger.LoggerI) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	config.ConnConfig.LogLevel = pgx.LogLevelInfo
	config.ConnConfig.Logger = &PGXStdLogger{
		log: log,
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &Store{
		db:  pool,
		log: log,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Url() storage.UrlRepoI {
	if s.url == nil {
		s.url = NewUrlRepo(s.db, s.log)
	}

	return s.url
}

func (s *Store) User() storage.UsersRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db, s.log)
	}

	return s.user
}
