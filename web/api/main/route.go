package main

import (
	"github.com/RTE/web/api/presentation/handler"
	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine) {
	api := r.Group("api")
	{
		api.GET("token/:address", handler.HandleGetUserToken)
	}
}
