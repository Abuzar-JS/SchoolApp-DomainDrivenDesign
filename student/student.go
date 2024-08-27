package student

import (
	"data/school"
	"data/student/domain"
	"data/student/domain/student"
	"data/student/infrastructure/postgres"
	"data/student/presentation/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type client struct {
	studentRepo student.StudentRepository
}

type Client interface {
	GetStudentByIdClient(studentId int) (Student domain.Student, err error)
}

func (c *client) GetStudentByIdClient(studentId int) (Student domain.Student, err error) {
	return c.studentRepo.GetStudentById(studentId)
}

func InitiateAndRegister(router *gin.Engine, database *gorm.DB, validate *validator.Validate, sc school.Client) Client {

	http.RegisterRoutes(router, database, validate, sc)

	// return client
	return &client{
		studentRepo: postgres.NewStudentPostgres(database),
	}
}
