package http

import (
	"data/course/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCourseByID(
	service application.GetCoursebyID,
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

		courseId := ctx.Param("course_id")
		crID, err := strconv.Atoi(courseId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		request := application.GetRequestByCourseID{
			SchoolID:  scID,
			StudentID: stID,
			CourseID:  crID,
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
