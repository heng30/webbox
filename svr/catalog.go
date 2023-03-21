package svr

import (
	// "errors"
	"github.com/gin-gonic/gin"
	// "io"
	// "local/db"
	"net/http"
)

func catalog(r gin.IRouter) {
	r.GET("/catalog", func(c *gin.Context) {
		// TODO: update token timestamp
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "test"})
	})

	// r.POST("/:token/:name", postFactory("token", "name"))
}

/*

func getFactory(token, cname string) func(*gin.Context) {
	return
    func(c *gin.Context) {
		table := c.Param(token)
        name := c.Param(cname)
		if tokens, err := db.Query(table, name); err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "data": tokens})
		} else {
			c.JSON(http.StatusNotFound, errorBody(err))
		}
	}
}

func postFactory(token, cname string) func(*gin.Context) {
	return func(c *gin.Context) {
		table := c.Param(token)
        name := c.Param(cname)
		tokens := make([]byte, 0, 4096)
		data := make([]byte, 1024)
		for {
			n, err := c.Request.Body.Read(data)
			if n > 0 {
				tokens = append(tokens, data[:n]...)
			}

			if err != nil {
				if err != io.EOF {
					c.JSON(http.StatusBadRequest, errorBody(err))
					return
				} else {
					break
				}
			}
			if n <= 0 {
				break
			}
		}

		if len(tokens) <= 0 {
			c.JSON(http.StatusBadRequest, errorBody(errors.New("Not support empty body")))
			return
		}

		if err := db.Update(table, name, string(tokens)); err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0})
		} else {
			c.JSON(http.StatusNotFound, errorBody(err))
		}
	}
}
*/
