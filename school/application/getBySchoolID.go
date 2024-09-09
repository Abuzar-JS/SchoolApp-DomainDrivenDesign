package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/school/domain"
	"github.com/Abuzar-JS/Go-StudentApp/school/domain/school"
)

type GetBySchoolId func(ctx context.Context, schoolID int) (*domain.School, error)

func NewGetBySchoolID(
	schoolRepo school.Repository,
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
