package handlers

import (
	"shortener-url/api/http"
	"shortener-url/pkg/util"
	"shortener-url/structs"
	"shortener-url/structs/v1"

	"github.com/gin-gonic/gin"
)

// CreateUsers godoc
// @ID create_user
// @Router /v1/user [POST]
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param user body v1.CreateUser true "CreateUserBody"
// @Success 201 {object} v1.GetUsersById "GetUsersByIdBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) CreateUsers(c *gin.Context) {
	req := v1.CreateUser{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.srvs.User().CreateUsers(c.Request.Context(), &req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetUserByID godoc
// @ID get_user_by_id
// @Router /v1/user/{id} [GET]
// @Summary Get User
// @Description Get User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} v1.GetUsersById "GetUsersByIdBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUserByID(c *gin.Context) {

	id := c.Param("id")
	if !util.IsValidUUID(id) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.srvs.User().GetUsersById(c.Request.Context(), &structs.ById{Id: id})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateUser godoc
// @ID update_user_token
// @Router /v1/user-token [PUT]
// @Summary Update User token
// @Description Update User token
// @Tags Users
// @Accept json
// @Produce json
// @Param reqBody body v1.UpdateUserToken true "UpdateUserTokenBody"
// @Success 202 {object} string "Success"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	req := v1.UpdateUserToken{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.srvs.User().UpdateUser(c.Request.Context(), &req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Accepted, "Success")
}

// DeleteUser godoc
// @ID delete_user
// @Router /v1/user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 202 {object} string "Success"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	if !util.IsValidUUID(id) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	err := h.srvs.User().DeleteUsers(c.Request.Context(), &structs.ById{Id: id})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Accepted, "Success")
}

// GetUserList godoc
// @ID get_user_list
// @Router /v1/user [GET]
// @Summary Get user list
// @Description Get user list
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Success 200 {object} v1.GetUserListResponse "GetUserListResponseBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {

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

	resp, err := h.srvs.User().GetUserList(c.Request.Context(), &structs.ListRequest{
		Limit: int32(limit),
		Page:  int32(page),
	})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
