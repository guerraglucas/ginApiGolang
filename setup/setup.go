package setup

import (
	"github.com/guerraglucas/ginApi/controllers"
	databases "github.com/guerraglucas/ginApi/db"
	"github.com/guerraglucas/ginApi/repositories"
	"github.com/guerraglucas/ginApi/routes"
)

func Setup() {
	// Connect to database
	databases.ConnectToPostgres()
	db := databases.DB

	// Setup repositories
	studentRepository := repositories.NewStudentRepository(db)

	// Setup controllers
	studentController := controllers.NewStudentController(studentRepository)

	// Setup routes
	routes.StartRoutes(studentController)
	// Connect to redis
	// Connect to other services
}
