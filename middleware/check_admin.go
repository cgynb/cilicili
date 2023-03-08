package middleware

import (
	"cilicili/orm"
	"cilicili/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *orm.User
		if val, ok := c.Get("user"); ok && val != nil {
			user, _ = val.(*orm.User)
		}
		if user.Role != 1 {
			c.Abort()
			c.JSON(http.StatusForbidden, response.RespErr(http.StatusForbidden, "you're not admin", "admin required"))
		} else {
			c.Next()
		}
	}
}
