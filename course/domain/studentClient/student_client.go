package studentClient

import "context"

type StudentClient interface {
	GetStudentByIdClient(ctx context.Context, studentID int) (Student, error)
}
