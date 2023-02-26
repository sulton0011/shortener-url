package handlers

import (
	"context"
	"shortener-url/api/http"
	"shortener-url/pkg/util"

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
