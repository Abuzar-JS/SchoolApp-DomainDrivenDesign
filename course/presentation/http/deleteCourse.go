package http

import (
	"data/course/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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


