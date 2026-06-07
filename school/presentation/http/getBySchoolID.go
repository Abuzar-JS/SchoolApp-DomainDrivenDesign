package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"

	"github.com/gin-gonic/gin"
)

// GetSchoolByID godoc
// @Summary Get a school by ID
// @Tags schools
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /schools/{school_id} [get]
func NewGetBySchoolID(
	getById application.GetBySchoolId,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		schoolId := ctx.Param("school_id")
		ID, err := strconv.Atoi(schoolId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		school, err := getById(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "school found",
			"data":    school,
		})
	}
}
