package student

import "data/student/domain"

type StudentRepository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetBySchoolID(studentID int) ([]domain.Student, error)
	GetStudentById(studentId int) (Student domain.Student, err error)
}

type WriteRepository interface {
	Delete(studentId int) error
	Save(student *domain.Student) error
	Update(id int, student domain.Student) error
}
