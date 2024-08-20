package application

import (
	"context"
	"data/school/domain"
	"data/school/domain/school"
	"fmt"
)

type CreateSchoolRequest struct {
	Name string
}

func (s CreateSchoolRequest) Validate(ctx context.Context) error {
	if s.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	return nil
}

type CreateSchool func(ctx context.Context, request CreateSchoolRequest) (*domain.School, error)

func NewCreateSchool(
	schoolRepo school.SchoolRepository,
) CreateSchool {
	return func(ctx context.Context, request CreateSchoolRequest) (*domain.School, error) {
		err := request.Validate(ctx)
		if err != nil {
			return nil, err
		}

		schoolRequest := domain.School{
			Name: request.Name,
		}

		err = schoolRepo.Save(&schoolRequest)
		if err != nil {
			return nil, fmt.Errorf("school creation failed")
		}

		return &schoolRequest, nil

	}
}
