package svr

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "local/token"
)

func login(r gin.IRouter) {
	r.GET("/login", func(c *gin.Context) {
		token := token.GetToken()
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": token})
	})
}
