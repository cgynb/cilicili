package api

import (
	"cilicili/middleware"
	"cilicili/orm"
	"cilicili/response"
	"cilicili/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	name := c.PostForm("username")
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	user, ok, msg := orm.CreateUser(name, email, pwd)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(user))
	} else {
		c.JSON(http.StatusForbidden, response.RespErr(403, msg, "register error"))
	}
}

func Login(c *gin.Context) {
	name := c.PostForm("username")
	pwd := c.PostForm("password")
	user, ok := orm.GetUser("name", name)
	if ok && utils.CheckPassword(pwd, user.Password) {
		token, _ := middleware.GenToken(&middleware.UserClaims{
			User:           *user,
			StandardClaims: jwt.StandardClaims{},
		})
		c.Header("Authorization", token)
		c.JSON(http.StatusOK, response.RespOk(user))
	} else {
		c.JSON(http.StatusForbidden, response.RespErr(403, "please login again", "name doesn't exist or password error"))
	}
}

func UpdateInfo(c *gin.Context) {
	newUserName, haveUsername := c.GetPostForm("new_username")
	newAvatar, haveAvatar := c.GetPostForm("new_avatar")
	newPassword, havePassword := c.GetPostForm("new_password")
	var ok bool
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	if haveAvatar {
		ok = orm.UpdateUser(user.ID, "avatar", newAvatar)
	} else if haveUsername {
		ok = orm.UpdateUser(user.ID, "name", newUserName)
	} else if havePassword {
		ok = orm.UpdateUser(user.ID, "password", newPassword)
	} else {
		c.JSON(http.StatusForbidden, response.RespErr(403, "please input something you want to update", "don't have explicit attribute to update"))
	}
	if ok {
		c.JSON(http.StatusOK, response.RespOk(user))
	} else {
		c.JSON(http.StatusForbidden, response.RespErr(400, "update fail", "update fail"))
	}
}

func SearchUser(c *gin.Context) {
	username, haveUsername := c.GetQuery("username")
	userId, haveUserId := c.GetQuery("user_id")
	userVideo, haveUserVideo := c.GetQuery("user_video")
	page, _ := strconv.Atoi(c.Query("page"))
	var users []*orm.User
	var user *orm.User
	var ok bool
	if haveUserVideo {
		users, ok = orm.GetUsers("id in (select user_id from videos where video_name like \"%"+userVideo+"%\")", page)
	} else if haveUsername {
		users, ok = orm.GetUsers("name LIKE "+"\"%"+username+"%\"", page)
	} else if haveUserId {
		user, ok = orm.GetUser("id", userId)
		users = append(users, user)
	}
	if ok {
		c.JSON(http.StatusOK, response.RespOk(gin.H{
			"users": users,
		}))
	} else {
		c.JSON(http.StatusInternalServerError, response.RespErr(http.StatusInternalServerError, "search fail", "db error"))
	}
}
