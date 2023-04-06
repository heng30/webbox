package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Method() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == "GET" || method == "POST" {
			c.Next()
		} else {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"code": -1, "data": fmt.Sprintf("method: %s not allowed", method)})
			c.AbortWithStatus(http.StatusMethodNotAllowed)
		}

	}
}
