package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"local/config"
	"local/db"
	tokenMap "local/token"
	"net/http"
	"strings"
)

func Auth(testMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rurl := c.Request.RequestURI
		if testMode || rurl == "/ping" {
		} else if strings.Contains(rurl, "/login") {
			username, _ := c.GetQuery("username")
			password, _ := c.GetQuery("password")

			if username == "" || password == "" {
				c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"code": -1, "data": "username or password is empety"})
				c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				return
			}

			if !validUser(username, password) {
				c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"code": -1, "data": fmt.Sprintf("username: %s or password: %s is wrong", username, password)})
				c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				return
			}
		} else {
			token, _ := c.GetQuery("authtoken")
			if token == "" {
				c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"code": -1, "data": "authtoken is empety"})
				c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				return
			}

			if !validToken(token) {
				c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"code": -1, "data": fmt.Sprintf("authtoken: %s does not exist", token)})
				c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
				return
			}
		}
		c.Next()
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
