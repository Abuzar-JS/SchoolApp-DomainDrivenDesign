package http

import (
	"data/course/application"
	coursePostgres "data/course/infrastructure/postgres"
	schoolPostgres "data/school/infrastructure/postgres"
	studentPostgres "data/student/infrastructure/postgres"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Courses router
func RegisterCourseRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {
	studentRepo := studentPostgres.NewStudentPostgres(database)
	schoolRepo := schoolPostgres.NewSchoolPostgres(database)

	courseRepo := coursePostgres.NewCoursePostgres(database)

	courseRouter := router.Group("/api/v1/schools/:school_id/students/:student_id")

	courseRouter.GET("/courses", GetCourseByStudentID(
		application.NewGetCourseByStudentID(courseRepo, studentRepo, schoolRepo),
	))
	courseRouter.GET("/courses/:course_id", GetCourseByID(
		application.NewGetCourseByID(courseRepo, studentRepo, schoolRepo),
	))

	courseRouter.POST("/course", NewCreateCourse(
		application.NewCreateCourse(courseRepo, studentRepo, schoolRepo),
	))

	courseRouter.PUT("/courses/:course_id", NewUpdateCourse(
		application.NewUpdateCourse(courseRepo, studentRepo, schoolRepo),
	))

	courseRouter.DELETE("/courses/:course_id", NewDeleteCourse(
		application.NewDeleteCourse(courseRepo, studentRepo, schoolRepo),
	))

	return router
}
