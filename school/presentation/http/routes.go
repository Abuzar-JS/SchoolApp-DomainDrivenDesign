package http

import (
	"data/school/application"
	"data/school/infrastructure/postgres"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) {
	schoolRepo := postgres.NewSchoolPostgres(database)

	schoolRouter := router.Group("/api/v1")

	//Create School
	schoolRouter.POST("/school", NewCreateSchool(
		application.NewCreateSchool(schoolRepo),
	))

	//Get All Schools
	schoolRouter.GET("/schools", NewGetAllSchool(
		application.NewGetAllSchool(schoolRepo),
	))
	// Get School By ID
	schoolRouter.GET("/schools/:school_id", NewGetBySchoolID(
		application.NewGetBySchoolID(schoolRepo),
	))

	//Update School
	schoolRouter.PUT("/schools/:school_id", NewUpdateSchool(
		application.NewUpdateSchool(schoolRepo),
	))

	//Delete School
	schoolRouter.DELETE("/schools/:school_id", NewDeleteSchool(
		application.NewDeleteSchool(schoolRepo),
	))
}
