package main

import (
	"cilicili/api"
	"cilicili/config"
	"cilicili/orm"
	"cilicili/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.Init()
	redis.Init()
	orm.Init()
	err := api.Router().Run(":80")
	if err != nil {
		return
	}
}
