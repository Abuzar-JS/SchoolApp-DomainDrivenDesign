package router

import (
	"data/course/controller"
	"data/course/repository"
	"data/course/service"
	schoolRepo "data/school/repository"
	studentRepo "data/student/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Courses router
func CourseRouter(router *gin.Engine, db *gorm.DB, validate *validator.Validate) *gin.Engine {
	courseRepository := repository.NewCourseRepositoryImpl(db)
	schoolRepository := schoolRepo.NewSchoolRepositoryImpl(db)
	StudentRepository := studentRepo.NewStudentRepositoryImpl(db)

	courseService := service.NewCourseServiceImpl(courseRepository, validate, schoolRepository, StudentRepository)

	courseController := controller.NewCourseController(courseService)

	courseRouter := router.Group("/api/v1/schools/:school_id/students/:student_id")

	courseRouter.GET("/courses", courseController.FindByStudentID)
	courseRouter.GET("/courses/:course_id", courseController.FindById)
	courseRouter.POST("/course", courseController.Create)
	courseRouter.PUT("/courses/:course_id", courseController.Update)
	courseRouter.DELETE("/courses/:course_id", courseController.Delete)

	return router
}
