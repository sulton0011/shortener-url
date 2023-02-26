package storage

import (
	"context"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
)

type StorageI interface {
	CloseDB()
	Url() UrlRepoI
	User() UsersRepoI
}

type UrlRepoI interface {
	Create(ctx context.Context, req *structV1.CreateUrlRequest) (resp *structV1.GetUrlResponse, err error)
	GetByPK(ctx context.Context, req *structs.ById) (resp *structV1.GetUrlResponse, err error)
}

type UsersRepoI interface {
	CreateUsers(ctx context.Context, req *structV1.CreateUser) (resp *structV1.GetUsersById, err error)
	GetUsersById(ctx context.Context, req *structs.ById) (resp *structV1.GetUsersById, err error)
	DeleteUsers(ctx context.Context, req *structs.ById) (err error)
	GetUserList(ctx context.Context, req *structs.ListRequest) (resp *structV1.GetUserListResponse, err error)
	GetByLogin(ctx context.Context, login string) (resp *structV1.GetUsersByLogin, err error)
}
