package http

import (
	"data/student/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
