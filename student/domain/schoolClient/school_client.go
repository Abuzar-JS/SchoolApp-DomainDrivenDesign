package schoolClient

import (
	"context"
)

type SchoolClient interface {
	GetBySchoolIdClient(ctx context.Context, schoolID int) (School, error)
}
