package http

import (
	"data/school/application"
	models "data/school/presentation/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewCreateSchool(
	service application.CreateSchool,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body models.CreateSchoolRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		request := application.CreateSchoolRequest{
			Name: body.Name,
		}

		school, err := service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "School created successfully",
			"school":  school,
		})
	}
}
