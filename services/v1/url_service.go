package v1

import (
	"context"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/logger"
	"shortener-url/services"
	"shortener-url/storage"
	structV1 "shortener-url/structs/v1"
)

type UrlService struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
	err  errors.ErrorService
}

func NewUrlService(strg storage.StorageI, cfg config.Config, log logger.LoggerI) services.UrlServiceI {
	return &UrlService{
		strg: strg,
		cfg:  cfg,
		log:  log,
		err:  *errors.NewErrorService(log, config.SrvsUrl, cfg.HTTPPort),
	}
}

func (s *UrlService) Create(ctx context.Context, req *structV1.CreateUrlRequest) (resp *structV1.GetUrlResponse, err error) {
	defer s.err.Wrap(&err, "Create", req)
	resp, err = s.strg.Url().Create(ctx, req)

	return
}
