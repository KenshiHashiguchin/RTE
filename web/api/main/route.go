package main

import (
	"github.com/RTE/web/api/presentation/handler"
	"github.com/RTE/web/api/presentation/middleware"
	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine) {
	api := r.Group("api")
	{
		api.GET("token/:address", handler.HandleGetUserToken)
		api.GET("/auth", handler.HandleAuth)

		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			//auth.POST("/reserve", handler.HandleReserve)
			auth.GET("/reserve", handler.HandleReserve)
		}
	}
}
