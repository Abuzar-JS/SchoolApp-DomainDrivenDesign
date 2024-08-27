package application

import (
	"context"
	"data/student/domain"
	"data/student/domain/student"
	"fmt"
)

type GetStudentBySchoolID func(schoolID int) ([]domain.Student, error)

func NewGetStudentBySchoolID(
	studentRepo student.StudentRepository,
	schoolClient domain.SchoolClient,
) GetStudentBySchoolID {

	return func(schoolID int) ([]domain.Student, error) {
		fmt.Println("sc id", schoolID)

		_, err := schoolClient.GetStudentByIdClient(context.Background(), schoolID)
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
