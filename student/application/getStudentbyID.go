package application

import (
	"data/student/domain"
	"data/student/domain/student"
	"fmt"
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
