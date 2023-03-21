package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guerraglucas/ginApi/models"
)

// ReturnAllStudents returns all students
func ReturnAllStudents(c *gin.Context) {

	c.JSON(200, gin.H{
		"students": []models.Student{
			{ID: 1, Name: "John", Age: 20},
			{ID: 2, Name: "Jane", Age: 21},
			{ID: 3, Name: "Jack", Age: 22},
		},
	})

}

// ReturnSingleStudent returns a single student based on the id passed in the url
func ReturnSingleStudent(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	c.JSON(200, gin.H{
		"student": models.Student{
			ID:   idInt,
			Name: "John",
			Age:  20,
		},
	})

}
