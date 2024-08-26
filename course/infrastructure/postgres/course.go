package postgres

import (
	"data/course/domain"
	"data/student/infrastructure/postgres"
	"fmt"

	"gorm.io/gorm"
)

type CoursePostgres struct {
	Db *gorm.DB
}

func NewCoursePostgres(Db *gorm.DB) *CoursePostgres {
	return &CoursePostgres{
		Db: Db,
	}
}

type Course struct {
	ID        int                `gorm:"primary_key;column:id"`
	Title     string             `gorm:"unique;not null;column:title"`
	StudentID int                `gorm:"not null;unique;column:student_id"`
	Student   []postgres.Student `gorm:"-"`
}

func (c Course) TableName() string {
	return "courses"
}

func (u *CoursePostgres) Delete(courseID int) error {

	var course domain.Course

	result := u.Db.Where("id = ?", courseID).Delete(&course)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no course found with id %d", courseID)
	}

	return nil
}

func (u *CoursePostgres) GetByStudentID(studentID int) ([]domain.Course, error) {
	var course []domain.Course
	result := u.Db.Where("student_id=?", studentID).Find(&course)

	if result.Error != nil {
		return nil, fmt.Errorf(" student not found")
	}

	return course, nil
}

func (u *CoursePostgres) GetByCourseID(courseID int) (Course domain.Course, err error) {
	var course domain.Course
	result := u.Db.First(&course, courseID)
	if result.Error != nil {
		return course, fmt.Errorf("course not found")
	}
	return course, nil
}

func (u *CoursePostgres) Save(course *domain.Course) error {
	result := u.Db.Create(course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *CoursePostgres) Update(id int, course domain.Course) error {

	result := u.Db.Model(domain.Course{}).Where("id=?", course.ID).Updates(course)
	if result.Error != nil {
		return fmt.Errorf("can't update")
	}

	return nil
}
