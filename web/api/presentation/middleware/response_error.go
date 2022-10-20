package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ResponseError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
		}
	}
}
