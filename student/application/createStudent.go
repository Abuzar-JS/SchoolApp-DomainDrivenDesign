package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/student/domain"
	"github.com/Abuzar-JS/Go-StudentApp/student/domain/student"
)

type CreateStudentRequest struct {
	Name     string
	Class    string
	SchoolID int
}

func (s CreateStudentRequest) Validate(ctx context.Context) error {
	if s.Name == "" || s.Class == "" {
		return fmt.Errorf("name or class cannot be empty")
	}
	if s.SchoolID <= 0 {
		return fmt.Errorf("school id must be greater than 0")
	}

	return nil
}

type CreateStudent func(ctx context.Context, request CreateStudentRequest) (*domain.Student, error)

func NewCreateStudent(
	studentRepo student.StudentRepository,
) CreateStudent {
	return func(ctx context.Context, request CreateStudentRequest) (*domain.Student, error) {
		err := request.Validate(ctx)
		if err != nil {
			return nil, err
		}

		studentModel := domain.Student{
			Name:     request.Name,
			Class:    request.Class,
			SchoolID: request.SchoolID,
		}

		err = studentRepo.Save(&studentModel)
		if err != nil {
			return nil, fmt.Errorf("student creation failed")
		}

		return &studentModel, nil
	}
}
