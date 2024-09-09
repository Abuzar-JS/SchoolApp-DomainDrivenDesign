package course

import (
	"github.com/Abuzar-JS/Go-StudentApp/course/domain/course"
	"github.com/Abuzar-JS/Go-StudentApp/course/infrastructure/postgres"
	"github.com/Abuzar-JS/Go-StudentApp/course/presentation/http"
	school "github.com/Abuzar-JS/Go-StudentApp/school"
	student "github.com/Abuzar-JS/Go-StudentApp/student"

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
