package svr

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"local/config"
	"net/http"
	"os"
    "log"
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

func readdir(r gin.IRouter) {
	r.GET("/readdir", func(c *gin.Context) {
		infos := []catalogInfo{}
		path, _ := c.GetQuery("dirname")
		apath, err := getAbsPath(path)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody(err))
			return
		}

		files, err := ioutil.ReadDir(apath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		for _, file := range files {
            path := file.Name();
            log.Println(path)
			if file.IsDir() {
				infos = append(infos, catalogInfo{1, path, 0})
			} else {
				infos = append(infos, catalogInfo{0, path, file.Size()})
			}

		}

		jsonBytes, err := json.Marshal(infos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorBody(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "data": string(jsonBytes)})
	})
}
