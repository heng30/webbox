package svr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"local/config"
	"local/middlewares"
    "local/token"
	"log"
    "time"
)

type routerCb func(gin.IRouter)

func Start() {
	r := gin.Default()

	mcbs := []gin.HandlerFunc{middlewares.Auth(config.AppConf.TestMode), middlewares.Cors()}
	if config.AppConf.EnableTLS {
		mcbs = append(mcbs, middlewares.LoadTls(config.AppConf.ListenAddr))
	}
	for _, cb := range mcbs {
		r.Use(cb)
	}

	rcbs := []routerCb{ping, login, catalog}
	for _, cb := range rcbs {
		cb(r)
	}

    go func() {
        for {
            token.ClearTimeoutToken(config.AppConf.TokenDuration);
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
	return gin.H{"code": -1, "error": fmt.Sprintf("%v", err)}
}
