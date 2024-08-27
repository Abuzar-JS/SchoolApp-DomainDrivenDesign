package school

import (
	"context"
	"data/school"
	"data/student/domain/schoolClient"
	"fmt"
)

type SchoolDomainClient struct {
	schoolClient school.Client
}

func NewSchoolDomainClient(schoolClient school.Client) *SchoolDomainClient {
	return &SchoolDomainClient{schoolClient: schoolClient}
}

func (sc SchoolDomainClient) GetBySchoolIdClient(ctx context.Context, schoolID int) (schoolClient.School, error) {
	school, err := sc.schoolClient.GetBySchoolIdClient(ctx, schoolID)
	if err != nil {
		return schoolClient.School{}, fmt.Errorf("failed to get school from school domain client")
	}
	return schoolClient.School{ID: school.ID}, nil
}
