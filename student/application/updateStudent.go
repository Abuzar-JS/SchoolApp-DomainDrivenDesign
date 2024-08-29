package application

import (
	"context"
	"data/student/domain/student"
	"fmt"
)

type UpdateStudentRequest struct {
	Name     *string
	Class    *string
	SchoolID *int
}

func (s UpdateStudentRequest) Validate(ctx context.Context) error {
	if *s.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if *s.SchoolID == 0 {
		return fmt.Errorf("id must be greater then 0")
	}

	if *s.Class == "" {
		return fmt.Errorf("class cannot be empty")
	}

	return nil

}

type UpdateStudent func(ctx context.Context, studentID int, request UpdateStudentRequest) error

func NewUpdateStudent(
	studentRepo student.StudentRepository,
) UpdateStudent {
	return func(ctx context.Context, studentID int, request UpdateStudentRequest) error {

		err := request.Validate(ctx)
		if err != nil {
			return err
		}
		studentData, err := studentRepo.GetStudentById(studentID)
		if err != nil {
			return fmt.Errorf("student can't update ")
		}

		if request.Name != nil {
			studentData.Name = *request.Name
		}

		if request.Class != nil {
			studentData.Class = *request.Class
		}

		if request.SchoolID != nil {
			studentData.SchoolID = *request.SchoolID
		}

		if err := studentRepo.Update(studentID, studentData); err != nil {
			return fmt.Errorf("update request failed: %w", err)
		}
		return nil
	}
}
