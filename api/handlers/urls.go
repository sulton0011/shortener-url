package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	hp "net/http"
	"shortener-url/api/http"
	"shortener-url/pkg/util"
	"time"

	"shortener-url/structs"
	v1 "shortener-url/structs/v1"

	"github.com/gin-gonic/gin"
)

// CreateUrls godoc
// @ID create_urls
// @Security ApiKeyAuth
// @Router /v1/urls [POST]
// @Summary Create Urls
// @Description Create Urls
// @Tags Urls
// @Accept json
// @Produce json
// @Param user body v1.CreateUrlRequest true "CreateUrlRequestBody"
// @Success 201 {object} v1.GetUrlResponse "GetUrlResponseBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) CreateUrl(c *gin.Context) {
	req := v1.CreateUrlRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.srvs.Url().Create(c.Value("ctx").(context.Context), &req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetUrlByID godoc
// @ID get_url_by_id
// @Security ApiKeyAuth
// @Router /v1/urls/{id} [GET]
// @Summary Get urls
// @Description Get urls
// @Tags Urls
// @Accept json
// @Produce json
// @Param id path string true "url id"
// @Success 200 {object} v1.GetUrlResponse "GetUrlResponseBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUrlByID(c *gin.Context) {

	id := c.Param("id")
	if !util.IsValidUUID(id) {
		h.handleResponse(c, http.InvalidArgument, "id is an invalid uuid")
		return
	}

	resp, err := h.srvs.Url().GetByPK(c.Value("ctx").(context.Context), &structs.ById{Id: id})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUrlByShort godoc
// @ID get_url_by_short
// @Security ApiKeyAuth
// @Router /{short_url} [GET]
// @Summary Get urls
// @Description Get urls
// @Tags Urls
// @Accept json
// @Produce json
// @Param short_url path string true "short_url"
// @Success 200 {object} v1.GetUrlResponse "GetUrlResponseBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUrlByShort(c *gin.Context) {
	var (
		isNil bool
		resp  *v1.GetUrlResponse
	)
	short_url := c.Param("short_url")

	result, err := h.clientRedis.Get(short_url).Result()
	if err != nil {
		fmt.Println(err)
		isNil = true
	}
	fmt.Println(isNil)
	if !isNil {
		err = json.Unmarshal([]byte(result), &resp)
		if err != nil {
			fmt.Println("Error while unmarshal in get link by short link", err)
			isNil = true
		}
	}
	if isNil {
		resp, err = h.srvs.Url().GetByShort(c.Request.Context(), &structs.ShortUrl{ShortUrl: short_url})
		if err != nil {
			h.handleResponse(c, http.InternalServerError, err.Error())
			return
		}

		jsonRes, err := json.Marshal(resp)
		if err == nil {
			err = h.clientRedis.Set(resp.ShortUrl, jsonRes, time.Duration(10*time.Minute)).Err()
			if err != nil {
				fmt.Println("Error while set short link data: ", err)
			}
		}
	}

	_, err = h.srvs.Url().Update(c.Request.Context(), &v1.UpdateUrlRequest{
		Id:           resp.Id,
		Title:        resp.Title,
		ShortUrl:     resp.ShortUrl,
		ExpiresAt:    resp.ExpiresAt,
		ExpiresCount: resp.ExpiresCount,
		UsedCount:    resp.UsedCount + 1,
	})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	c.Redirect(hp.StatusMovedPermanently, resp.LongUrl)
}

// UpdateUrlByID godoc
// @ID update_url_by_id
// @Security ApiKeyAuth
// @Router /v1/urls/{id} [PUT]
// @Summary Put url
// @Description Put url
// @Tags Urls
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param url body v1.UpdateUrlRequest true "UpdateUrlRequestBody"
// @Success 200 {object} v1.GetUrlResponse "GetUrlResponseBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) UpdateUrl(c *gin.Context) {
	req := v1.UpdateUrlRequest{}
	id := c.Param("id")
	if !util.IsValidUUID(id) {
		h.handleResponse(c, http.InvalidArgument, "id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	req.Id = id

	resp, err := h.srvs.Url().Update(c.Value("ctx").(context.Context), &req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUrlList godoc
// @ID get_url_list
// @Security ApiKeyAuth
// @Router /v1/urls [GET]
// @Summary Get url list
// @Description Get url list
// @Tags Urls
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Success 200 {object} v1.GetUrlListResponse "GetUrlListResponse"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUrlList(c *gin.Context) {

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	page, err := h.getPageParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.srvs.Url().GetList(c.Request.Context(), &structs.ListRequest{
		Limit: int32(limit),
		Page:  int32(page),
		Id:    c.Value("ctx").(context.Context).Value("user_id").(string),
	})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
