package v1

import (
	"context"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/logger"
	"shortener-url/pkg/security"
	"shortener-url/services"
	"shortener-url/storage"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
)

type SessionService struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
	err  errors.ErrorService
}

func NewSessionService(strg storage.StorageI, cfg config.Config, log logger.LoggerI) services.SessionServiceI {
	return &SessionService{
		strg: strg,
		cfg:  cfg,
		log:  log,
		err:  *errors.NewErrorService(log, config.SrvsSession, cfg.HTTPPort),
	}
}

func (s *SessionService) Login(ctx context.Context, req *structV1.Login) (resp *structV1.LoginResponse, err error) {
	defer s.err.Wrap(&err, "Login", req)

	respUser, err := s.strg.User().GetUsersById(ctx, &structs.ById{Id: req.Id})
	if err != nil {
		return resp, err
	}

	m := map[string]interface{}{
		"id": respUser.Id,
	}

	accessToken, err := security.GenerateJWT(m, config.AccessTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {

		return nil, err
	}

	resp = &structV1.LoginResponse{
		AccessToken: accessToken,
		Id:          respUser.Id,
		CreatedAt:   respUser.CreatedAt,
		UpdatedAt:   respUser.UpdatedAt,
		Name:        respUser.Name,
		Surname:     respUser.Surname,
		MiddleName:  respUser.MiddleName,
		PhoneNumber: respUser.Email,
	}

	return
}
