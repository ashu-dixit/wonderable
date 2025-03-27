package routes

import (
	"wonderable/controllers"
	"wonderable/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	parent := router.Group("/parent")
	{
		parent.GET("/student", middleware.AuthMiddleware("Parent"), controllers.GetStudentsByParent)
	}

	teacher := router.Group("/teacher")
	{
		teacher.GET("/student", middleware.AuthMiddleware("Teacher"), controllers.GetStudentsByParent)
	}
}
