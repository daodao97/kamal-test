package main

import (
	"fmt"
	"kamal/pkg/conf"
	"net/http"

	"github.com/daodao97/xgo/xapp"
	"github.com/daodao97/xgo/xredis"
	"github.com/gin-gonic/gin"
)

var version string

func main() {
	app := xapp.NewApp().
		AddStartup(
			conf.Init,
			func() error {
				return xredis.Init(conf.GetRedis())
			},
		).
		AddServer(xapp.NewHttp(":8001", h))

	if err := app.Run(); err != nil {
		fmt.Printf("Application error: %v\n", err)
	}
}

func h() http.Handler {
	e := xapp.NewGin()
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": version})
	})
	return e.Handler()
}
