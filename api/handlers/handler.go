package handlers

import (
	"shortener-url/api/http"
	"shortener-url/config"
	"shortener-url/pkg/logger"
	"shortener-url/pkg/security"
	"shortener-url/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	cfg  config.Config
	log  logger.LoggerI
	wd   websocket.Dialer
	wu   websocket.Upgrader
	srvs services.ServiceI
	// services client.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, srvs services.ServiceI) Handler {
	return Handler{
		cfg:  cfg,
		log:  log,
		srvs: srvs,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, data)
}

func (h *Handler) getPageParam(c *gin.Context) (page int, err error) {
	PageStr := c.DefaultQuery("page", h.cfg.DefaultPage)
	return strconv.Atoi(PageStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (limit int, err error) {
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) HasAccess(c *gin.Context) {
	reqToken := c.GetHeader("Authorization")
	if len(reqToken) < 25 {
		h.handleResponse(c, http.Forbidden, "token is empty")
		c.Abort()
		return
	}

	token, err := security.ExtractToken(reqToken)
	if err != nil {
		h.handleResponse(c, http.Forbidden, err.Error())
		c.Abort()
		return
	}

	tokenInfo, err := security.ParseClaims(token, h.cfg.SecretKey)
	if err != nil {
		h.handleResponse(c, http.Forbidden, err.Error())
		c.Abort()
		return
	}

	c.Set("id", tokenInfo.Id)

}
