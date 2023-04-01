package svr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func upload(r gin.IRouter) {
	r.POST("/upload", func(c *gin.Context) {
		filename, _ := c.GetQuery("filename")
		apath, err := getAbsPath(filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		err = c.SaveUploadedFile(file, apath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})
}
