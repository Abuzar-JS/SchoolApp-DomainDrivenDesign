package domain

import "data/school/infrastructure/postgres"

type Student struct {
	ID       int
	Name     string
	Class    string
	SchoolID int
	School   []postgres.School
}
