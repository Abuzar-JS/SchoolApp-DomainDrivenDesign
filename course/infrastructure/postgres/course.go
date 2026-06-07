package postgres

import (
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/course/domain"

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

type Courses []Course

func (cs Courses) toDomain() domain.Courses {
	getCourses := make(domain.Courses, len(cs))
	for i, c := range cs {
		getCourses[i] = c.toDomain()
	}

	return getCourses
}

type Course struct {
	ID        int    `gorm:"primaryKey;column:id"`
	Title     string `gorm:"unique;not null;column:title"`
	StudentID int    `gorm:"not null;column:student_id"`
}

func (c Course) toDomain() domain.Course {
	return domain.Course{
		ID:        c.ID,
		Title:     c.Title,
		StudentID: c.StudentID,
	}
}

func fromDomain(c domain.Course) Course {
	return Course{
		ID:        c.ID,
		Title:     c.Title,
		StudentID: c.StudentID,
	}
}

func (c Course) TableName() string {
	return "courses"
}

func (u *CoursePostgres) Delete(courseID int) error {

	var course Course

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

	var courses Courses

	result := u.Db.Where("student_id = ?", studentID).Find(&courses)
	if result.Error != nil {
		return nil, fmt.Errorf("courses not found: %w", result.Error)
	}

	return courses.toDomain(), nil
}

func (u *CoursePostgres) GetByCourseID(courseID int) (domain.Course, error) {
	var course Course
	result := u.Db.First(&course, courseID)
	if result.Error != nil {
		return domain.Course{}, fmt.Errorf("course not found")
	}
	return course.toDomain(), nil
}

func (u *CoursePostgres) Save(course domain.Course) (domain.Course, error) {
	c := fromDomain(course)
	result := u.Db.Create(&c)
	if result.Error != nil {
		return domain.Course{}, result.Error
	}
	return c.toDomain(), nil
}

func (u *CoursePostgres) Update(course domain.Course) error {
	c := fromDomain(course)

	result := u.Db.Model(&Course{}).Where("id = ?", course.ID).Updates(c)
	if result.Error != nil {
		return fmt.Errorf("can't update: %w", result.Error)
	}

	return nil
}
