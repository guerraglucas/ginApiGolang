package models

type StudentRepository interface {
	CreateStudent(Student) (Student, error)
	DeleteStudent(id int) (Student, error)
	UpdateStudent(Student) (Student, error)
	GetStudent(id int) (Student, error)
	GetAllStudents() ([]Student, error)
	SearchStudent(string) (Student, error)
}
