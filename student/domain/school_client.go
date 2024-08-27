package domain

import (
	"context"
)

type SchoolClient interface {
	GetStudentByIDClient(ctx context.Context, schoolID int) (School, error)
}
