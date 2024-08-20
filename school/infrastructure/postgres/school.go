package postgres

import (
	"data/school/domain"
	"fmt"

	"gorm.io/gorm"
)

type SchoolPostgres struct {
	Db *gorm.DB
}

func NewSchoolPostgres(Db *gorm.DB) *SchoolPostgres {
	return &SchoolPostgres{
		Db: Db,
	}
}

type School struct {
	ID   int    `gorm:"primaryKey;unique;not null" json:"id"`
	Name string `gorm:"type:varchar(255);unique;not null" json:"name"`
}

func (s School) TableName() string {
	return "schools"
}

func (s School) ToDomain() *domain.School {
	return &domain.School{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (u *SchoolPostgres) Delete(schoolId int) error {

	var school domain.School

	result := u.Db.Where("id = ?", schoolId).Delete(&school)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no school found with id %d", schoolId)
	}

	return nil
}

func (u *SchoolPostgres) GetAll() []domain.School {
	var School []domain.School
	result := u.Db.Order("id").Find(&School)
	if result.Error != nil {
		return nil
	}
	return School
}

func (u *SchoolPostgres) GetBySchoolID(schoolId int) (School domain.School, err error) {
	var school domain.School
	result := u.Db.First(&school, schoolId)
	if result.Error != nil {
		return school, fmt.Errorf("school not found")
	}

	return school, nil
}

func (u *SchoolPostgres) Save(school *domain.School) error {
	result := u.Db.Create(school)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *SchoolPostgres) Update(school domain.School) error {
	var updateSchool = domain.School{
		ID:   school.ID,
		Name: school.Name,
	}

	result := u.Db.Model(&school).Updates(updateSchool)
	if result.Error != nil {
		return fmt.Errorf("can't update")
	}

	return nil
}
