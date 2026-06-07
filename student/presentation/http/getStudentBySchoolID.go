package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/student/application"

	"github.com/gin-gonic/gin"
)

// GetStudentsBySchoolID godoc
// @Summary Get students by school ID
// @Tags students
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /schools/{school_id}/students [get]
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
