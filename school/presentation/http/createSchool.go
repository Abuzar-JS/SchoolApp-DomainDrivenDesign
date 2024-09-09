package http

import (
	"net/http"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"
	models "github.com/Abuzar-JS/Go-StudentApp/school/presentation/model"

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
