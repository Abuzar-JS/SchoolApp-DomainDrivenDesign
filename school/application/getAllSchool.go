package application

import (
	"data/school/domain"
	"data/school/domain/school"
)

type GetAllSchool func() []domain.School

func NewGetAllSchool(
	schoolRepo school.SchoolRepository,
) GetAllSchool {
	return func() []domain.School {
		schools := schoolRepo.GetAll()

		// var schools []domain.School

		// for _, value := range result {
		// 	School := domain.School{
		// 		ID:   value.ID,
		// 		Name: value.Name,
		// 	}
		// 	schools = append(schools, School)
		// }

		return schools

	}
}
