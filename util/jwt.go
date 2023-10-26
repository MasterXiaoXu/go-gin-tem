package util

import (

	"github.com/gin-gonic/gin"
	"time"
	// "fmt"
	"strings"
	jwt "github.com/appleboy/gin-jwt/v2"
)


var identityKey = "id"

type User struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func authorizator(data interface{}, ctx *gin.Context) bool {
	return strings.Contains(data.(string), "admin")
}

func authenticator(c *gin.Context) (interface{}, error) {
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	if user.Username == "admin" &&  user.Password == "admin" {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}



func JwtInit() *jwt.GinJWTMiddleware{

	ginJWTMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: authenticator,
		Authorizator:  authorizator,
		Unauthorized: func(c *gin.Context, code int, message string) {

			c.JSON(200, gin.H{
				"code":    code,
				"message": message,
			})
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims["username"]
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			
			
			// if user, ok := data.(userCenter.User); ok {
			// 	return jwt.MapClaims{"username": user.Username}
			// }
			return jwt.MapClaims{"username":"1111"}
		},
	}


	return ginJWTMiddleware
	


}


