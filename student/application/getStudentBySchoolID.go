package application

import (
	"data/school/domain/school"
	"data/student/domain"
	"data/student/domain/student"
	"fmt"
)

type GetStudentBySchoolID func(schoolID int) ([]domain.Student, error)

func NewGetStudentBySchoolID(
	studentRepo student.StudentRepository,
	schoolRepo school.SchoolRepository,
) GetStudentBySchoolID {

	return func(schoolID int) ([]domain.Student, error) {
		ID, err := schoolRepo.GetBySchoolID(schoolID)
		if err != nil {
			return nil, fmt.Errorf("no school found with ID %v", ID)
		}
		students, err := studentRepo.GetBySchoolID(schoolID)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve students for school ID %d: %w", schoolID, err)
		}
		return students, nil
	}
}
