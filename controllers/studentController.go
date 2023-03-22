package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guerraglucas/ginApi/models"
)

// StudentController is the controller for the Student model
type StudentController struct {
	StudentRepository models.StudentRepository
}

// CONSTRUCTOR NewStudentController returns a new StudentController
func NewStudentController(studentRepository models.StudentRepository) *StudentController {
	return &StudentController{
		StudentRepository: studentRepository,
	}
}

// ReturnAllStudents returns all students from the StudentRepository
func (r *StudentController) ReturnAllStudents(c *gin.Context) {
	students, err := r.StudentRepository.GetAllStudents()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting students",
		})
	}
	c.JSON(200, gin.H{
		"students": students,
	})
}

// ReturnSingleStudent returns a single student based on the id passed in the url
func (r *StudentController) ReturnSingleStudent(c *gin.Context) {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	student, err := r.StudentRepository.GetStudent(idConverted)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting student",
		})
	}
	c.JSON(200, gin.H{
		"student": student,
	})
}

// CreateStudent creates a new student from the postgres db
func (r *StudentController) CreateStudent(c *gin.Context) {
	var reqBody map[string]interface{}

	err := json.NewDecoder(c.Request.Body).Decode(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	name, ok := reqBody["name"].(string)
	if !ok {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	age, ok := reqBody["age"].(float64)
	if !ok {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	student, err := r.StudentRepository.CreateStudent(name, int(age))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating student",
		})
	}
	json.NewEncoder(c.Writer).Encode(&student)
}

func (r *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	_, err := r.StudentRepository.DeleteStudent(idConverted)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error deleting student",
		})
	}
	c.JSON(200, gin.H{
		"message": "Student deleted",
	})

}

func (r *StudentController) UpdateStudent(c *gin.Context) {
	var reqBody map[string]interface{}

	errJson := json.NewDecoder(c.Request.Body).Decode(&reqBody)
	if errJson != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	name, ok := reqBody["name"].(string)
	if !ok {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	age, ok := reqBody["age"].(float64)
	if !ok {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	_, err := r.StudentRepository.UpdateStudent(idConverted, name, int(age))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error updating student",
		})
	}
	c.JSON(200, gin.H{
		"message": "Student updated",
	})
}
