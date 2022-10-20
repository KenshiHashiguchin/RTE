package main

import (
	"github.com/RTE/web/api/presentation/handler"
	"github.com/RTE/web/api/presentation/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logSetting(r)
	r.Use(middleware.ResponseError()) // error response

	r.GET("api/token/:address", handler.HandleGetUserToken)
	r.Run(":80")
}
