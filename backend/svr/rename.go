
package svr

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func rename(r gin.IRouter) {
	r.GET("/rename", func(c *gin.Context) {
		oldname, _ := c.GetQuery("oldname")
		newname, _ := c.GetQuery("newname")
		aopath, err := getAbsPath(oldname)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		anpath, err := getAbsPath(newname)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		if err = os.Rename(aopath, anpath); err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}
