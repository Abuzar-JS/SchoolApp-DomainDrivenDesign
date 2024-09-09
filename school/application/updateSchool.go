package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/school/domain/school"
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

		if err := request.Validate(ctx); err != nil {
			return err
		}
		schoolData, err := schoolRepo.GetBySchoolID(request.ID)
		if err != nil {
			return fmt.Errorf("no school found with id %d", request.ID)
		}

		schoolData.SetName(request.Name)
		schoolData.SetID(request.ID)
		schoolRepo.Update(schoolData)

		return nil
	}
}
