package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Abuzar-JS/Go-StudentApp/auth"
	config "github.com/Abuzar-JS/Go-StudentApp/config"
	course "github.com/Abuzar-JS/Go-StudentApp/course"
	_ "github.com/Abuzar-JS/Go-StudentApp/docs"
	school "github.com/Abuzar-JS/Go-StudentApp/school"
	student "github.com/Abuzar-JS/Go-StudentApp/student"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Student App API
// @version 1.0
// @description API documentation for schools, students, and courses.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter the token as: Bearer <API_AUTH_TOKEN>
func main() {
	log.Info().Msg("Started Server! ")

	authToken := os.Getenv("API_AUTH_TOKEN")
	if authToken == "" {
		log.Fatal().Msg("API_AUTH_TOKEN is required")
	}

	db := config.DatabaseConnection()
	validate := validator.New()

	ginRouter := gin.Default()
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginRouter.Use(auth.RequireBearerToken(authToken))

	schoolClient := school.InitiateAndRegister(ginRouter, db, validate)
	studentClient := student.InitiateAndRegister(ginRouter, db, validate, schoolClient)
	course.InitiateAndRegister(ginRouter, db, validate, schoolClient, studentClient)

	server := &http.Server{
		Addr:    ":8080",
		Handler: ginRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server not started")
	}

}
