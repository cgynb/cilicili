package api

import (
	"cilicili/orm"
	"cilicili/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetComment(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	page, _ := strconv.Atoi(c.Query("page"))
	comments, subComments, ok := orm.GetComments(uint(videoId), page)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(gin.H{
			"comments":     comments,
			"sub_comments": subComments,
			"page":         page,
		}))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "get comment fail", "get comment fail"))
	}
}

func PostComment(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.PostForm("video_id"))
	replyId := c.PostForm("reply_id")
	text := c.PostForm("text")
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	com, ok, msg := orm.CreateComment(uint(videoId), user.ID, replyId, text)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(com))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "post comment fail", msg))
	}
}
