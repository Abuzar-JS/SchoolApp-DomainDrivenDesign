package application

import (
	"context"
	"data/course/domain/course"
	"data/course/domain/schoolClient"
	"data/course/domain/studentClient"
	"fmt"
)

type DeleteCourseRequest struct {
	CourseID  int
	StudentID int
	SchoolID  int
}

type DeleteCourse func(ctx context.Context, request DeleteCourseRequest) error

func NewDeleteCourse(
	courseRepo course.CourseRepository,
	studentClient studentClient.StudentClient,
	schoolClient schoolClient.SchoolClient,
) DeleteCourse {
	return func(ctx context.Context, request DeleteCourseRequest) error {

		_, err := schoolClient.GetBySchoolIdClient(context.Background(), request.SchoolID)
		if err != nil {
			return fmt.Errorf(" no school found with ID %v", request.SchoolID)

		}

		stID, err := studentClient.GetStudentByIdClient(context.Background(), request.StudentID)
		if err != nil {
			return fmt.Errorf(" no student found with ID %v", stID)
		}

		err = courseRepo.Delete(request.CourseID)

		if err != nil {
			return fmt.Errorf("id Does not Exist")

		}
		return nil
	}
}
