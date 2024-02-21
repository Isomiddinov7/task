package handler

import (
	"context"
	"database/sql"
	"net/http"
	"task/api/models"
	"task/config"
	"task/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err: "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	resp, err := h.strg.User().Create(ctx, &createUser)

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

func (h *Handler) GetByIDUser(c *gin.Context) {
	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.User().GetByID(ctx, &models.UserPrimaryKey{Id: id})
	if err == sql.ErrNoRows {
		handleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) GetListUser(c *gin.Context) {

	limit, err := getIntegerOrDefaultValue(c.Query("limit"), 10)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(c.Query("offset"), 0)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query offset")
		return
	}

	fields := c.Query("fields")
	// if len(fields) > 0 {
	// 	resp, err := h.strg.User().GetFields(context.Background(), &models.GetListUserRequest{
	// 		Limit:  limit,
	// 		Offset: offset,
	// 		Fields: fields,
	// 	})
	// 	if err != nil {
	// 		handleResponse(c, http.StatusBadRequest, "invalid query fields")
	// 		return
	// 	}
	// 	handleResponse(c, http.StatusOK, resp)
	// 	return
	// }
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query fields")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.User().GetList(ctx, &models.GetListUserRequest{
		Limit:  limit,
		Offset: offset,
		Fields: fields,
	})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) UpdateUser(c *gin.Context) {

	var updateUser models.UpdateUser

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateUser.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rowsAffected, err := h.strg.User().Update(ctx, &updateUser)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(c, http.StatusBadRequest, "no rows affected")
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.User().GetByID(ctx, &models.UserPrimaryKey{Id: updateUser.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.User().Delete(ctx, &models.UserPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}

// ------------------------------------------------------------------------------------------------------------------------------
func (h *Handler) MultiCreate(c *gin.Context) {
	var createUsers []models.CreateUser

	err := c.ShouldBindJSON(&createUsers)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err: "+err.Error())
		return
	}
	for i := range createUsers {

		ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
		defer cancel()
		resp, err := h.strg.User().Create(ctx, &createUsers[i])
		if err != nil {
			handleResponse(c, http.StatusInternalServerError, err)
			return
		}
		handleResponse(c, http.StatusCreated, resp)
	}

}

func (h *Handler) MultiUpdate(c *gin.Context) {
	var updateUsers []models.UpdateUser

	err := c.ShouldBindJSON(&updateUsers)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	for i := range updateUsers {

		ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
		defer cancel()
		rowsAffected, err := h.strg.User().Update(ctx, &updateUsers[i])
		if err != nil {
			handleResponse(c, http.StatusInternalServerError, err)
			return
		}

		if rowsAffected == 0 {
			handleResponse(c, http.StatusBadRequest, "no rows affected")
			return
		}

		ctx, cancel = context.WithTimeout(context.Background(), config.CtxTimeout)
		defer cancel()

		resp, err := h.strg.User().GetByID(ctx, &models.UserPrimaryKey{Id: updateUsers[i].Id})
		if err != nil {
			handleResponse(c, http.StatusInternalServerError, err)
			return
		}
		handleResponse(c, http.StatusAccepted, resp)
	}

}

func (h *Handler) MultiDelete(c *gin.Context) {
	var ids []models.UserPrimaryKey
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	for i := range ids {
		ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
		defer cancel()

		err := h.strg.User().Delete(ctx, &models.UserPrimaryKey{Id: ids[i].Id})
		if err != nil {
			handleResponse(c, http.StatusInternalServerError, err)
			return
		}

		handleResponse(c, http.StatusNoContent, nil)
	}

}

func (h *Handler) GetFields(c *gin.Context) {

}
