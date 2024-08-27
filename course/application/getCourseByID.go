package application

import (
	"context"
	"data/course/domain"
	"data/course/domain/course"
	"data/course/domain/schoolClient"
	"data/course/domain/studentClient"
	"fmt"
)

type GetRequestByCourseID struct {
	CourseID  int
	StudentID int
	SchoolID  int
}

type GetCoursebyID func(request GetRequestByCourseID) (domain.Course, error)

func NewGetCourseByID(
	courseRepo course.CourseRepository,
	studentClient studentClient.StudentClient,
	schoolClient schoolClient.SchoolClient,
) GetCoursebyID {
	return func(request GetRequestByCourseID) (domain.Course, error) {

		_, err := schoolClient.GetBySchoolIdClient(context.Background(), request.SchoolID)
		if err != nil {
			return domain.Course{}, fmt.Errorf(" no school found with ID %v", request.SchoolID)

		}

		stID, err := studentClient.GetStudentByIdClient(context.Background(), request.StudentID)
		if err != nil {
			return domain.Course{}, fmt.Errorf(" no student found with ID %v", stID)
		}

		course, err := courseRepo.GetByCourseID(request.CourseID)

		if err != nil {
			return domain.Course{}, fmt.Errorf("could not retrieve course: %w", err)
		}

		return course, nil
	}
}
