package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/school/domain/school"
)

type DeleteSchool func(ctx context.Context, schoolID int) error

func NewDeleteSchool(
	schoolRepo school.Repository,
) DeleteSchool {
	return func(ctx context.Context, schoolID int) error {
		if schoolID == 0 {
			return fmt.Errorf("schoolID must be greater than 0")
		}

		err := schoolRepo.Delete(schoolID)
		if err != nil {
			return fmt.Errorf("id does not exist")
		}

		return nil
	}
}
