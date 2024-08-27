package main

import (
	"data/config"
	"data/course"
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
