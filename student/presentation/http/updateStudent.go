package http

import (
	"data/student/application"
	models "data/student/presentation/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewUpdateStudent(
	update application.UpdateStudent,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.UpdateStudentRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		studentID := ctx.Param("student_id")
		ID, err := strconv.Atoi(studentID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if body.Class == nil && body.Name == nil && body.SchoolID == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "atleast single field is required to update student",
			})
			return
		}

		updateRequest := application.UpdateStudentRequest{
			Name:     body.Name,
			Class:    body.Class,
			SchoolID: body.SchoolID,
		}

		err = update(ctx.Request.Context(), ID, updateRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Student updated successfully",
		})
	}
}
