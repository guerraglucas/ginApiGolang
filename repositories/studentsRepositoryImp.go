package repositories

import (
	"database/sql"

	m "github.com/guerraglucas/ginApi/models"
)

type StudentRepositoryImp struct {
	db *sql.DB
}

// Constructor for StudentRepositoryImp
func NewStudentRepository(db *sql.DB) *StudentRepositoryImp {
	return &StudentRepositoryImp{db}
}

func (r *StudentRepositoryImp) CreateStudent(student m.Student) (m.Student, error) {
	err := r.db.QueryRow("INSERT INTO students (name, age) VALUES ($1, $2) RETURNING id", student.Name, student.Age).Scan(&student.Id)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentRepositoryImp) DeleteStudent(id int) (m.Student, error) {
	student := m.Student{}
	err := r.db.QueryRow("DELETE FROM students WHERE id = $1 RETURNING id, name, age", id).Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentRepositoryImp) UpdateStudent(student m.Student) (m.Student, error) {

	err := r.db.QueryRow("UPDATE students SET name = $1, age = $2 WHERE id = $3 RETURNING id, name, age", student.Name, student.Age, student.Id).Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentRepositoryImp) GetStudent(id int) (m.Student, error) {
	student := m.Student{}
	err := r.db.QueryRow("SELECT * FROM students WHERE id = $1", id).Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentRepositoryImp) GetAllStudents() ([]m.Student, error) {
	var students []m.Student
	rows, err := r.db.Query("SELECT * FROM students")
	if err != nil {
		return students, err
	}
	for rows.Next() {
		student := m.Student{}
		err = rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return students, err
		}
		students = append(students, student)
	}
	return students, nil
}

// this verifies that StudentRepositoryImp implements StudentRepository
var _ m.StudentRepository = &StudentRepositoryImp{}
