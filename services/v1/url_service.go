package v1

import (
	"context"
	"shortener-url/config"
	"shortener-url/pkg/errors"
	"shortener-url/pkg/logger"
	"shortener-url/services"
	"shortener-url/storage"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
	v1 "shortener-url/structs/v1"
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
	s.err.Info("Create", req)

	resp, err = s.strg.Url().Create(ctx, req)
	if err != nil {
		return
	}
	resp.GetShortUrl(s.cfg.HTTPScheme, "localhost", s.cfg.HTTPPort)
	resp.GetQrCode(resp.ShortUrl, config.SizeQrCode)

	return
}

func (s *UrlService) GetByPK(ctx context.Context, req *structs.ById) (resp *structV1.GetUrlResponse, err error) {
	defer s.err.Wrap(&err, "GetByPK", req)
	s.err.Info("GetByPK", req)

	resp, err = s.strg.Url().GetByPK(ctx, req)
	if err != nil {
		return
	}

	resp.GetShortUrl(s.cfg.HTTPScheme, "localhost", s.cfg.HTTPPort)
	resp.GetQrCode(resp.ShortUrl, config.SizeQrCode)

	return
}

func (s *UrlService) GetByShort(ctx context.Context, req *structs.ShortUrl) (resp *structV1.GetUrlResponse, err error) {
	defer s.err.Wrap(&err, "GetShort", req)
	s.err.Info("GetShort", req)

	resp, err = s.strg.Url().GetByShort(ctx, req)
	if err != nil {
		return
	}

	resp.GetShortUrl(s.cfg.HTTPScheme, "localhost", s.cfg.HTTPPort)
	resp.GetQrCode(resp.ShortUrl, config.SizeQrCode)

	return
}

func (s *UrlService) Update(ctx context.Context, req *v1.UpdateUrlRequest) (resp *v1.Message, err error) {
	defer s.err.Wrap(&err, "Update", req)
	s.err.Info("Update", req)

	resp, err = s.strg.Url().Update(ctx, req)
	if err != nil {
		return
	}

	return
}

func (s *UrlService) GetList(ctx context.Context, req *structs.ListRequest) (resp *v1.GetUrlListResponse, err error) {
	defer s.err.Wrap(&err, "GetAll", req)
	s.err.Info("GetAll", req)

	resp, err = s.strg.Url().GetAll(ctx, req)
	if err != nil {
		return
	}

	return
}
