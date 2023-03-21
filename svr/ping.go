package svr

import "github.com/gin-gonic/gin"

func ping(r gin.IRouter) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
