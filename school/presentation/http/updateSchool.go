package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"
	models "github.com/Abuzar-JS/Go-StudentApp/school/presentation/model"

	"github.com/gin-gonic/gin"
)

// UpdateSchool godoc
// @Summary Update a school
// @Tags schools
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Param request body models.UpdateSchoolRequest true "School update payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /schools/{school_id} [put]
func NewUpdateSchool(
	update application.UpdateSchool,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.UpdateSchoolRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
		}

		schoolID := ctx.Param("school_id")
		id, err := strconv.Atoi(schoolID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		body.ID = id

		updateRequest := application.UpdateSchoolRequest{
			ID:   body.ID,
			Name: body.Name,
		}

		err = update(ctx.Request.Context(), updateRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "School updated successfully",
		})
	}

}
