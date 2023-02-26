package v1

import (
	"shortener-url/config"
	"shortener-url/pkg/logger"
	"shortener-url/services"
	"shortener-url/storage"
)

type Service struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI

	url     services.UrlServiceI
	session services.SessionServiceI
	user    services.UserServiceI
}

func NewService(strg storage.StorageI, cfg config.Config, log logger.LoggerI) services.ServiceI {
	return &Service{
		strg: strg,
		cfg:  cfg,
		log:  log,
	}
}

func (s *Service) Url() services.UrlServiceI {
	if s.Url == nil {
		return NewUrlService(s.strg, s.cfg, s.log)
	}
	return s.url
}

func (s *Service) Session() services.SessionServiceI {
	if s.session == nil {
		return NewSessionService(s.strg, s.cfg, s.log)
	}
	return s.session
}

func (s *Service) User() services.UserServiceI {
	if s.user == nil {
		return NewUserService(s.strg, s.cfg, s.log)
	}
	return s.user
}
