package domain

import "data/student/infrastructure/postgres"

type Course struct {
	ID        int
	Title     string
	StudentID int
	Student   []postgres.Student
}
