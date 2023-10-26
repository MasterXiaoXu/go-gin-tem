

package routes

import (
   "bv/app"
   "bv/util"
   "github.com/gin-gonic/gin"
   jwt "github.com/appleboy/gin-jwt/v2"
)



func SetRouter() *gin.Engine  {
   r :=gin.Default()

   ginJWTMiddleware := util.JwtInit()
   jwtMiddleware, _ := jwt.New(ginJWTMiddleware)
   /**
   用户User路由组
    */
   userGroup :=r.Group("user/v1")
   userGroup.Use(jwtMiddleware.MiddlewareFunc())
   {
      //增加用户User

      userGroup.GET("/refresh_token", jwtMiddleware.RefreshHandler)
      userGroup.POST("/register",app.Register,)
      userGroup.POST("/login",app.Login)

      
   }
   
   return r
}