package svr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func download(r gin.IRouter) {
	r.GET("/download", func(c *gin.Context) {
		filename, _ := c.GetQuery("filename")

		apath, err := getAbsPath(filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		file, err := os.Open(apath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
		c.Writer.WriteHeader(http.StatusOK)

		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}
	})
}
