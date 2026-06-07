package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/course/domain/course"
	"github.com/Abuzar-JS/Go-StudentApp/course/domain/schoolClient"
	"github.com/Abuzar-JS/Go-StudentApp/course/domain/studentClient"
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

		student, err := studentClient.GetStudentByIdClient(context.Background(), request.StudentID)
		if err != nil {
			return fmt.Errorf(" no student found with ID %v", request.StudentID)
		}
		if student.SchoolID != request.SchoolID {
			return fmt.Errorf("student %v does not belong to school %v", request.StudentID, request.SchoolID)
		}

		course, err := courseRepo.GetByCourseID(request.CourseID)
		if err != nil {
			return fmt.Errorf("id Does not Exist")
		}
		if course.StudentID != request.StudentID {
			return fmt.Errorf("course %v does not belong to student %v", request.CourseID, request.StudentID)
		}

		err = courseRepo.Delete(request.CourseID)

		if err != nil {
			return fmt.Errorf("id Does not Exist")

		}
		return nil
	}
}
