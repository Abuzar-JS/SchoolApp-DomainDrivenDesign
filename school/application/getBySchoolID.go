package application

import (
	"context"
	"data/school/domain"
	"data/school/domain/school"
)

type GetBySchoolId func(ctx context.Context, schoolID int) *domain.School

func NewGetBySchoolID(
	schoolRepo school.SchoolRepository,
) GetBySchoolId {
	return func(ctx context.Context, schoolID int) *domain.School {
		School, err := schoolRepo.GetBySchoolID(schoolID)
		if err != nil {
			return nil
		}

		schoolResponse := domain.School{
			ID:   School.ID,
			Name: School.Name,
		}

		return &schoolResponse
	}

}
