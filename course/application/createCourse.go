package application

import (
	"context"
	"data/course/domain"
	"data/course/domain/course"
	"data/school/domain/school"
	"data/student/domain/student"
	"fmt"
)

type CreateCourseRequest struct {
	Title     string
	StudentID int
	SchoolID  int
}

func (c CreateCourseRequest) Validate(ctx context.Context) error {
	if c.Title == "" {
		return fmt.Errorf("title of the course cannot be empty")
	}
	if c.StudentID <= 0 {
		return fmt.Errorf("student id must be greater than 0")
	}
	return nil
}

type CreateCourse func(ctx context.Context, request CreateCourseRequest) (*domain.Course, error)

func NewCreateCourse(
	courseRepo course.CourseRepository,
	studentRepo student.StudentRepository,
	schoolRepo school.SchoolRepository,
) CreateCourse {
	return func(ctx context.Context, request CreateCourseRequest) (*domain.Course, error) {

		scID, err := schoolRepo.GetBySchoolID(request.SchoolID)
		if err != nil {
			return nil, fmt.Errorf(" no school found with ID %v", scID)

		}

		stID, err := studentRepo.GetStudentById(request.StudentID)
		if err != nil {
			return nil, fmt.Errorf(" no student found with ID %v", stID)
		}

		err = request.Validate(ctx)
		if err != nil {
			return nil, err
		}

		courseModel := domain.Course{
			Title:     request.Title,
			StudentID: request.StudentID,
		}

		err = courseRepo.Save(&courseModel)
		if err != nil {
			return nil, fmt.Errorf("course creation failed")
		}

		return &courseModel, nil

	}
}
