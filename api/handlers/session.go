package handlers

import (
	"shortener-url/api/http"
	"shortener-url/structs/v1"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login
// @Router /v1/login [POST]
// @Summary Login
// @Description Login
// @Tags Session
// @Accept json
// @Produce json
// @Param login body v1.Login true "LoginRequestBody"
// @Success 201 {object} http.Response{data=v1.LoginResponse} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 401 {object} http.Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {
	var login v1.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.srvs.Session().Login(
		c.Request.Context(),
		&login,
	)

	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}
