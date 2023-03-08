package api

import (
	"cilicili/orm"
	"cilicili/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostCollect(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.PostForm("video_id"))
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	col, ok, msg := orm.CreateCollect(uint(videoId), user.ID)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(col))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "post comment fail", msg))
	}
}
