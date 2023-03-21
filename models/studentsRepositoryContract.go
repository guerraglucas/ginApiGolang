package models

type StudentRepository interface {
	CreateStudent(name string, age int) Student
	DeleteStudent(id int) Student
	UpdateStudent(id int, name string, age int) Student
	GetStudent(id int) Student
	GetAllStudents() []Student
}
