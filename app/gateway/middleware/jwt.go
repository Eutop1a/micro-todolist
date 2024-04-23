package middleware

import (
	"net/http"
	"time"
	"todo_list/pkg/ctl"
	"todo_list/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "权限过期，请重新登录",
			})
			c.Abort()
			return
		}
		ctl.InitUserinfo(c.Request.Context(), &ctl.UserInfo{Id: claims.Id})
		c.Next()
	}
}
