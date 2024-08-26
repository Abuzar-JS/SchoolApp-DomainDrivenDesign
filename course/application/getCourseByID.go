package application

import (
	"data/course/domain"
	"data/course/domain/course"
	"data/school/domain/school"
	"data/student/domain/student"
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
	studentRepo student.StudentRepository,
	schoolRepo school.SchoolRepository,
) GetCoursebyID {
	return func(request GetRequestByCourseID) (domain.Course, error) {

		scID, err := schoolRepo.GetBySchoolID(request.SchoolID)
		if err != nil {
			return domain.Course{}, fmt.Errorf(" no school found with ID %v", scID)

		}

		stID, err := studentRepo.GetStudentById(request.StudentID)
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
