package api

import (
	"cilicili/orm"
	"cilicili/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BanUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("user_id"))
	ok := orm.BanUser(uint(userId))
	if ok {
		c.JSON(http.StatusOK, response.RespOk(nil))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "db err", "db err"))
	}
}

func DelComment(c *gin.Context) {
	commentId, _ := strconv.Atoi(c.PostForm("comment_id"))
	ok := orm.DelComment(uint(commentId))
	if ok {
		c.JSON(http.StatusOK, response.RespOk(nil))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "db err", "db err"))
	}
}

func AuditVideo(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.PostForm("video_id"))
	ok := orm.DelVideo(uint(videoId))
	if ok {
		c.JSON(http.StatusOK, response.RespOk(nil))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "db err", "db err"))
	}
}
