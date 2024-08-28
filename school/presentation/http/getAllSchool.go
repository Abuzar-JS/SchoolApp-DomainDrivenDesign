package http

import (
	"data/school/application"
	models "data/school/presentation/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGetAllSchool(
	getAll application.GetAllSchool,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		schoolResponse := getAll()
		webResponse := models.Response{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   schoolResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}
