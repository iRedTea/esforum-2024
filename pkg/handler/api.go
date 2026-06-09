package handler

import (
	"esforum/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type threadRequest struct {
	Title string `json:"title"`
}

// thread        godoc
// @Summary      Create new thread (only for admins)
// @Tags         thread
// @Produce      json
// @Param        thread  body   handler.threadRequest  true       "Thread data in JSON"
// @Param        Authorization   header    string        true   	"JWT Bearer token (authorization)"
// @Success      200 {object} esforum.Thread "created thread"
// @Router       /api/thread [post]
func (h *Handler) create(c *gin.Context) {
	var input threadRequest

	if !util.MatchRole(c, "admin") {
		util.NewErrorResponse(c, http.StatusForbidden, "no perms")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	thread, err := h.repo.CreateThread(input.Title)
	if err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, thread)
}

// thread/delete          godoc
// @Summary               Delete thread by id (only for admins)
// @Tags                  thread
// @Produce               json
// @Param                 id  path      int  true  "id of thread"
// @Param                 Authorization    header    string    true   	"JWT Bearer token (authorization)"
// @Success               200
// @Router                /api/thread/{id} [delete]
func (h *Handler) delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	_, err := h.repo.GetThread(id)
	if err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	current, _ := c.Get(userCtx)
	currentUser, err := h.services.AuthorizeById(current.(int64))
	if (err != nil) || (currentUser.Access != "admin") {
		util.NewErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	err = h.repo.DeleteThread(id)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "{ \"status\": \"ok\" }")
}
