package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/student/application"

	"github.com/gin-gonic/gin"
)

func NewGetByStudentID(
	service application.GetByStudentID,
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

		studentID := ctx.Param("student_id")
		fmt.Println(studentID)
		ID, err = strconv.Atoi(studentID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		student, err := service(ID)
		if err != nil {
			ctx.JSON(404, gin.H{
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
