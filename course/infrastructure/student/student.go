package student

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/Go-StudentApp/course/domain/studentClient"
	student "github.com/Abuzar-JS/Go-StudentApp/student"
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
