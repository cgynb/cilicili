package middleware

import (
	"cilicili/config"
	"cilicili/orm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	orm.User
	jwt.StandardClaims
}

func GenToken(claims *UserClaims) (string, bool) {
	claims.ExpiresAt = time.Now().Add(time.Duration(config.Conf.TokenConfig.EffectTime)).Unix()
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Conf.TokenConfig.SecretKey))
	if err != nil {
		return "", false
	}
	return sign, true
}

func ParseToken(tokenString string) (*UserClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.TokenConfig.SecretKey), nil
	})
	if err != nil {
		return nil, false
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, false
	}
	return claims, true
}
