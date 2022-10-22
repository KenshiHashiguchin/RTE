package main

import (
	"github.com/RTE/web/api/presentation/handler"
	merchantHandler "github.com/RTE/web/api/presentation/handler/merchant"
	"github.com/RTE/web/api/presentation/middleware"
	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine) {
	api := r.Group("api")
	{
		api.GET("token/:address", handler.HandleGetUserToken)
		api.GET("/auth", handler.HandleAuth) //TODO POST

		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/reserve", handler.HandleReserve)
		}

		merchantRoute(api)
	}
}

func merchantRoute(api *gin.RouterGroup) {
	merchant := api.Group("/merchant")
	{
		merchant.GET("token/:address", merchantHandler.HandleGetToken)
		merchant.GET("/auth", merchantHandler.HandleAuth) //TODO POST
	}
}
