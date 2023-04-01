package middlewares

import (
	"github.com/gin-gonic/gin"
	"local/config"
	"local/db"
	tokenMap "local/token"
	"net/http"
	"strings"
)

func Auth(testMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == "GET" || method == "POST" {
			rurl := c.Request.RequestURI
			if testMode || rurl == "/ping" {
			} else if strings.Contains(rurl, "/login") {
				username, _ := c.GetQuery("username")
				password, _ := c.GetQuery("password")

				if username == "" || password == "" {
					c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				}

				if !validUser(username, password) {
					c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				}
			} else {
				token := c.Request.Header.Get("Auth-Token")
				if token == "" {
					c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				}

				if !validToken(token) {
					c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				}
			}
			c.Next()
		}
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func validToken(token string) bool {
	if tokenMap.IsTokenTimeout(token, config.AppConf.TokenDuration) {
		tokenMap.DelToken(token)
		return false
	} else {
		tokenMap.UpdateTokenTimestamp(token)
		return true
	}
}

func validUser(username, password string) bool {
	if users, err := db.QueryUsers(); err == nil {
		for _, item := range users {
			if item.Name == username && item.Passwd == password {
				return true
			}
		}
	}
	return false
}
