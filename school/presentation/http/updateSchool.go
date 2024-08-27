package http

import (
	"data/school/application"
	models "data/school/presentation/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewUpdateSchool(
	update application.UpdateSchool,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.UpdateSchoolRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		schoolID := ctx.Param("school_id")
		id, err := strconv.Atoi(schoolID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		body.ID = id

		updateRequest := application.UpdateSchoolRequest{
			ID:   body.ID,
			Name: body.Name,
		}

		err = update(ctx.Request.Context(), updateRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "School updated successfully",
		})
	}

}
