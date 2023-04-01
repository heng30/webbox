package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTls(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
            SSLHost: host,
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
