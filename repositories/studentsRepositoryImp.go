package repositories

import (
	m "github.com/guerraglucas/ginApi/models"
	"gorm.io/gorm"
)

type StudentRepositoryImp struct {
	db *gorm.DB
}

// Constructor for StudentRepositoryImp
func NewStudentRepository(db *gorm.DB) *StudentRepositoryImp {
	return &StudentRepositoryImp{db}
}

func (r *StudentRepositoryImp) CreateStudent(name string, age int) m.Student {

	student := m.Student{Name: name, Age: age}
	r.db.Create(&student)
	return student
}

func (r *StudentRepositoryImp) DeleteStudent(id int) m.Student {
	student := m.Student{}
	r.db.Delete(&student, id)
	return student
}

func (r *StudentRepositoryImp) UpdateStudent(id int, name string, age int) m.Student {
	student := m.Student{}
	r.db.Model(&student).Where("id = ?", id).Updates(m.Student{Name: name, Age: age})
	return student
}

func (r *StudentRepositoryImp) GetStudent(id int) m.Student {
	student := m.Student{}
	r.db.First(&student, id)
	return student
}

func (r *StudentRepositoryImp) GetAllStudents() []m.Student {
	students := []m.Student{}
	r.db.Find(&students)
	return students
}

// this verifies that StudentRepositoryImp implements StudentRepository
var _ m.StudentRepository = &StudentRepositoryImp{}
