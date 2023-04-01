package svr

import (
	"errors"
	"github.com/gin-gonic/gin"
	"local/config"
	"net/http"
	"os"
)

func mkdir(r gin.IRouter) {
	r.GET("/mkdir", func(c *gin.Context) {
		dirname, _ := c.GetQuery("dirname")

		apath, err := getAbsPath(dirname)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		if err = os.MkdirAll(apath, 0750); err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}

func deldir(r gin.IRouter) {
	r.GET("/deldir", func(c *gin.Context) {
		if !config.AppConf.CanDelete {
			c.JSON(http.StatusForbidden, errorBody(errors.New("configuration is set to not allow delete directory")))
			return
		}

		dirname, _ := c.GetQuery("dirname")
		apath, err := getAbsPath(dirname)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		if err = os.RemoveAll(apath); err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}
