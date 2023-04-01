package svr

import (
	"errors"
	"github.com/gin-gonic/gin"
	"local/config"
	"net/http"
	"os"
)

func delfile(r gin.IRouter) {
	r.GET("/delfile", func(c *gin.Context) {
		if !config.AppConf.CanDelete {
			c.JSON(http.StatusForbidden, errorBody(errors.New("configuration is set to not allow delete file")))
			return
		}

		filename, _ := c.GetQuery("filename")
		apath, err := getAbsPath(filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		if err = os.Remove(apath); err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}
