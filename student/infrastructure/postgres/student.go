package postgres

import (
	"data/student/domain"
	"fmt"

	"gorm.io/gorm"
)

type StudentPostgres struct {
	Db *gorm.DB
}

func NewStudentPostgres(Db *gorm.DB) *StudentPostgres {
	return &StudentPostgres{
		Db: Db,
	}
}

type Student struct {
	ID       int    `gorm:"primaryKey;unique;not null;column:id" json:"id"`
	Name     string `gorm:"not null;column:name" json:"name"`
	Class    string `gorm:"not null;column:class" json:"class"`
	SchoolID int    `gorm:"not null;column:school_id" json:"school_id"`
}

func (s Student) TableName() string {
	return "students"
}

func (u *StudentPostgres) Delete(studentID int) error {

	var student domain.Student

	result := u.Db.Where("id = ?", studentID).Delete(&student)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no student found with id %d", studentID)
	}

	return nil
}

func (u *StudentPostgres) GetBySchoolID(schoolID int) ([]domain.Student, error) {
	var Student []domain.Student

	result := u.Db.Where("school_id=?", schoolID).Find(&Student)
	if result.Error != nil {
		return Student, fmt.Errorf("student not found")
	}

	return Student, nil
}

func (u *StudentPostgres) GetStudentById(studentId int) (Student domain.Student, err error) {
	var student domain.Student
	result := u.Db.First(&student, studentId)

	if result.Error != nil {
		return student, fmt.Errorf("student not found")
	}

	return student, nil
}

func (u *StudentPostgres) Save(student *domain.Student) error {
	result := u.Db.Create(student)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *StudentPostgres) Update(id int, student domain.Student) error {

	result := u.Db.Model(domain.Student{}).Where("id=?", student.ID).Updates(student)

	if result.Error != nil {
		return fmt.Errorf("can't update")
	}

	return nil
}
