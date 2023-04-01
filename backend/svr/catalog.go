package svr

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type catalogInfo struct {
	Type int    `json:"type"` // 0: file; 1 directory
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func catalog(r gin.IRouter) {
	r.GET("/catalog", func(c *gin.Context) {
		infos := []catalogInfo{}
		path, _ := c.GetQuery("path")
		apath, err := getAbsPath(path)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		err = filepath.Walk(apath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				if path == "." || path == ".." || path == apath {
					return nil
				}

				infos = append(infos, catalogInfo{1, path[len(apath):], 0})
			} else {
				infos = append(infos, catalogInfo{0, path[len(apath):], info.Size()})
			}

			return nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
		} else {
			jsonBytes, err := json.Marshal(infos)
			if err != nil {
				c.JSON(http.StatusInternalServerError, errorBody(err))
				return
			}

			c.JSON(http.StatusOK, gin.H{"code": 0, "data": string(jsonBytes)})
		}
	})
}
