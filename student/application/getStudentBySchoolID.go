package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/student/domain"
	"github.com/Abuzar-JS/Go-StudentApp/student/domain/schoolClient"
	"github.com/Abuzar-JS/Go-StudentApp/student/domain/student"
)

type GetStudentBySchoolID func(schoolID int) ([]domain.Student, error)

func NewGetStudentBySchoolID(
	studentRepo student.StudentRepository,
	schoolClient schoolClient.SchoolClient,
) GetStudentBySchoolID {

	return func(schoolID int) ([]domain.Student, error) {

		_, err := schoolClient.GetBySchoolIdClient(context.Background(), schoolID)
		if err != nil {
			return nil, fmt.Errorf("no school found with ID %d", schoolID)

		}
		students, err := studentRepo.GetBySchoolID(schoolID)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve students for school ID %d: %w", schoolID, err)
		}
		return students, nil
	}
}
