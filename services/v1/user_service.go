package v1

import (
	"context"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/logger"
	"shortener-url/pkg/security"
	"shortener-url/pkg/util"
	"shortener-url/services"
	"shortener-url/storage"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
)

type UserService struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
	err  errors.ErrorService
}

func NewUserService(strg storage.StorageI, cfg config.Config, log logger.LoggerI) services.UserServiceI {
	return &UserService{
		strg: strg,
		cfg:  cfg,
		log:  log,
		err:  *errors.NewErrorService(log, config.SrvsUser, cfg.HTTPPort),
	}
}

func (s *UserService) CreateUsers(ctx context.Context, req *structV1.CreateUser) (resp *structV1.GetUsersById, err error) {
	defer s.err.Wrap(&err, "CreateUsers", req)

	if !util.IsValidEmail(req.Email) {
		return nil, errors.New("invalid email")
	}

	if len(req.Login) < 6 {
		return nil, errors.New("login must not be less than 6 characters")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("password must not be less than 6 characters")
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "error generating hash password")
	}
	req.Password = hashedPassword

	resp, err = s.strg.User().CreateUsers(ctx, req)
	if err != nil {
		return
	}
	
	return s.strg.User().GetUsersById(ctx, &structs.ById{Id: resp.Id})
}

func (s *UserService) GetUsersById(ctx context.Context, req *structs.ById) (resp *structV1.GetUsersById, err error) {
	defer s.err.Wrap(&err, "GetUsersById", req)
	resp, err = s.strg.User().GetUsersById(ctx, req)
	return
}

func (s *UserService) DeleteUsers(ctx context.Context, req *structs.ById) (err error) {
	defer s.err.Wrap(&err, "DeleteUsers", req)
	err = s.strg.User().DeleteUsers(ctx, req)
	return
}

func (s *UserService) GetUserList(ctx context.Context, req *structs.ListRequest) (resp *structV1.GetUserListResponse, err error) {
	defer s.err.Wrap(&err, "GetUserList", req)
	resp, err = s.strg.User().GetUserList(ctx, req)
	return
}
