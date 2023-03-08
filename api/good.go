package api

import (
	"cilicili/orm"
	"cilicili/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostGood(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.PostForm("video_id"))
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	gd, ok, msg := orm.CreateGood(uint(videoId), user.ID)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(gd))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "post comment fail", msg))
	}
}
