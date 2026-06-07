package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/student/application"

	"github.com/gin-gonic/gin"
)

// DeleteStudent godoc
// @Summary Delete a student
// @Tags students
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Param student_id path int true "Student ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /schools/{school_id}/students/{student_id} [delete]
func NewDeleteStudent(
	delete application.DeleteStudent,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		studentID := ctx.Param("student_id")
		ID, err := strconv.Atoi(studentID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = delete(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "student deleted successfully",
		})

	}

}
