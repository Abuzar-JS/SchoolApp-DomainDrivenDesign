package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/student/application"

	"github.com/gin-gonic/gin"
)

// GetStudentByID godoc
// @Summary Get a student by ID
// @Tags students
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Param student_id path int true "Student ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /schools/{school_id}/students/{student_id} [get]
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
