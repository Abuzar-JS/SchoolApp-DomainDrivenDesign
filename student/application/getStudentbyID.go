package application

import (
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/student/domain"
	"github.com/Abuzar-JS/Go-StudentApp/student/domain/student"
)

type GetByStudentID func(studentID int) (domain.Student, error)

func NewGetByStudentID(
	studentRepo student.StudentRepository,
) GetByStudentID {
	return func(studentID int) (domain.Student, error) {
		student, err := studentRepo.GetStudentById(studentID)
		if err != nil {
			return domain.Student{}, fmt.Errorf("could not retrieve student: %w", err)
		}
		return student, nil
	}
}
