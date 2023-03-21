package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/guerraglucas/ginApi/controllers"
)

func StartRoutes() {
	{
		r := gin.Default()
		r.GET("/", controller.Index)
		r.GET("/students", controller.ReturnAllStudents)
		r.GET("/students/:id", controller.ReturnSingleStudent)
		r.Run(":8080")
	}
}
