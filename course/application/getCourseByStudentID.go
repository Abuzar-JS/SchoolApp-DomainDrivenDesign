package application

import (
	"data/course/domain"
	"data/course/domain/course"
	"data/school/domain/school"
	"data/student/domain/student"
	"fmt"
)

type GetCourseRequestByStudentID struct {
	StudentID int
	SchoolID  int
}

type GetCourseByStudentID func(GetCourseRequestByStudentID) ([]domain.Course, error)

func NewGetCourseByStudentID(
	courseRepo course.CourseRepository,
	studentRepo student.StudentRepository,
	schoolRepo school.SchoolRepository,
) GetCourseByStudentID {
	return func(request GetCourseRequestByStudentID) ([]domain.Course, error) {

		scID, err := schoolRepo.GetBySchoolID(request.SchoolID)
		if err != nil {
			return nil, fmt.Errorf(" no school found with ID %v", scID)

		}

		stID, err := studentRepo.GetStudentById(request.StudentID)
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
