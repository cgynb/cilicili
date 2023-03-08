package api

import (
	"cilicili/orm"
	"cilicili/redis"
	"cilicili/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UploadVideo(c *gin.Context) {
	videoUrl := c.PostForm("video_url")
	videoType := c.PostForm("video_type")
	videoName := c.PostForm("video_name")
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	v, ok, msg := orm.CreateVideo(user.ID, videoType, videoUrl, videoName)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(gin.H{
			"video": v,
		}))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, msg, msg))
	}

}

func SendBulletChat(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.PostForm("video_id"))
	videoTime := c.PostForm("video_time")
	text := c.PostForm("text")
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	b, ok, msg := orm.CreateBulletChat(uint(videoId), user.ID, text, videoTime)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(b))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "send bullet chat fail", msg))
	}
}

func SearchVideo(c *gin.Context) {
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	var searchText string
	conditions := make(map[string]any)
	if c.Query("video_id") != "" {
		searchText = c.Query("video_id")
		conditions["id"], _ = strconv.Atoi(c.Query("video_id"))
	}
	if c.Query("video_type") != "" {
		searchText = c.Query("video_type")
		conditions["type"] = c.Query("video_type")
	}
	if c.Query("video_name") != "" {
		searchText = c.Query("video_name")
		conditions["video_name"] = c.Query("video_name")
	}
	if c.Query("user_id") != "" {
		searchText = c.Query("user_id")
		conditions["user_id"], _ = strconv.Atoi(c.Query("user_id"))
	}
	v, ok := orm.GetVideos(conditions)
	redis.AddSearchHistory(strconv.Itoa(int(user.ID)), searchText)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(v))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "search fail", "search fail"))
	}
}

func GetLeagueTable(c *gin.Context) {
	l, ok := redis.GetClickList()
	if ok {
		c.JSON(http.StatusOK, response.RespOk(l))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "get league table fail", "get league table fail"))
	}
}

func AddClick(c *gin.Context) {
	videoId, err := strconv.ParseFloat(c.PostForm("video_id"), 64)
	clickNum, ok := redis.AddClick(videoId)
	if err != nil && ok {
		c.JSON(http.StatusForbidden, response.RespErr(http.StatusForbidden, "add click fail", "add click fail"))
	} else {
		c.JSON(http.StatusOK, response.RespOk(gin.H{
			"video_id":  videoId,
			"click_num": clickNum,
		}))
	}
}

func GetSearchHistory(c *gin.Context) {
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	v, ok := redis.GetSearchHistory(strconv.Itoa(int(user.ID)))
	if ok {
		c.JSON(http.StatusOK, response.RespOk(v))
	} else {
		c.JSON(http.StatusOK, response.RespErr(http.StatusInternalServerError, "get history fail", "get history fail"))
	}
}
