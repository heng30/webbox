package svr

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"local/config"
	"local/middlewares"
	"local/token"
	"log"
	"path/filepath"
	"time"
)

type routerCb func(gin.IRouter)

func Start() {
	r := gin.Default()

	mcbs := []gin.HandlerFunc{
		middlewares.Cors(),
		middlewares.Method(),
		middlewares.Auth(config.AppConf.TestMode),
	}
	if config.AppConf.EnableTLS {
		mcbs = append(mcbs, middlewares.LoadTls(config.AppConf.ListenAddr))
	}
	for _, cb := range mcbs {
		r.Use(cb)
	}

	rcbs := []routerCb{ping, login, catalog, upload, mupload, download, mkdir, rename, readdir, deldir, delfile, configure}
	for _, cb := range rcbs {
		cb(r)
	}

	go func() {
		for {
			token.ClearTimeoutToken(config.AppConf.TokenDuration)
			time.Sleep(5 * time.Minute)
		}
	}()

	if config.AppConf.EnableTLS {
		if err := r.RunTLS(config.AppConf.ListenAddr, config.CertFile, config.KeyFile); err != nil {
			log.Fatalln("Error:", err)

		}
	} else {
		if err := r.Run(config.AppConf.ListenAddr); err != nil {
			log.Fatalln("Error:", err)
		}
	}
}

func errorBody(err error) gin.H {
	return gin.H{"code": -1, "data": fmt.Sprintf("%v", err)}
}

func getAbsPath(filename string) (string, error) {
	if !filepath.HasPrefix(filename, "/") {
		return "", errors.New("Invalid filename: " + filename)
	}

	path := config.AppConf.RootPath + filename
	log.Println(path)
	apath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if !filepath.HasPrefix(apath, config.AppConf.RootPath) {
		return "", errors.New("Invalid filename: " + filename)
	}

	return apath, nil
}
