package main

import (
	"github.com/RTE/web/api/presentation/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logSetting(r)
	r.Use(middleware.ResponseError()) // error response

	route(r)

	r.Run(":80")
}
