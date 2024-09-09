package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/course/domain"
	"github.com/Abuzar-JS/Go-StudentApp/course/domain/course"
	"github.com/Abuzar-JS/Go-StudentApp/course/domain/schoolClient"
	"github.com/Abuzar-JS/Go-StudentApp/course/domain/studentClient"
)

type GetCourseRequestByStudentID struct {
	StudentID int
	SchoolID  int
}

type GetCourseByStudentID func(GetCourseRequestByStudentID) ([]domain.Course, error)

func NewGetCourseByStudentID(
	courseRepo course.CourseRepository,
	studentClient studentClient.StudentClient,
	schoolClient schoolClient.SchoolClient,
) GetCourseByStudentID {
	return func(request GetCourseRequestByStudentID) ([]domain.Course, error) {

		_, err := schoolClient.GetBySchoolIdClient(context.Background(), request.SchoolID)
		if err != nil {
			return nil, fmt.Errorf(" no school found with ID %v", request.SchoolID)

		}

		stID, err := studentClient.GetStudentByIdClient(context.Background(), request.StudentID)
		if err != nil {
			return nil, fmt.Errorf(" no student found with ID %v", stID)
		}

		courses, err := courseRepo.GetByStudentID(request.StudentID)
		if err != nil {
			return nil, fmt.Errorf("no course found against the student")
		}

		return courses, nil
	}
}
