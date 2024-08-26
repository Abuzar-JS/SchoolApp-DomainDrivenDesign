package http

import (
	"data/course/application"
	models "data/course/presentation/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewCreateCourse(
	service application.CreateCourse,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		schoolID := ctx.Param("school_id")
		scID, err := strconv.Atoi(schoolID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		studentID := ctx.Param("student_id")
		stID, err := strconv.Atoi(studentID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		body := models.CreateCourseRequest{
			StudentID: stID,
			SchoolID:  scID,
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		request := application.CreateCourseRequest{
			Title:     body.Title,
			StudentID: body.StudentID,
			SchoolID:  body.SchoolID,
		}

		course, err := service(ctx.Request.Context(), request)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Course created successfully",
			"course":  course,
		})

	}
}
