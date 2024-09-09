package school

import "github.com/Abuzar-JS/Go-StudentApp/school/domain"

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetBySchoolID(schoolId int) (School domain.School, err error)
	GetAll() []domain.School
}

type WriteRepository interface {
	Delete(schoolId int) error
	Save(school *domain.School) error
	Update(school domain.School) error
}
