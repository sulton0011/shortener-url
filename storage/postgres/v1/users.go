package v1

import (
	"context"
	"database/sql"
	"fmt"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/helper"
	"shortener-url/pkg/logger"
	"shortener-url/pkg/util"
	"shortener-url/storage"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db  *pgxpool.Pool
	err errors.ErrorStorage
}

func NewUserRepo(db *pgxpool.Pool, log logger.LoggerI) storage.UsersRepoI {
	return &UserRepo{
		db:  db,
		err: *errors.NewErrorStorage(log, config.StrgUser),
	}
}

func (r *UserRepo) CreateUsers(ctx context.Context, req *structV1.CreateUser) (resp *structV1.GetUsersById, err error) {
	defer r.err.Wrap(&err, "CreateUsers", req)
	var id string
	resp = &structV1.GetUsersById{}
	query := `INSERT INTO users(
		name,
		surname,
		middle_name,
		email,
		login,
		password) 
		VALUES ($1, $2, $3, $4, $5, $6) returning id;`
	err = r.db.QueryRow(ctx, query,
		req.Name,
		req.Surname,
		req.MiddleName,
		req.Email,
		req.Login,
		req.Password,
	).Scan(
		&id,
	)
	if err != nil {
		return nil, err
	}
	resp.Id = id
	return
}

func (r *UserRepo) GetUsersById(ctx context.Context, req *structs.ById) (resp *structV1.GetUsersById, err error) {
	defer r.err.Wrap(&err, "GetUsersById", req)
	resp = &structV1.GetUsersById{}

	if !util.IsValidUUID(req.Id) {
		return resp, errors.New("id not  valid UUID")
	}

	query := `SELECT 
		id,
		name,
		surname,
		middle_name,
		email,
		created_at,
		updated_at
	FROM users WHERE id = $1`
	var (
		createdAt, updatedAt sql.NullString
	)
	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Surname,
		&resp.MiddleName,
		&resp.Email,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return
	}
	if createdAt.Valid {
		resp.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		resp.UpdatedAt = updatedAt.String
	}
	return resp, nil
}

func (r *UserRepo) DeleteUsers(ctx context.Context, req *structs.ById) (err error) {
	defer r.err.Wrap(&err, "DeleteUsers", req)
	row, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	if err != nil {
		return
	}
	if row.RowsAffected() == 0 {
		return errors.New("no rows were affected")
	}
	return
}

func (r *UserRepo) GetUserList(ctx context.Context, req *structs.ListRequest) (resp *structV1.GetUserListResponse, err error) {
	defer r.err.Wrap(&err, "GetUserList", req)
	resp = &structV1.GetUserListResponse{}

	var (
		params = make(map[string]interface{})
	)

	query := `SELECT 
		id,
		name,
		surname,
		middle_name,
		email,
		created_at,
		updated_at
	FROM "users"`
	// filter := " WHERE 1 = 1"
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
		fmt.Println(req.Page)
		fmt.Println(req.Limit)
		limit = ` limit :limit`
	}

	err = r.db.QueryRow(ctx, "select count(1) from users").Scan(
		&resp.Count,
	)
	if err != nil {
		return nil, errors.Wrap(err, "r.db.QueryRow")
	}

	q := query + order + offset + limit

	q, arr := helper.ReplaceQueryParams(q, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return resp, errors.Wrap(err, "r.db.Query")
	}
	defer rows.Close()

	for rows.Next() {
		user := &structV1.GetUsersById{}

		var (
			createdAt, updatedAt sql.NullString
		)
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Surname,
			&user.MiddleName,
			&user.Email,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return resp, errors.Wrap(err, "rows.Scan")
		}
		if createdAt.Valid {
			user.CreatedAt = createdAt.String
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.String
		}

		resp.Users = append(resp.Users, *user)
	}
	return
}
