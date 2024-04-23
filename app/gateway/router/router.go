package router

import (
	"todo_list/app/gateway/http"
	"todo_list/app/gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
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
			authed.POST("task", http.CreateTaskHandler)
			authed.POST("update_task", http.UpdateTaskHandler)
			authed.POST("delete_task", http.DeleteTaskHandler)
			authed.GET("list_task", http.ListTaskHandler)
			authed.GET("get_task", http.GetTaskHandler)
		}

	}
	return ginRouter
}
