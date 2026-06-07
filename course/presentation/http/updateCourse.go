package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/course/application"
	models "github.com/Abuzar-JS/Go-StudentApp/course/presentation/model"

	"github.com/gin-gonic/gin"
)

// UpdateCourse godoc
// @Summary Update a course
// @Tags courses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Param student_id path int true "Student ID"
// @Param course_id path int true "Course ID"
// @Param request body models.UpdateCourseRequest true "Course update payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /schools/{school_id}/students/{student_id}/courses/{course_id} [put]
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

		body := models.UpdateCourseRequest{}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}
		if body.Title == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "title is required to update course",
			})
			return
		}
		body.CourseID = crID
		body.StudentID = &stID
		body.SchoolID = scID

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
