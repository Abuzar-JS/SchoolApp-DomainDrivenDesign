package http

import (
	"data/student/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewGetStudentBySchoolID(
	service application.GetStudentBySchoolID,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		schoolID := ctx.Param("school_id")
		ID, err := strconv.Atoi(schoolID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		student, err := service(ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "student found",
			"data":    student,
		})
	}
}
