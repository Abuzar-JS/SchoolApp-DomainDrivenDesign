package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/course/application"

	"github.com/gin-gonic/gin"
)

// DeleteCourse godoc
// @Summary Delete a course
// @Tags courses
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Param student_id path int true "Student ID"
// @Param course_id path int true "Course ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /schools/{school_id}/students/{student_id}/courses/{course_id} [delete]
func NewDeleteCourse(
	service application.DeleteCourse,
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

		courseID := ctx.Param("course_id")
		crID, err := strconv.Atoi(courseID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		request := application.DeleteCourseRequest{
			CourseID:  crID,
			StudentID: stID,
			SchoolID:  scID,
		}

		err = service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "course deleted successfully",
		})

	}
}
