package application

import (
	"context"
	"data/school/domain"
	"data/school/domain/school"
	"fmt"
)

type GetBySchoolId func(ctx context.Context, schoolID int) (*domain.School, error)

func NewGetBySchoolID(
	schoolRepo school.SchoolRepository,
) GetBySchoolId {
	return func(ctx context.Context, schoolID int) (*domain.School, error) {
		School, err := schoolRepo.GetBySchoolID(schoolID)
		if err != nil {
			return nil, fmt.Errorf("no school found")
		}

		schoolResponse := domain.School{

			ID:   School.ID,
			Name: School.Name,
		}

		return &schoolResponse, nil
	}

}
