package school

import (
	"context"
	"data/school"
	"data/student/domain"
	"fmt"
)

type SchoolDomainClient struct {
	schoolClient school.Client
}

func NewSchoolDomainClient(schoolClient school.Client) *SchoolDomainClient {
	return &SchoolDomainClient{schoolClient: schoolClient}
}

func (sc SchoolDomainClient) GetStudentByIdClient(ctx context.Context, schoolID int) (domain.School, error) {
	school, err := sc.schoolClient.GetStudentByIdClient(ctx, schoolID)
	if err != nil {
		return domain.School{}, fmt.Errorf("failed to get school from school domain client")
	}
	return domain.School{ID: school.ID}, nil
}
