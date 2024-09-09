package course

import "github.com/Abuzar-JS/Go-StudentApp/course/domain"

type CourseRepository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetByStudentID(studentID int) ([]domain.Course, error)
	GetByCourseID(courseID int) (Course domain.Course, err error)
}

type WriteRepository interface {
	Delete(courseID int) error
	Save(course domain.Course) (domain.Course, error)
	Update(course domain.Course) error
}
