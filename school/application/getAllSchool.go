package application

import (
	"github.com/Abuzar-JS/Go-StudentApp/school/domain"
	"github.com/Abuzar-JS/Go-StudentApp/school/domain/school"
)

type GetAllSchool func() []domain.School

func NewGetAllSchool(
	schoolRepo school.Repository,
) GetAllSchool {
	return func() []domain.School {
		result := schoolRepo.GetAll()

		var schools []domain.School

		for _, value := range result {
			School := domain.School{
				ID:   value.ID,
				Name: value.Name,
			}
			schools = append(schools, School)
		}

		return schools

	}
}
