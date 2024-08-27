package http

import (
	"data/school"
	schoolPostgres "data/school/infrastructure/postgres"
	"data/student/application"
	"data/student/domain"
	"data/student/infrastructure/postgres"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate, sc school.Client) *gin.Engine {
	studentRepo := postgres.NewStudentPostgres(database)

	schoolRepo := schoolPostgres.NewSchoolPostgres(database)

	studentRouter := router.Group("/api/v1/schools")

	studentRouter.GET("/:school_id/students", NewGetStudentBySchoolID(
		application.NewGetStudentBySchoolID(studentRepo, domain.SchoolClient),
	))

	studentRouter.GET("/:school_id/students/:student_id", NewGetByStudentID(
		application.NewGetByStudentID(studentRepo),
	))

	studentRouter.POST("/:school_id/student", NewCreateStudent(
		application.NewCreateStudent(studentRepo),
	))

	studentRouter.PUT("/:school_id/students/:student_id", NewUpdateStudent(
		application.NewUpdateStudent(studentRepo),
	))

	studentRouter.DELETE("/:school_id/students/:student_id", NewDeleteStudent(
		application.NewDeleteStudent(studentRepo),
	))

	return router

}
