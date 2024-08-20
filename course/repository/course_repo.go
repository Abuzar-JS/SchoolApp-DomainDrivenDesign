package repository

import "data/course/model"

type CourseRepository interface {
	Save(course *model.Course) error
	Update(id int, course model.Course) error
	Delete(courseId int) error
	FindById(courseId int) (Course model.Course, err error)
	FindByStudentID(studentID int) ([]model.Course, error)
}
