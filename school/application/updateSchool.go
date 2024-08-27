package application

import (
	"context"
	"data/school/domain/school"
	"fmt"
)

type UpdateSchoolRequest struct {
	ID   int
	Name string
}

func (s UpdateSchoolRequest) Validate(ctx context.Context) error {
	if s.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if s.ID == 0 {
		return fmt.Errorf("id must be greater then 0")
	}

	return nil
}

type UpdateSchool func(ctx context.Context, request UpdateSchoolRequest) error

func NewUpdateSchool(
	schoolRepo school.Repository,
) UpdateSchool {
	return func(ctx context.Context, request UpdateSchoolRequest) error {
		schoolData, err := schoolRepo.GetBySchoolID(request.ID)
		if err != nil {
			return fmt.Errorf("cant't update school ")
		}

		schoolData.SetName(request.Name)
		schoolRepo.Update(schoolData)

		return nil
	}
}
