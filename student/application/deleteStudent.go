package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/student/domain/student"
)

type DeleteStudent func(ctx context.Context, studentID int) error

func NewDeleteStudent(
	studentRepo student.StudentRepository,
) DeleteStudent {
	return func(ctx context.Context, studentID int) error {
		if studentID == 0 {
			return fmt.Errorf("studentID must be greater than 0")
		}

		err := studentRepo.Delete(studentID)
		if err != nil {
			return fmt.Errorf("if does not exist")
		}
		return nil
	}

}
