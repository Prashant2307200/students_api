package storage

import "github.com/Prashant2307200/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudentsList() ([]types.Student, error)
}
