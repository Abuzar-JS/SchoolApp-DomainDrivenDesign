package repository

import (
	"data/course/model"
	"fmt"

	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	Db *gorm.DB
}

func NewCourseRepositoryImpl(Db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{Db: Db}
}

func (u *CourseRepositoryImpl) Delete(courseId int) error {

	var course model.Course

	result := u.Db.Where("id = ?", courseId).Delete(&course)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no course found with id %d", courseId)
	}

	return nil
}

func (u *CourseRepositoryImpl) FindByStudentID(studentID int) ([]model.Course, error) {
	var course []model.Course
	result := u.Db.Where("student_id=?", studentID).Find(&course)

	if result.Error != nil {
		return nil, fmt.Errorf(" student not found")
	}

	return course, nil
}

func (u *CourseRepositoryImpl) FindById(courseId int) (Course model.Course, err error) {
	var course model.Course
	result := u.Db.First(&course, courseId)
	if result.Error != nil {
		return course, fmt.Errorf("course not found")
	}
	return course, nil
}

func (u *CourseRepositoryImpl) Save(course *model.Course) error {
	result := u.Db.Create(course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *CourseRepositoryImpl) Update(id int, course model.Course) error {

	result := u.Db.Model(model.Course{}).Where("id=?", course.ID).Updates(course)
	if result.Error != nil {
		return fmt.Errorf("can't update")
	}

	return nil
}
