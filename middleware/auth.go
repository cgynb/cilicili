package middleware

import (
	"cilicili/orm"
	"cilicili/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		userClaims, ok := ParseToken(token)
		if !ok {
			c.Abort()
			c.JSON(http.StatusForbidden, response.RespErr(http.StatusForbidden, "please login again", "auth error"))
		} else if userClaims.Forbid == 1 {
			c.Abort()
			c.JSON(http.StatusForbidden, response.RespErr(http.StatusForbidden, "you account has been banned", "you account has been banned"))
		} else {
			id := strconv.Itoa(int(userClaims.ID))
			user, ok := orm.GetUser("id", id)
			if ok {
				c.Set("user", user)
				token, _ = GenToken(&UserClaims{
					User:           *user,
					StandardClaims: jwt.StandardClaims{},
				})
				c.Header("Authorization", token)
				c.Next()
			} else {
				c.Abort()
				c.JSON(http.StatusForbidden, response.RespErr(http.StatusForbidden, "please login again", "auth error"))
			}
		}
	}
}
