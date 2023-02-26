package v1

import (
	"context"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/helper"
	"shortener-url/pkg/logger"
	"shortener-url/storage"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
	"time"

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
	resp = &structV1.GetUrlResponse{}
	query := `INSERT INTO urls(
		title,
		long_url,
		short_url,
		expires_at,
		expires_count,
		created_by) VALUES ($1, $2, $3, $4, $5, $6) returning id;`
	var id string
	err = r.db.QueryRow(ctx, query,
		req.Title,
		req.LongUrl,
		req.GetShortUrl(),
		helper.DefaultValue(req.ExpiresAt, time.Now().Add(config.ExpiresDateUrl).Format(config.DatabaseTimeLayout)),
		req.ExpiresCount,
		ctx.Value("user_id"),
	).Scan(&id)
	if err != nil {
		return nil, errors.Wrap(err, "r.db.QueryRow")
	}

	return r.GetByPK(ctx, &structs.ById{Id: id})
}

func (r *urlRepo) GetByPK(ctx context.Context, req *structs.ById) (resp *structV1.GetUrlResponse, err error) {
	defer r.err.Wrap(&err, "GetByPK", req)
	resp = &structV1.GetUrlResponse{}
	query := `INSERT INTO urls(
		id,
		title,
		long_url,
		short_url,
		expires_at,
		expires_count,
		created_by) VALUES ($1, $2, $3, $4, $5, $6, $7);`
	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Title,
		&resp.LongUrl,
		&resp.ShortUrl,
		&resp.ExpiresAt,
		&resp.ExpiresCount,
		&resp.CreatedBy,
	)
	if err != nil {
		return nil, errors.Wrap(err, "r.db.QueryRow")
	}

	return resp, nil
}
