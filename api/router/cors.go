package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method != "OPTIONS" {

			c.Next()

		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
