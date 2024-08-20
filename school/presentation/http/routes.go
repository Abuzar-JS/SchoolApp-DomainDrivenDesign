package http

import (
	"data/school/application"
	"data/school/infrastructure/postgres"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {
	schoolRepo := postgres.NewSchoolPostgres(database)

	schoolRouter := router.Group("/api/v1")

	schoolRouter.GET("/schools", NewGetAllSchool(
		application.NewGetAllSchool(schoolRepo),
	))

	schoolRouter.GET("/schools/:school_id", NewGetBySchoolID(
		application.NewGetBySchoolID(schoolRepo),
	))

	schoolRouter.POST("/school", NewCreateSchool(
		application.NewCreateSchool(schoolRepo),
	))

	schoolRouter.PUT("/schools/:school_id", NewUpdateSchool(
		application.NewUpdateSchool(schoolRepo),
	))

	schoolRouter.DELETE("/schools/:school_id", NewDeleteSchool(
		application.NewDeleteSchool(schoolRepo),
	))

	return router
}
