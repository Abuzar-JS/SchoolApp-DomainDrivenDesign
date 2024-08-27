package student

import (
	"context"
	"data/course/domain/studentClient"
	"data/student"
	"fmt"
)

type StudentDomainClient struct {
	studentClient student.Client
}

func NewStudentDomainClient(studentClient student.Client) *StudentDomainClient {
	return &StudentDomainClient{studentClient: studentClient}
}

func (sc StudentDomainClient) GetStudentByIdClient(ctx context.Context, studentID int) (studentClient.Student, error) {
	student, err := sc.studentClient.GetStudentByIdClient(studentID)
	if err != nil {
		return studentClient.Student{}, fmt.Errorf("failed to get student from student domain client")
	}
	return studentClient.Student{ID: student.ID}, nil
}
