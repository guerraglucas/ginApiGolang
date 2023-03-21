package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/guerraglucas/ginApi/controllers"
)

func StartRoutes(c *controller.StudentController) {
	{
		r := gin.Default()
		r.GET("/", controller.Index)
		r.GET("/:name", controller.ReturnGreetings)
		r.GET("/students", c.ReturnAllStudents)
		r.POST("/students", c.CreateStudent)
		r.GET("/students/:id", c.ReturnSingleStudent)
		r.PUT("/students/:id", c.UpdateStudent)
		r.DELETE("/students/:id", c.DeleteStudent)
		r.Run(":8080")
	}
}
