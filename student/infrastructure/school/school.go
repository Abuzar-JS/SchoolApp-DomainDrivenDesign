package school

import (
	"context"
	"fmt"

	school "github.com/Abuzar-JS/Go-StudentApp/school"
	"github.com/Abuzar-JS/Go-StudentApp/student/domain/schoolClient"
)

type SchoolDomainClient struct {
	schoolClient school.Client
}

func NewSchoolDomainClient(schoolClient school.Client) *SchoolDomainClient {
	return &SchoolDomainClient{schoolClient: schoolClient}
}

func (sc SchoolDomainClient) GetBySchoolIdClient(ctx context.Context, schoolID int) (schoolClient.School, error) {
	school, err := sc.schoolClient.GetBySchoolID(ctx, schoolID)
	if err != nil {
		return schoolClient.School{}, fmt.Errorf("failed to get school from school domain client")
	}
	return schoolClient.School{ID: school.ID}, nil
}
