package main

import (
	"data/config"
	courseRoutes "data/course/presentation/http"
	"data/school"
	"data/student"

	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server! ")

	db := config.DatabaseConnection()
	validate := validator.New()

	ginRouter := gin.Default()

	sClient := school.InitiateAndRegister(ginRouter, db, validate)
	student.InitiateAndRegister(ginRouter, db, validate, sClient)

	courseRoutes.RegisterCourseRoutes(ginRouter, db, validate)

	server := &http.Server{
		Addr:    ":8080",
		Handler: ginRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server not started")
	}

}
