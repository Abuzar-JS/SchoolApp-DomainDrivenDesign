package http

import (
	"data/school/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewDeleteSchool(
	delete application.DeleteSchool,
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

		err = delete(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "school deleted successfully",
		})
	}
}
