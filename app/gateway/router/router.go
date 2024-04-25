package router

import (
	"todo_list/app/gateway/http"
	"todo_list/app/gateway/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})

		v1.POST("/user/register", http.UserRegisterHandler)
		v1.POST("/user/login", http.UserLoginHandler)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", http.ListTaskHandler)
			authed.POST("task", http.CreateTaskHandler)
			authed.GET("task/:id", http.GetTaskHandler)
			authed.PUT("task/:id", http.UpdateTaskHandler)
			authed.DELETE("task/:id", http.DeleteTaskHandler)
		}
	}
	return ginRouter
}
