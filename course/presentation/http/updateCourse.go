package http

import (
	"data/course/application"
	models "data/course/presentation/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewUpdateCourse(
	service application.UpdateCourse,
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

		body := models.UpdateCourseRequest{
			CourseID:  crID,
			StudentID: &stID,
			SchoolID:  scID,
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		if body.Title == nil && body.StudentID == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "atleast one field is required to update course",
			})
			return
		}

		request := application.UpdateCourseRequest{
			Title:     body.Title,
			CourseID:  body.CourseID,
			StudentID: body.StudentID,
			SchoolID:  body.SchoolID,
		}

		err = service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Course updated successfully",
		})
	}
}
