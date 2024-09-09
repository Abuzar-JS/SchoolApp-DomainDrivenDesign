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
	for _, c := range cs {
		getCourses = append(getCourses, c.toDomain())
	}

	return getCourses
}

type Course struct {
	ID        int    `gorm:"primary_key;column:id"`
	Title     string `gorm:"unique;not null;column:title"`
	StudentID int    `gorm:"not null;unique;column:student_id"`
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

func (u *CoursePostgres) GetByStudentID(filterType string, id int) (domain.Courses, error) {

	var course Courses

	query := u.Db

	if filterType == "student" {
		query = u.Db.Where("student_id=?", id)
	}

	result := query.Find(&course)
	if result.Error != nil {
		return nil, fmt.Errorf("courses not found: %w", result.Error)
	}

	return course.toDomain(), nil
}

func (u *CoursePostgres) GetByCourseID(courseID int) (Course domain.Course, err error) {
	var course domain.Course
	result := u.Db.First(&course, courseID)
	if result.Error != nil {
		return course, fmt.Errorf("course not found")
	}
	return course, nil
}

func (u *CoursePostgres) Save(course domain.Course) error {
	c := fromDomain(course)
	result := u.Db.Create(c)
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
