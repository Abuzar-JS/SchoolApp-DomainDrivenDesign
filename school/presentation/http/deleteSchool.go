package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"

	"github.com/gin-gonic/gin"
)

// DeleteSchool godoc
// @Summary Delete a school
// @Tags schools
// @Produce json
// @Security BearerAuth
// @Param school_id path int true "School ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /schools/{school_id} [delete]
func NewDeleteSchool(
	delete application.DeleteSchool,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		schoolID := ctx.Param("school_id")
		id, err := strconv.Atoi(schoolID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		err = delete(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "school deleted successfully",
		})
	}
}
