package http

import (
	school "github.com/Abuzar-JS/Go-StudentApp/school"
	"github.com/Abuzar-JS/Go-StudentApp/student/application"
	"github.com/Abuzar-JS/Go-StudentApp/student/infrastructure/postgres"
	schoolClient "github.com/Abuzar-JS/Go-StudentApp/student/infrastructure/school"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate, sc school.Client) *gin.Engine {
	studentRepo := postgres.NewStudentPostgres(database)

	studentRouter := router.Group("/api/v1/schools")

	//Get Student By School ID
	studentRouter.GET("/:school_id/students", NewGetStudentBySchoolID(
		application.NewGetStudentBySchoolID(
			studentRepo,
			schoolClient.NewSchoolDomainClient(sc),
		),
	))

	//Get Student By Student ID
	studentRouter.GET("/:school_id/students/:student_id", NewGetByStudentID(
		application.NewGetByStudentID(studentRepo),
	))

	// Create Student
	studentRouter.POST("/:school_id/student", NewCreateStudent(
		application.NewCreateStudent(studentRepo),
	))

	//Update Student
	studentRouter.PUT("/:school_id/students/:student_id", NewUpdateStudent(
		application.NewUpdateStudent(studentRepo),
	))

	//Delete Student
	studentRouter.DELETE("/:school_id/students/:student_id", NewDeleteStudent(
		application.NewDeleteStudent(studentRepo),
	))

	return router

}
