package api

import (
	"cilicili/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.GET("/", SearchUser)
	user.POST("/register", Register)
	user.POST("/login", Login)
	user.PUT("/info", middleware.Auth(), UpdateInfo)

	admin := v1.Group("/admin")
	admin.Use(middleware.Auth(), middleware.CheckAdmin())
	admin.PUT("/ban", BanUser)
	admin.DELETE("/comment", DelComment)
	admin.DELETE("/video", AuditVideo)

	video := v1.Group("/video")
	video.Use(middleware.Auth())
	video.GET("/", SearchVideo)
	video.POST("/", UploadVideo)
	video.GET("/comment", GetComment)
	video.POST("/comment", PostComment)
	video.POST("/good", PostGood)
	video.POST("/collect", PostCollect)
	video.POST("/bullet_chat", SendBulletChat)
	video.GET("/click", GetLeagueTable)
	video.POST("/click", AddClick)
	video.GET("/history", GetSearchHistory)

	return r
}
