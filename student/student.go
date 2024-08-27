package student

import (
	"data/school"
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
}

func InitiateAndRegister(router *gin.Engine, database *gorm.DB, validate *validator.Validate, sc school.Client) Client {
	// register tariff routes

	http.RegisterRoutes(router, database, validate, sc)

	// return clien
	return &client{
		studentRepo: postgres.NewStudentPostgres(database),
	}
}
