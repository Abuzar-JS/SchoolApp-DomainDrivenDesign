package service

import (
	"data/course/controller/request"
	"data/course/controller/response"
	"data/course/model"
)

type CourseService interface {
	Create(course request.CreateCourseRequest) (model.Course, error)
	Update(id int, course request.UpdateCourseRequest) error
	Delete(courseId int) error
	FindById(request GetCourseRequest) (response.CourseResponse, error)
	FindByStudentID(request GetCourseRequest) ([]response.CourseResponse, error)
}
