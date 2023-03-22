package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guerraglucas/ginApi/models"
	"github.com/guerraglucas/ginApi/utils"
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
		utils.HttpErrorHandler(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

// ReturnSingleStudent returns a single student based on the id passed in the url
func (r *StudentController) ReturnSingleStudent(c *gin.Context) {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	student, err := r.StudentRepository.GetStudent(idConverted)

	if err != nil {
		utils.HttpErrorHandler(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

// CreateStudent creates a new student from the postgres db
func (r *StudentController) CreateStudent(c *gin.Context) {
	var newStudent models.Student
	err := c.ShouldBind(&newStudent)
	if err != nil {
		utils.HttpErrorHandler(err, c)
		return
	}

	student, err := r.StudentRepository.CreateStudent(newStudent)
	if err != nil {
		utils.HttpErrorHandler(err, c)
		return
	}
	json.NewEncoder(c.Writer).Encode(&student)
}

func (r *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	_, err := r.StudentRepository.DeleteStudent(idConverted)
	if err != nil {
		utils.HttpErrorHandler(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted",
	})

}

func (r *StudentController) UpdateStudent(c *gin.Context) {
	var studentToUpdate models.Student
	err := c.ShouldBind(&studentToUpdate)
	if err != nil {
		utils.HttpErrorHandler(err, c)
		return
	}

	_, errRepo := r.StudentRepository.UpdateStudent(studentToUpdate)
	if errRepo != nil {
		utils.HttpErrorHandler(errRepo, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated",
	})
}
