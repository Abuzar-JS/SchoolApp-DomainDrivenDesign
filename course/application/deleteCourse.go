package application

import (
	"context"
	"data/course/domain/course"
	"data/school/domain/school"
	"data/student/domain/student"
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
	studentRepo student.StudentRepository,
	schoolRepo school.SchoolRepository,
) DeleteCourse {
	return func(ctx context.Context, request DeleteCourseRequest) error {

		scID, err := schoolRepo.GetBySchoolID(request.SchoolID)
		if err != nil {
			return fmt.Errorf(" no school found with ID %v", scID)

		}

		stID, err := studentRepo.GetStudentById(request.StudentID)
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
