package http

import (
	"data/course/application"
	schoolClt "data/course/infrastructure/school"
	studentClt "data/course/infrastructure/student"

	coursePostgres "data/course/infrastructure/postgres"
	"data/school"
	"data/student"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Courses router
func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate, sc school.Client, st student.Client) *gin.Engine {

	courseRepo := coursePostgres.NewCoursePostgres(database)

	courseRouter := router.Group("/api/v1/schools/:school_id/students/:student_id")

	//Get Course student ID
	courseRouter.GET("/courses", GetCourseByStudentID(
		application.NewGetCourseByStudentID(courseRepo,
			studentClt.NewStudentDomainClient(st),
			schoolClt.NewSchoolDomainClient(sc),
		),
	))

	// Get course by course ID
	courseRouter.GET("/courses/:course_id", GetCourseByID(
		application.NewGetCourseByID(
			courseRepo,
			studentClt.NewStudentDomainClient(st),
			schoolClt.NewSchoolDomainClient(sc)),
	))

	//Create Course
	courseRouter.POST("/course", NewCreateCourse(
		application.NewCreateCourse(courseRepo,
			studentClt.NewStudentDomainClient(st),
			schoolClt.NewSchoolDomainClient(sc),
		),
	))

	//update course
	courseRouter.PUT("/courses/:course_id", NewUpdateCourse(
		application.NewUpdateCourse(courseRepo,
			studentClt.NewStudentDomainClient(st),
			schoolClt.NewSchoolDomainClient(sc),
		),
	))

	//Delete course
	courseRouter.DELETE("/courses/:course_id", NewDeleteCourse(
		application.NewDeleteCourse(courseRepo,
			studentClt.NewStudentDomainClient(st),
			schoolClt.NewSchoolDomainClient(sc),
		),
	))

	return router
}
