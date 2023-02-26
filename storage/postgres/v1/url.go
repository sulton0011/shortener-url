package v1

import (
	"context"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/logger"
	"shortener-url/storage"
	structV1 "shortener-url/structs/v1"

	"github.com/jackc/pgx/v4/pgxpool"
)

type urlRepo struct {
	db  *pgxpool.Pool
	err errors.ErrorStorage
}

func NewUrlRepo(db *pgxpool.Pool, log logger.LoggerI) storage.UrlRepoI {
	return &urlRepo{
		db:  db,
		err: *errors.NewErrorStorage(log, config.StrgUrl),
	}
}

func (r *urlRepo) Create(ctx context.Context, req *structV1.CreateUrlRequest) (resp *structV1.GetUrlResponse, err error) {
	defer r.err.Wrap(&err, "Create", req)
	query := `INSERT INTO urls() VALUES ();`
	_, err = r.db.Exec(ctx, query)
	if err != nil {
		return nil, err
	}
	// return r.GetByPK(ctx, &st.GetByPK{Url: req.Url})
	return &structV1.GetUrlResponse{}, nil
}
