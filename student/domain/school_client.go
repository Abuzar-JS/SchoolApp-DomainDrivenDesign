package domain

import (
	"context"
)

type SchoolClient interface {
	GetStudentByIdClient(ctx context.Context, schoolID int) (School, error)
}
