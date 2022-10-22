package main

import (
	"github.com/RTE/web/api/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func session(r *gin.Engine) {
	store := cookie.NewStore([]byte(util.Env("APP_SECRET_KEY")))
	r.Use(sessions.Sessions("user_session", store))
}
