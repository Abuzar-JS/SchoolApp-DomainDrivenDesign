package application

import (
	"context"
	"data/course/domain/course"
	"data/school/domain/school"
	"data/student/domain/student"
	"fmt"
)

type UpdateCourseRequest struct {
	Title     *string
	StudentID *int
	SchoolID  int
	CourseID  int
}

func (c UpdateCourseRequest) Validate(ctx context.Context) error {
	if *c.Title == "" {
		return fmt.Errorf("title of the course cannot be empty")
	}
	if *c.StudentID <= 0 {
		return fmt.Errorf("student id must be greater than 0")
	}
	return nil
}

type UpdateCourse func(ctx context.Context, request UpdateCourseRequest) error

func NewUpdateCourse(
	courseRepo course.CourseRepository,
	studentRepo student.StudentRepository,
	schoolRepo school.SchoolRepository,
) UpdateCourse {
	return func(ctx context.Context, request UpdateCourseRequest) error {

		scID, err := schoolRepo.GetBySchoolID(request.SchoolID)
		if err != nil {
			return fmt.Errorf(" no school found with ID %v", scID)

		}

		stID, err := studentRepo.GetStudentById(*request.StudentID)
		if err != nil {
			return fmt.Errorf(" no student found with ID %v", stID)
		}

		courseData, err := courseRepo.GetByCourseID(request.CourseID)
		if err != nil {
			return fmt.Errorf("can't update course ")
		}

		if request.Title != nil {
			courseData.Title = *request.Title
		}

		if request.StudentID != nil {
			courseData.StudentID = *request.StudentID
		}
		if err := courseRepo.Update(request.CourseID, courseData); err != nil {
			return fmt.Errorf("update request failed: %w", err)
		}
		return nil
	}

}
