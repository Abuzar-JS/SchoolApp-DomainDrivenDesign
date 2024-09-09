package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"

	"github.com/gin-gonic/gin"
)

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
