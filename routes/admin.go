package routes

import (
	"wonderable/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.POST("assign/student/teacher", controllers.AssignStudentToTeacher) // ✅ Assign student to teacher
		admin.POST("assign/student/parent", controllers.AssignStudentToParent)   // ✅ Assign student to parent
		admin.POST("add/student", controllers.AddStudent)   // ✅ Assign student to parent
		admin.POST("/", controllers.AddSubject) // ✅ Add subject
			
	}
}	




