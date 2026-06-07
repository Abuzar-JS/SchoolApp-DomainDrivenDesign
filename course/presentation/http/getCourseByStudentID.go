package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/course/application"

	"github.com/gin-gonic/gin"
)

// GetCoursesByStudentID godoc
// @Summary Get courses by student ID
// @Tags courses
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Param student_id path int true "Student ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /schools/{school_id}/students/{student_id}/courses [get]
func GetCourseByStudentID(
	service application.GetCourseByStudentID,
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

		request := application.GetCourseRequestByStudentID{
			SchoolID:  scID,
			StudentID: stID,
		}

		course, err := service(request)
		if err != nil {
			ctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "course found",
			"data":    course,
		})
	}

}
