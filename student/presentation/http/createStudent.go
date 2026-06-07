package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/student/application"
	models "github.com/Abuzar-JS/Go-StudentApp/student/presentation/model"

	"github.com/gin-gonic/gin"
)

// CreateStudent godoc
// @Summary Create a student
// @Tags students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Param request body models.CreateStudentRequest true "Student payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /schools/{school_id}/student [post]
func NewCreateStudent(
	service application.CreateStudent,
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

		var body models.CreateStudentRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		request := application.CreateStudentRequest{
			Name:     body.Name,
			Class:    body.Class,
			SchoolID: ID,
		}

		student, err := service(ctx.Request.Context(), request)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Student created successfully",
			"student": student,
		})

	}
}
