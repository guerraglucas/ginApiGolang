package models

type StudentRepository interface {
	CreateStudent(name string, age int) (Student, error)
	DeleteStudent(id int) (Student, error)
	UpdateStudent(id int, name string, age int) (Student, error)
	GetStudent(id int) (Student, error)
	GetAllStudents() ([]Student, error)
}
