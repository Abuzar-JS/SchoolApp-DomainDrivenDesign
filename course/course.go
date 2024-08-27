package course

import (
	"data/course/domain/course"
	"data/course/infrastructure/postgres"
	"data/course/presentation/http"
	"data/school"
	"data/student"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type client struct {
	courseRepo course.CourseRepository
}

type Client interface {
}

func InitiateAndRegister(router *gin.Engine, database *gorm.DB, validate *validator.Validate, sc school.Client, st student.Client) Client {

	http.RegisterRoutes(router, database, validate, sc, st)

	return &client{

		courseRepo: postgres.NewCoursePostgres(database),
	}
}
