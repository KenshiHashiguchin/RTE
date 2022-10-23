package middleware

import (
	"github.com/RTE/web/api/domain/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		authSession := session.Get("auth")
		log.Print(authSession)
		if authSession != nil {
			_, err := model.GetAuthUserByTokenString(authSession.(string))
			if err == nil {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
}