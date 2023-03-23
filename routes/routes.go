package routes

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	controller "github.com/guerraglucas/ginApi/controllers"
)

func StartRoutes(c *controller.StudentController) {
	{
		gin.SetMode(gin.DebugMode)
		r := gin.New()
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		r.GET("/", controller.Index)
		r.GET("/:name", controller.ReturnGreetings)
		r.GET("/students", c.ReturnAllStudents)
		r.POST("/students", c.CreateStudent)
		r.GET("/students/:id", c.ReturnSingleStudent)
		r.PATCH("/students/:id", c.UpdateStudent)
		r.DELETE("/students/:id", c.DeleteStudent)
		r.Run(":8080")
	}
}
