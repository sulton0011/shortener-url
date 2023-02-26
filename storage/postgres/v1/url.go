package v1

import (
	"context"
	"database/sql"
	"fmt"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/helper"
	"shortener-url/pkg/logger"
	"shortener-url/storage"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
	v1 "shortener-url/structs/v1"
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
	query := `
	select
		id,
		title,
		long_url,
		short_url,
		expires_count,
		created_by,
		expires_at,
		created_at,
		updated_at,
		used_count
	from urls where id = $1`
	fmt.Println(query)
	var (
		ExpiresAt sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString
	)
	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Title,
		&resp.LongUrl,
		&resp.ShortUrl,
		&resp.ExpiresCount,
		&resp.CreatedBy,
		&ExpiresAt,
		&CreatedAt,
		&UpdatedAt,
		&resp.UsedCount,
	)
	if err != nil {
		return nil, errors.Wrap(err, "r.db.QueryRow")
	}

	if ExpiresAt.Valid {
		resp.ExpiresAt = ExpiresAt.String
	}
	if CreatedAt.Valid {
		resp.CreatedAt = CreatedAt.String
	}
	if UpdatedAt.Valid {
		resp.UpdatedAt = UpdatedAt.String
	}
	return resp, nil
}

func (r *urlRepo) GetByShort(ctx context.Context, req *structs.ShortUrl) (resp *structV1.GetUrlResponse, err error) {
	defer r.err.Wrap(&err, "GetByShort", req)
	resp = &structV1.GetUrlResponse{}
	query := `
	select
		id,
		title,
		long_url,
		short_url,
		expires_count,
		created_by,
		expires_at,
		created_at,
		updated_at,
		used_count
	from urls 
	where short_url = $1 AND (used_count <= expires_count OR CURRENT_TIMESTAMP < expires_at)`
	fmt.Println(query)
	var (
		ExpiresAt sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString
	)
	err = r.db.QueryRow(ctx, query, req.ShortUrl).Scan(
		&resp.Id,
		&resp.Title,
		&resp.LongUrl,
		&resp.ShortUrl,
		&resp.ExpiresCount,
		&resp.CreatedBy,
		&ExpiresAt,
		&CreatedAt,
		&UpdatedAt,
		&resp.UsedCount,
	)
	if err != nil {
		return nil, errors.Wrap(err, "r.db.QueryRow")
	}

	if ExpiresAt.Valid {
		resp.ExpiresAt = ExpiresAt.String
	}
	if CreatedAt.Valid {
		resp.CreatedAt = CreatedAt.String
	}
	if UpdatedAt.Valid {
		resp.UpdatedAt = UpdatedAt.String
	}
	return resp, nil
}

func (r *urlRepo) Update(ctx context.Context, req *v1.UpdateUrlRequest) (resp *v1.Message, err error) {
	defer r.err.Wrap(&err, "Update", req)

	query := `
	UPDATE urls
	SET
		title=$1,
		short_url=$2,
	   	updated_at=CURRENT_TIMESTAMP,
		expires_count=$3,
		expires_at=$4,
		used_count=$5,
   	WHERE id=$6`
	_, err = r.db.Exec(ctx, query,
		req.Title,
		req.ShortUrl,
		req.ExpiresCount,
		req.ExpiresAt,
		req.UsedCount,
		req.Id,
	)
	if err != nil {
		return nil, errors.Wrap(err, "r.db.Exec")
	}
	return &v1.Message{Message: "Success!"}, nil
}

func (r *urlRepo) GetAll(ctx context.Context, req *structs.ListRequest) (resp *structV1.GetUrlListResponse, err error) {
	defer r.err.Wrap(&err, "GetAll", req)
	var (
		ExpiresAt sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString

		params = make(map[string]interface{})
	)

	resp = &structV1.GetUrlListResponse{}
	query := `
	select
		id,
		title,
		long_url,
		short_url,
		expires_count,
		created_by,
		expires_at,
		created_at,
		updated_at
	from urls where created_by = :user_id`
	order := " ORDER BY created_at DESC"
	limit := " LIMIT 10"
	offset := " OFFSET 0"

	if req.Page > 0 {
		req.Page = (req.Page - 1) * req.Limit
		params["offset"] = req.Page
		offset = ` offset :offset`
	}
	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = ` limit :limit`
	}
	if len(req.Id) > 0 {
		params["user_id"] = req.Id
	}

	q := query + order + offset + limit
	q, arr := helper.ReplaceQueryParams(q, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, errors.Wrap(err, "r.db.Query")
	}

	for rows.Next() {
		url := &structV1.GetUrlResponse{}

		err := rows.Scan(
			&url.Id,
			&url.Title,
			&url.LongUrl,
			&url.ShortUrl,
			&url.ExpiresCount,
			&url.CreatedBy,
			&ExpiresAt,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "r.db.Scan")
		}

		if ExpiresAt.Valid {
			url.ExpiresAt = ExpiresAt.String
		}
		if CreatedAt.Valid {
			url.CreatedAt = CreatedAt.String
		}
		if UpdatedAt.Valid {
			url.UpdatedAt = UpdatedAt.String
		}

		resp.Urls = append(resp.Urls, *url)
	}

	return resp, nil
}
