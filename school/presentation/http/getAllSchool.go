package http

import (
	"net/http"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"
	models "github.com/Abuzar-JS/Go-StudentApp/school/presentation/model"

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
