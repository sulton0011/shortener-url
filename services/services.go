package services

import (
	"context"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
)

type ServiceI interface {
	Url() UrlServiceI
	Session() SessionServiceI
	User() UserServiceI
}

type UrlServiceI interface {
	Create(ctx context.Context, req *structV1.CreateUrlRequest) (resp *structV1.GetUrlResponse, err error)
	GetByPK(ctx context.Context, req *structs.ById) (resp *structV1.GetUrlResponse, err error)
}

type SessionServiceI interface {
	Login(ctx context.Context, req *structV1.Login) (resp *structV1.LoginResponse, err error)
}

type UserServiceI interface {
	CreateUsers(ctx context.Context, req *structV1.CreateUser) (resp *structV1.GetUsersById, err error)
	GetUsersById(ctx context.Context, req *structs.ById) (resp *structV1.GetUsersById, err error)
	DeleteUsers(ctx context.Context, req *structs.ById) (err error)
	GetUserList(ctx context.Context, req *structs.ListRequest) (resp *structV1.GetUserListResponse, err error)
}
