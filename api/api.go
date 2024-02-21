package api

import (
	"task/api/handler"
	"task/config"
	"task/storage"

	"github.com/gin-gonic/gin"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	v1 := r.Group("/v1")

	//User ...
	v1.POST("/users", handler.CreateUser)
	v1.POST("/users/multi", handler.MultiCreate)
	v1.GET("/users/:id", handler.GetByIDUser)
	v1.GET("/users", handler.GetListUser)
	v1.PUT("/users/:id", handler.UpdateUser)
	v1.PUT("/users/multi", handler.MultiUpdate)
	v1.DELETE("/users/:id", handler.DeleteUser)
	v1.DELETE("/users/multi", handler.MultiDelete)

}
