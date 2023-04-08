package svr

import (
	"github.com/gin-gonic/gin"
	"local/config"
	"net/http"
)

type conf struct {
	CanDelete bool `json:"canDelete"`
    UploadChunkSize int `json:"uploadChunkSize"`
}

func configure(r gin.IRouter) {
	r.GET("/configure", func(c *gin.Context) {
		cf := conf {
			CanDelete: config.AppConf.CanDelete,
            UploadChunkSize: config.AppConf.UploadChunkSize,
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "data": cf})
	})
}
